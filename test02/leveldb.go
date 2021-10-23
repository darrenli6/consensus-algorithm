package test02

import "fmt"

type DB struct {
	path string
	data map[string][]byte


}
// 模拟开启
func New(path string) (*DB,error){
	self:=DB{
		path:path,
		data:make(map[string][]byte),
	}
	return &self,nil
}
// 模拟关闭
func (self *DB) Close() error{
	return nil
}

func (self *DB) Put(key []byte,value []byte) error{
	self.data[string(key)]=value
	return nil
}

// get
func (self *DB) Get(key []byte) ([]byte,error){
   if v,ok:=self.data[string(key)];ok{
   	  return v,nil
   }else{
   	return nil,fmt.Errorf("NotFound")
   }
}


func (self *DB) Delete(key []byte) (error){
	if _,ok:=self.data[string(key)];ok{
		delete(self.data,string(key))
		return nil
	}else{
		return fmt.Errorf("NotFound")
	}
}


// 模拟遍历
func (self *DB) Iterator() Iterator{
	return NewDefaultIterator(self.data)
}