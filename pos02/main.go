package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//pos思维 按照币的数量 你分配记账权的比例



// 定义全节点

type PNode struct {
	//持有币的数量
	Tokens int
	// 币龄
	Days int
	//地址
	Address  string

}

// Pblock

type PBlock struct {

	Index int
	Data string
	PreHash string
	Hash string
	Timestamp string
	Validator *PNode


}

// 生成创世区块
func firstBlock() PBlock{

	//创建区块
	var firstBlock =PBlock{0,"创世区块","","",
		time.Now().String(),&PNode{0,0,""}}

    firstBlock.Hash=hex.EncodeToString(BlockHash(&firstBlock))
    return firstBlock

}

// 计算hash的方法
func BlockHash(block *PBlock) []byte{
	hashed:=strconv.Itoa(block.Index)+block.Data+
		block.PreHash+block.Timestamp+block.Validator.Address

	h:=sha256.New()
	h.Write([]byte(hashed))
	hash:=h.Sum(nil)
	return hash
}
//创建5个全节点
var nodes=make([]PNode,5)
// 存放节点的地址
var addr=make([]*PNode,15)

func InitNodes(){
	nodes[0]=PNode{1,1,"0x12341"}
	nodes[1]=PNode{2,1,"0x12342"}
	nodes[2]=PNode{3,1,"0x12343"}
	nodes[3]=PNode{4,1,"0x12344"}
	nodes[4]=PNode{5,1,"0x12345"}

	count:=0
	for i:=0;i<len(nodes);i++{
		// 持币数量
		for j:=0;j<nodes[i].Tokens*nodes[i].Days ;j++{
			addr[count]=&nodes[i]
			count++
		}
	}

	fmt.Println("节点 tokens days Address \n ")
	fmt.Printf("%v \n",nodes)
	fmt.Println("生产者列表")
	for i:=0;i<len(addr);i++{
		fmt.Printf(" %v \n",addr[i].Address)
	}
	fmt.Println()

}


//创建新节点 实现pos算法
func CreateNewBlock(lastBlock *PBlock,data string) PBlock{

	var newBlock PBlock
	newBlock.Index=lastBlock.Index
	newBlock.Timestamp=time.Now().String()
	newBlock.Data=data
	newBlock.PreHash=lastBlock.Hash

	// 需要休眠一下
	time.Sleep(100000000)

	rand.Seed(time.Now().Unix())
	// 产生0到15的随机数
	var rd=rand.Intn(15)
	//选择出矿工
	node:=addr[rd]
	fmt.Printf("  由 %s 根据pos算法产生了区块\n",node.Address)

	//验证者  实际的挖矿节点
 	newBlock.Validator=node
 	// 模拟获取奖励
 	node.Tokens+=1
 	newBlock.Hash=hex.EncodeToString(BlockHash(&newBlock))

 	return newBlock



}

func main(){
	InitNodes()
	//创建创世区块
	var firstBlock= firstBlock()

	for i:=0;i<30;i++{
		var newBlock=CreateNewBlock(&firstBlock,"新区块")
		fmt.Println("新区块信息")
		fmt.Printf("hash %s\n",newBlock.Hash)
	}
}