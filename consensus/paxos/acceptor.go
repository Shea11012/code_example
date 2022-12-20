package paxos

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type Acceptors struct {
	lis net.Listener
	// 服务器ID
	id int
	// 接受者承诺的提案编号，为0，则表示接受者没有收到过任何prepare消息
	minProposal int
	// 接受者已接受的提案编号
	acceptedNumber int
	// 接受者已接受的提案值，如果没有接受任何提案，则为nil
	acceptedValue interface{}
	// 学习者ID列表
	learners []int
}

func newAcceptors(id int, learners []int) *Acceptors {
	acceptor := &Acceptors{
		id: id,
		learners: learners,
	}

	acceptor.server()
	return acceptor
}

func (a *Acceptors) Prepare(args *MsgArgs, reply *MsgReply) error {
	if args.Number > a.minProposal {
		a.minProposal = args.Number
		reply.Number = a.acceptedNumber
		reply.Value = a.acceptedValue
		reply.Ok = true
	} else {
		reply.Ok = false
	}

	return nil
}

func (a *Acceptors) Accept(args *MsgArgs, reply *MsgReply) error {
	if args.Number >= a.minProposal {
		a.minProposal = args.Number
		a.acceptedNumber = args.Number
		a.acceptedValue = args.Value
		reply.Ok = true

		for _, l := range a.learners {
			go func(l int) {
				addr := fmt.Sprintf("127.0.0.1:%d", l)
				args.From = a.id
				args.To = l
				resp := new(MsgReply)
				ok := call(addr,"Learner.Learn", args, resp)
				if !ok {
					return
				}
			}(l)
		}
	} else {
		reply.Ok = false
	}

	return nil
}

func (a *Acceptors) server() {
	rpcs := rpc.NewServer()
	rpcs.Register(a)
	addr := fmt.Sprintf(":%d",a.id)
	l,e := net.Listen("tcp",addr)
	if e != nil {
		log.Fatal("listen error:", e)
	}

	a.lis = l

	go func ()  {
		for {
			conn,err := a.lis.Accept()
			if err != nil {
				continue
			}

			go rpcs.ServeConn(conn)
		}	
	}()
}

func (a *Acceptors) close() {
	a.lis.Close()
}
