package test02

import (
	"fmt"
	"testing"
)

// 测试迭代器

func Test_leveldb(t *testing.T) {
	 var err error


	 //建立连接
	 db,err:=New("")
	 check(err)
	 //put
	 err=db.Put([]byte("k1"),[]byte("v1"))
	 check(err)
	err=db.Put([]byte("k2"),[]byte("v2"))
	check(err)
	err=db.Put([]byte("k3"),[]byte("v3"))
	check(err)
	err=db.Put([]byte("k4"),[]byte("v4"))
	check(err)

	_,err=db.Get([]byte("k1"))
	check(err)
	//
	err=db.Delete([]byte("k1"))

	//_,err=db.Get([]byte("k1"))
	//check(err)

	iter:=db.Iterator()

	for iter.Next()  {
		fmt.Printf("%s :%s \n",iter.Key(),string(iter.Value()))
	}
}

func check(err error){
	if err!=nil{
		panic(err)
	}
}
