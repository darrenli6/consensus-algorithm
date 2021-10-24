package main

import (
	"crypto/sha256"
	"encoding/hex"
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