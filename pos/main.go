package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//实现pos算法

// 定义区块
type Block struct {
	Index int
	TimeStamp string
	BPM int
	HashCode string
	PrevHash string
	// 区块验证者
	Validator string

}

//创建区块链 数组
var Blockchain []Block

//生成新的区块
// address 矿工地址
func CenerateNextBlock(oldBlock Block,BPM int,adress string) Block{

	var newBlock Block
	newBlock.Index=oldBlock.Index+1
	newBlock.TimeStamp=time.Now().String()
	newBlock.PrevHash=oldBlock.HashCode
	newBlock.BPM=BPM
	//挖矿节点地址
	newBlock.Validator=adress
	// 哈希计算
	newBlock.HashCode=GenerateHashValue(newBlock)
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


//存放几个节点 有几个用户在参与
var n [2] Node

// 用于记录挖矿地址
var addr [6000] string

// 网络上的全节点

type Node struct {
	// 记录有多少个token
	tokens int
	// 节点地址
	address string

}

func main(){

	//var firstBlock Block
	//myBlock :=CenerateNextBlock(firstBlock,1,"abc")
	//fmt.Println(myBlock)

	n[0]=Node{tokens:1000,address:"abcd123"}
	n[1]=Node{tokens:5000,address:"abc123"}

	//以下pos
	var count=0
	for i:=0;i<len(n);i++{
		for j:=0;j<n[i].tokens;j++{
			addr[count]=n[i].address
			count++
		}
	}

	//设置随机种子
	rand.Seed(time.Now().Unix())
	// 通过随机值
	var rd=rand.Intn(6000)
    // 随机选矿工
	var adds=addr[rd]

	//创建创世区块
	var firstBlock Block
	firstBlock.BPM=100
	firstBlock.PrevHash="0"
	firstBlock.TimeStamp=time.Now().String()
	firstBlock.Validator="abc123"
	firstBlock.Index=1
	firstBlock.HashCode=GenerateHashValue(firstBlock)

	// 将区块加到区块链中
	Blockchain=append(Blockchain,firstBlock)

	//第二个区块
	var scondBlock=CenerateNextBlock(firstBlock,200,adds)
	Blockchain =append(Blockchain,scondBlock)

	fmt.Println(Blockchain)


	}

