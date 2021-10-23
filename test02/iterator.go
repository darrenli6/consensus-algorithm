package test02

import "fmt"

//定义
type Iterator interface {

	Next() bool
	Key() []byte
	Value() []byte
	Close() error

}

// 定义一个键值对结构体

type Pair struct {

	Key []byte
	Value []byte

}

// 定义一个迭代器
type DefaultIterator struct {
	data[] Pair
	index int
	length int
}

func NewDefaultIterator(data map[string][]byte) *DefaultIterator{
	//创建默认迭代器
	self:=new(DefaultIterator)
	self.index=-1
	self.length=len(data)
	for k,v:=range data{
		p:=Pair{
			Key:[]byte(k),
			Value:v,
		}
		self.data=append(self.data,p)
	}
	return self
}

// 是否存在下一个值
func (self *DefaultIterator) Next() bool{
	if self.index < self.length-1{
	    self.index++
		return true
	}
	return false
}

// 是否存在下一个值
func (self *DefaultIterator) Key() []byte{
	if self.index == -1 || self.index>=self.length{
		 panic(fmt.Errorf("越界"))
	}
	return self.data[self.index].Key

}
// 是否存在下一个值
func (self *DefaultIterator) Value() []byte{
	if  self.index>=self.length{
		panic(fmt.Errorf("越界"))
	}
	return self.data[self.index].Value

}

func (self *DefaultIterator) Close() error{

	return nil
}