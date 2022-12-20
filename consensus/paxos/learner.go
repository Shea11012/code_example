package paxos

import "net"

type Learner struct {
	lis net.Listener
	// 学习者ID
	id int
	// 记录接受者已接受的提案，[接受者，id]请求消息
	accpetedMsg map[int]MsgArgs
}