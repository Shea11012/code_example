package snowflake

import (
	"errors"
	"sync"
	"time"
)

// snowflake leaf bit划分
// 1 | 41 | 10 | 12

const (
	Epoch     int64 = 1663459200000 // 2022-09-18 00:00:00 +0000 UTC
	TimeShift       = 41
	// SequenceBit 在同一毫秒内可自增的数量
	SequenceBit = 12
	// NodeBit 可部署的节点数量
	NodeBit = 63 - TimeShift - SequenceBit

	sequenceMax = -1 ^ (-1 << SequenceBit)
	nodeMax     = -1 ^ (-1 << NodeBit)
	nodeShift   = SequenceBit
)

var (
	ErrClockMoveBackward = errors.New("clock moved backwards")
)

type Node struct {
	mu             sync.Mutex
	lastTime       int64 // 上一次的时间戳
	sequence       int64
	nodeID         int64
	GetCurrentTime func() int64
}

func NewNode(nodeID int64) *Node {
	n := &Node{}
	n.nodeID = nodeID
	n.GetCurrentTime = time.Now().UTC().UnixMilli
	n.lastTime = n.GetCurrentTime()

	return n
}

func (n *Node) NextId() (int64, error) {
	n.mu.Lock()
	defer n.mu.Unlock()

	currentTime := n.GetCurrentTime()
	if currentTime < n.lastTime {
		return 0, ErrClockMoveBackward
	}

	if currentTime == n.lastTime {
		n.sequence = (n.sequence + 1) & sequenceMax
		// 当sequence等于0时，表示seq到了1ms内的最大值了,需要等待下1ms再次生成
		if n.sequence == 0 {
			for currentTime <= n.lastTime {
				currentTime = n.GetCurrentTime()
			}
		}
	} else {
		n.sequence = 0
	}

	n.lastTime = currentTime

	r := currentTime<<TimeShift | (n.nodeID << nodeShift) | (n.sequence)
	return r, nil
}
