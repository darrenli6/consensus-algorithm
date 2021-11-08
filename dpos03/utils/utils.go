package utils

import (
	"fmt"
	"github.com/godotenv/godotenv-master"
	"strconv"

	"os"
	"log"
	"crypto/sha256"
	"encoding/hex"
	"bytes"
	"encoding/gob"
)



type UUID [16]byte

//保留两位小数浮点数
func Decimal(value float64) float64 {

	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

//加载配置信息
func LoadingEnv(fileName string) {
	err := godotenv.Load(fileName)
	if err != nil {
		log.Fatal(err)
	}
}

//根据key查找.env文件中的值
func GetEnvValue(key string) string {
	return os.Getenv(key)
}

//计算哈希
func CalculateHash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

//将string转成固定长度数组
func ConvertStrToBytes(str string) []byte {
	var bytes [constLength]byte
	for i, c := range str {
		bytes[i] = byte(c)
	}
	return bytes[:]
}

//序列化对象
func Serialize(data interface{}) []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(data)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}