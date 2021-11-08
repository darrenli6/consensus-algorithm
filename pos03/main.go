package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go-spew/go-spew-master/spew"
	"github.com/godotenv/godotenv-master"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

type Block struct {
	Index int
	TimeStamp string
	BPM int
	HashCode string
	PrevHash string
	Validator string
}
//声明链
var Blockchain []Block


// 临时缓冲区
var tempBlocks []Block
//声明候选人,
//任何一个节点提议一个新块的时候，将它发送到这个信道
var condidateBlocks=make(chan Block)

//公告的信道 用于网络广播的内容
var announcements =make(chan string)

//锁
var mutex=&sync.Mutex{}

//验证者列表
// 存放地址和tokens
var validators=make(map[string]int)

//生成区块
func generateBlock(oldBlock Block,BPM int,address string) (Block){

	var newBlock Block
	newBlock.Index=oldBlock.Index+1
	newBlock.TimeStamp=time.Now().String()
	newBlock.BPM=BPM
	newBlock.PrevHash=oldBlock.HashCode
	newBlock.Validator=address

	newBlock.HashCode=GenerateHashValue(newBlock)
	return newBlock
}

// 计算hash的方法
func GenerateHashValue(block Block) string{
	var hashcode=block.PrevHash+
		block.TimeStamp+block.Validator+
		strconv.Itoa(block.BPM)+strconv.Itoa(block.Index)

	return calculateHash(hashcode)
}


func calculateHash(s string) string{
	//哈希
	var sha=sha256.New()
	sha.Write([]byte(s))
	hashed:=sha.Sum(nil)
	return hex.EncodeToString(hashed)

}

func main(){


	// 加载本地.env文件
	err:=godotenv.Load("D:\\project\\Go\\consensus-algorithm\\pos03\\.env")
	if err!=nil {
       log.Fatal(err)
	}

	//创建创世区块
	genesisBlock:=Block{}
	genesisBlock=Block{0,time.Now().String(),0,GenerateHashValue(genesisBlock),"",""}
	spew.Dump(genesisBlock)


	//读port
	port:=os.Getenv("PORT")

	//启动服务器
	server,err:=net.Listen("tcp",":"+port)
    if err!=nil{
		log.Fatal(err)
	}
	//打印监听到的端口
	log.Println("Http server listening on port",port)

	defer server.Close()

	go func(){
        for cadidate := range condidateBlocks{

			//加上锁
			mutex.Lock()
			//候选人中有数据 添加到临时缓冲区
			tempBlocks=append(tempBlocks,cadidate)

			mutex.Unlock()
		}
	}()


	// 查谁去挖矿
	go func() {
		for {
			//根据tokens个数去做比重划分
			pickWinner()
		}
	}()


	// 接受验证者节点的链接
	for {
		//等待终端的链接
		conn,err:=server.Accept()
		if err!=nil{
			log.Fatal(err)
		}
		//连上的情况下，处理终端发来的消息
		go handleConn(conn)
	}

}

// 由POS的主要逻辑
// 实现获取记账权的节点
// 根据令牌数量
func pickWinner(){
   // 休眠30秒
	time.Sleep(30*time.Second)
	//锁
	mutex.Lock()
	temp:=tempBlocks

	mutex.Unlock()

	//声明一个彩票池，存放每一个验证者的地址
	lotteryPool :=[]string{}

	if len(temp)>0{
		// 有验证者
		// 根据被标记的令牌的数量对他们进行加权
		// 遍历temp
		OUTER:
			for _,block:=range temp{
				// 查看是否已在彩票池
				for _,node:=range lotteryPool {
					//如果遍历的节点在彩票池 让跳出
					if block.Validator==node{
						//跳出
						continue OUTER
					}
				}


				// 锁
				mutex.Lock()
				// 地址和tokens
				setValidators :=validators
				mutex.Unlock()

				// 获取验证者的tokens的个数
				// k 当前验证者的tokens
				k,ok:=setValidators[block.Validator]
                if ok{
					//向彩票池加k条数据
                     // 将所有的验证者加入到一个数组中
					for i:=0;i<k;i++{
						lotteryPool=append(lotteryPool,block.Validator)
					}
				}



			}


        //设置随机种子 保证随机性
		s:=rand.NewSource(time.Now().Unix())
		r:=rand.New(s)
		//通过随机值（0 彩票池长度） 随机获得记账的节点
		lotteryWinner:= lotteryPool[r.Intn(len(lotteryPool))]
        // 把获胜者的区块添加到整条区块链上
		//通知其他节点关于获胜者的消息
		for _,block :=range temp{

			//是否是被选中
			if block.Validator == lotteryWinner{
				// 是拥有记账权的节点
				mutex.Lock()

				Blockchain=append(Blockchain,block)
				mutex.Unlock()

				//广播消息
				for _ =range validators{
					// 将获胜者地址放到公告中
				    announcements <- "\nvalidator:"+lotteryWinner+"\n"
				}
				break
			}
		}
	}


	//临时缓冲区为空的情况
	mutex.Lock()
	tempBlocks=[]Block{}
	mutex.Unlock()




}

//处理终端cmd发来的信息
func handleConn(conn net.Conn){
    // 释放资源
	defer conn.Close()

	go func(){

		// 打印获胜者的消息
           for {
			   msg := <- announcements
			   io.WriteString(conn,msg)

		   }

	} ()

	//验证者的地址
	var address string
	// cmd 验证者输入拥有的tokens
	io.WriteString(conn,"Enter token balance:")

	scanBalance := bufio.NewScanner(conn)

	for scanBalance.Scan() {
		//获取输入的数据 转成int
		// 获取余额，持币的数量
		balance,err := strconv.Atoi(scanBalance.Text())
		if err!=nil{
			log.Printf("%v  not a number : %v",scanBalance.Text(),err)
		}

		//生成验证者的地址
		address = calculateHash(time.Now().String())

		//将验证者的地址
		validators[address] =balance
		fmt.Println(validators)

		break


	}


	// 输入交易信息
	io.WriteString(conn,"\n Enter a new BPM :")
	//获取输入的数据 转化为int
	scanBPM:= bufio.NewScanner(conn)

	go func() {
		//多次输入交易信息
		for scanBPM.Scan(){
			bmp,err:=strconv.Atoi(scanBPM.Text())
			if err!=nil{
				log.Printf("%v not a number %v",scanBPM.Text(),err)
				//从map中移除验证者信息 以后就没有记账权了
				//对恶意节点的惩罚
				delete(validators,address)
				conn.Close()
			}
            //取到区块
			mutex.Lock()

			oldLastIndex:=Blockchain[len(Blockchain)-1]
			mutex.Unlock()

			//创建新的区块
			newBlock:=generateBlock(oldLastIndex,bmp,address)

			if err!=nil{
				log.Println(err)
				continue
			}
			//验证区块
			if isBlockValid(newBlock,oldLastIndex){
				//验证通过 将新的区块 ，加入通道

				condidateBlocks <- newBlock

			}



		}
	}()

    //周期性打印区块消息
	for  {
		time.Sleep(time.Minute)
		mutex.Lock()

		//json输出
		output,err:=json.Marshal(Blockchain)

		mutex.Unlock()
		if err!=nil{
			log.Fatal(err)
		}
		//输出cmd
		io.WriteString(conn,string(output)+"\n")


	}



}

func isBlockValid(newBlock,oldBlock Block)bool{

	// 检查index
	if oldBlock.Index+1 != newBlock.Index{
		return false
	}
	//prehash

	if oldBlock.HashCode!=newBlock.PrevHash {
		return false
	}

	//再次验证
	if GenerateHashValue(newBlock) !=newBlock.HashCode{
		return false

	}
	return true

}
