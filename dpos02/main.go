package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//定义全节点

type Node struct {
	Name string
	Votes int
}

//区块

type Block struct {
	Index int
	TimeStamp string
	Prehash string
	Hash string
	Data []byte
	delegate *Node
}

func firstBlock() Block{
	gene:=Block{0,time.Now().String(),"","",[]byte("first blokc"),nil}

	gene.Hash=string(blockHash(gene))
	return gene

}
// 计算hash
func blockHash(block Block) []byte{
	hash:=strconv.Itoa(block.Index)+block.TimeStamp+block.Prehash+hex.EncodeToString(block.Data)

	h:=sha256.New()
	h.Write([]byte(hash))
	hashed:=h.Sum(nil)
	return hashed
}

//生成新的区块
func (node *Node)GenerateNewBlock(lastBlock Block,data []byte) Block{
	var newBlock=Block{lastBlock.Index+1,time.Now().String(),
		lastBlock.Hash,"",data,nil}
	newBlock.Hash=hex.EncodeToString(blockHash(newBlock))
	newBlock.delegate=node
	return newBlock

}


// 创建10个节点
var NodeAddr=make([]Node,10)

// 创建节点
func CreateNode(){
	for i:=0;i<10;i++{
		name:=fmt.Sprintf("节点 %d 票数",i)
		NodeAddr[i]=Node{name,0}
	}
}

// 简单模拟一下投票
func Vote(){

	for i:=0;i<10;i++{
		rand.Seed(time.Now().UnixNano())
		time.Sleep(100000)
		vote:=rand.Intn(10000)
	     // 为10个节点进行投票
		 // 每个节点的票数
		 NodeAddr[i].Votes=vote
		 fmt.Printf("节点 [%d] 票数[%d] \n",i,vote)
	}
}

// 选出票数最多的前三名
func SortNodes()[] Node{
	// 10个节点
	n:=NodeAddr

	for i := 0; i < len(n)-1; i++ {
		for j := i+1; j < len(n); j++ {
			if  n[i].Votes<n[j].Votes {
				n[i],n[j] = n[j],n[i]
			}
		}
	}



	return n[:3]
}

func main() {


	CreateNode()
	fmt.Printf("创建节点\n")
	fmt.Println(NodeAddr)
	fmt.Print("节点票数\n")
	// 票数
	Vote()
	//选出前三
	nodes:=SortNodes()

	fmt.Printf("中奖者是 \n ")

	fmt.Println(nodes)
	//创世区块
	first:=firstBlock()

	lastBlock:=first

	fmt.Println("开始生成区块")
	for i:=0;i<len(nodes);i++{
		fmt.Printf(" [%s %d] 生成新的区块\n  ",nodes[i].Name,nodes[i].Votes)
		lastBlock=nodes[i].GenerateNewBlock(lastBlock,[]byte(fmt.Sprintf("new Block %d",i)))
	}


}