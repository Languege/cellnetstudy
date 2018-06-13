package main

import (
	"cellnetstudy/mychat/proto/pb"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	"github.com/davyxu/cellnet/proc"
	"github.com/davyxu/golog"

	"fmt"
	_ "github.com/davyxu/cellnet/peer/tcp"
	_ "github.com/davyxu/cellnet/proc/tcp"
	"reflect"
)

var log = golog.New("server")

func main() {

	// 创建一个事件处理队列，整个服务器只有这一个队列处理事件，服务器属于单线程服务器
	queue := cellnet.NewEventQueue()

	// 创建一个tcp的侦听器，名称为server，连接地址为127.0.0.1:8801，所有连接将事件投递到queue队列,单线程的处理（收发封包过程是多线程）
	p := peer.NewGenericPeer("tcp.Acceptor", "server", "127.0.0.1:18801", queue)

	// 设定封包收发处理的模式为tcp的ltv(Length-Type-Value), Length为封包大小，Type为消息ID，Value为消息内容
	// 每一个连接收到的所有消息事件(cellnet.Event)都被派发到用户回调, 用户使用switch判断消息类型，并做出不同的处理
	proc.BindProcessorHandler(p, "tcp.ltv", func(ev cellnet.Event) {

		switch msg := ev.Message().(type) {
		// 有新的连接
		case *cellnet.SessionAccepted:
			log.Debugln("server accepted")
		// 有连接断开
		case *cellnet.SessionClosed:
			log.Debugln("session closed: ", ev.Session().ID())
		// 收到某个连接的ChatREQ消息
		case *pb.ChatREQ:

			// 准备回应的消息
			ack := pb.ChatACK{
				Content: msg.Content,       // 聊天内容
				Id:      ev.Session().ID(), // 使用会话ID作为发送内容的ID
			}

			// 在Peer上查询SessionAccessor接口，并遍历Peer上的所有连接，并发送回应消息（即广播消息）
			p.(cellnet.SessionAccessor).VisitSession(func(ses cellnet.Session) bool {

				ses.Send(&ack)

				return true
			})

			ev.Session().Send(&ack)
		case *pb.ChatP2P: //期望实现将消息发给指定用户
			log.Debugln("ChatP2P Content:", msg.Content, "From:", ev.Session().ID(), " To:", msg.Uid)
			// ev.Session().Send(msg)
			//获取接受方session
			sesId := new(int64)
			p.(cellnet.ContextSet).GetContext(msg.Uid, sesId)
			log.Debugln("revSession: ", *sesId)
			ses := p.(cellnet.SessionAccessor).GetSession(*sesId)
			ses.Send(msg)

		case *pb.BindUid:
			log.Debugln("Uid ", msg.Uid, " Bind Session:", ev.Session().ID())
			p.(cellnet.ContextSet).SetContext(msg.Uid, ev.Session().ID())

			sesId := new(int64)
			ok := p.(cellnet.ContextSet).GetContext(msg.Uid, sesId)
			if ok {
				fmt.Print("ok", *sesId)
			}

			ses := p.(cellnet.SessionAccessor).GetSession(*sesId)

			fmt.Print(reflect.TypeOf(ses))

			ses.Send(msg)
		}

	})

	// 开始侦听
	p.Start()

	// 事件队列开始循环
	queue.StartLoop()

	// 阻塞等待事件队列结束退出( 在另外的goroutine调用queue.StopLoop() )
	queue.Wait()

}
