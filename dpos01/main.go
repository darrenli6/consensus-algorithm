package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Block struct {
	Index int
	TimeStamp string
	BPM int
	HachCode string
	PrevHash string
	Validator string

}

var BlockChain []Block

// 生成区块
func GenerateNextBlock(oldBlock Block,BPM int,adds string) Block{
	var newBlock Block
	newBlock.Index=oldBlock.Index+1
	newBlock.PrevHash=oldBlock.HachCode
	newBlock.BPM=BPM
	newBlock.TimeStamp=oldBlock.TimeStamp
	newBlock.Validator=adds

	//计算hash
	newBlock.HachCode=GenerateHashValue(newBlock)
	return newBlock
}

// 计算hash的方法
func GenerateHashValue(block Block) string{
	var hashcode=block.PrevHash+
		block.TimeStamp+block.Validator+
		strconv.Itoa(block.BPM)+strconv.Itoa(block.Index)

	//哈希
	var sha=sha256.New()
	sha.Write([]byte(hashcode))
	hashed:=sha.Sum(nil)
	return hex.EncodeToString(hashed)

}

var delegate = []string{"aaa","bbb","ccc","ddd"}

// 模拟对委托人未知进行随机处理
// 后面让四个委托人轮询挖矿
// 随机位置i处理，被攻击的概率变小，降低风险

func RandDelegate(){
	rand.Seed(time.Now().Unix())
	//产生一个0到2的随机值
	var r=rand.Intn(3)
	t:=delegate[r]
	delegate[r]=delegate[3]
	delegate[3]=t
}


func main(){
	//测试随机未知
	fmt.Println(delegate)
	RandDelegate()
	fmt.Println(delegate)

	var firstBlock Block
	// 将创世区块加入区块链
	BlockChain=append(BlockChain,firstBlock)

	var n=0
	for  {
		// 每三十秒产生新的区块
		time.Sleep(time.Second*3)
		var nextBlock=GenerateNextBlock(firstBlock,1,delegate[n])
		// 轮询挖矿人
		n++
		n=n%4
		firstBlock=nextBlock
		BlockChain = append(BlockChain,nextBlock)
		fmt.Println(BlockChain)
	}
}