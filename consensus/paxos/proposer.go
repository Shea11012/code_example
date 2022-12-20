package paxos

import "fmt"

type Proposer struct {
	// 服务器ID
	id int
	// 当前提议者已知的最大轮次
	round int
	// 提议编号
	number int
	// 接收者id列表
	acceptors []int
}

func (p *Proposer) Propose(value interface{}) interface{} {
	p.round++
	p.number = p.proposalNumber()

	// 第一阶段
	prepareCount := 0
	maxNumber := 0
	for _, aid := range p.acceptors {
		args := &MsgArgs{
			Number: p.number,
			From:   p.id,
			To:     aid,
		}
		reply := new(MsgReply)
		ok := call(fmt.Sprintf("127.0.0.1:%d", aid), "Acceptor.Prepare", args, reply)
		if !ok {
			continue
		}

		if reply.Ok {
			prepareCount++
			if reply.Number > maxNumber {
				maxNumber = reply.Number
				value = reply.Value
			}
		}

		if prepareCount == p.majority() {
			break
		}
	}

	// 第二阶段
	acceptCount := 0
	if prepareCount >= p.majority() {
		for _, aid := range p.acceptors {
			args := &MsgArgs{
				Number: p.number,
				Value:  value,
				From:  p.id,
				To: aid,
			}
			reply := new(MsgReply)
			ok := call(fmt.Sprintf("127.0.0.1:%d", aid), "Acceptor.Accept", args, reply)
			if !ok {
				continue
			}

			if reply.Ok {
				acceptCount++
			}

			if acceptCount >= p.majority() {
				return value
			}
		}
	}

	return nil
}

func (p *Proposer) majority() int {
	return len(p.acceptors) / 2 + 1
}

// 提案编号=(轮次，服务器ID)
func (p *Proposer) proposalNumber() int {
	return p.round << 16 | p.id
}