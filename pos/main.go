package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
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
	newBlock.BPM=oldBlock.BPM
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

func main(){

	var firstBlock Block
	myBlock :=CenerateNextBlock(firstBlock,1,"abc")
	fmt.Println(myBlock)
}