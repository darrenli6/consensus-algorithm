package main

import "sync"

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
	
}