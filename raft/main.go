package main

import (
	"math/rand"
	"sync"
	"time"
)

//定义常量
const raftCount =3

type Leader struct {
	//任期
	Term int
	//领导编号
	LeaderId int
}

//创建存储leader的对象
//最初任期是0 -1代表没有编号
var leader=Leader{0,-1}

type Raft struct {
	//锁
	mu sync.Mutex
	//节点编号
	me int
	//当前任期
	currentTerm int
	//为哪个节点投票
	votedFor int
	//当前节点状态
	//0 follower 1 candidate 2 leader
	state int
	//发送最后一条消息的时间
	lastMessageTime int64
	//当前节点的领导
	currentLeader int
	//消息通道
	message chan bool
	// 选举通道
	electCh chan bool
	//心跳信号
	heartBeat chan bool
	// 返回心跳信号
	hearbeatRe chan bool

	//超时时间
	timeout int
}


func main(){
	//过程 创建三个节点 最初是follower状态
	//如果出现candidate状态的节点 则开始投票
	//产生leader

    //创建三个节点
	for i:=0;i<raftCount;i++{
		//定义make
		Make(i)
	}
}



// 创建Make
func Make(me int ) *Raft{

	rf:=&Raft{}
	rf.me=me
	//给 0 1 2 投票
	rf.votedFor=-1
    // 0 follwer状态
    rf.state=0

    rf.timeout=0

    // 最初没有领导
    rf.currentLeader=-1
    //设置任期
    rf.setTerm(0)

    // 通道
    rf.electCh=make(chan bool)
    rf.message=make(chan bool)

    rf.heartBeat=make(chan bool)

    rf.hearbeatRe =make(chan bool)

    //随机种子
    rand.Seed(time.Now().UnixNano())

    //选举的逻辑实现
    go rf.election()
    // 心跳检查
    go rf.sendLderHear()


	return nil

}

func (rf *Raft) setTerm(term int){
	rf.currentTerm=term

}

// 设置节点选举
func (rf *Raft) election(){

	for {

	}
}

// 产生随机值

func randRange(min,max int64) int64{
	// 用于心跳信息的时间间隔
	return rand.Int63n(max-min)+min
}

