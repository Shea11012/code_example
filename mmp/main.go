package main

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

const maxMMapSize = 0x8000000000
const maxMapStep = 1 << 30 // 1GB

func main() {
	file,err := os.OpenFile("my.db",os.O_RDWR|os.O_CREATE,0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	stat,err := os.Stat("my.db")
	if err != nil {
		panic(err)
	}

	size,err := mmapSize(int(stat.Size()))
	if err != nil {
		panic(err)
	}

	// 创建新文件时，大小是0，为0时会触发系统 SIGBUS 信号，所以需要通过系统调用 ftruncate 调整大小
	err = unix.Ftruncate(int(file.Fd()),int64(size))
	if err != nil {
		panic(err)
	}

	
	b,err := unix.Mmap(int(file.Fd()),0,size,unix.PROT_READ|unix.PROT_WRITE,unix.MAP_SHARED)
	if err != nil {
		panic(err)
	}

	for index,bb := range []byte("hello world") {
		b[index] = bb
	}

	err = unix.Munmap(b)
	if err != nil {
		panic(err)
	}

}

func mmapSize(size int) (int,error) {
	if size > maxMMapSize {
		return 0, fmt.Errorf("mmap too large")
	}

	// 从32KB增长到1GB
	for i := uint(15); i <= 30; i++ {
		if size <= 1<<i {
			return 1<<i,nil
		}
	}


	// 大于1GB,但不满足 maxMapStep 的倍数的，补足
	sz := int64(size)
	if remainder := sz % int64(maxMapStep); remainder > 0 {
		sz += int64(maxMapStep) - remainder
	}

	// 确保 map size 有多页
	pageSize := int64(os.Getpagesize())
	if (sz % pageSize) != 0 {
		sz = ((sz / pageSize) + 1 ) * pageSize
	}

	if sz > maxMMapSize {
		sz = maxMMapSize
	}

	return int(sz),nil
}