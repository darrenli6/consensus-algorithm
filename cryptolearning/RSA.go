package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
)

func RsaGenKey(bits int) error{

    // GenerateKey是使用随机数生成生成一对公钥和私钥

	privKey,err:= rsa.GenerateKey(rand.Reader,bits)

	if err!=nil{
		panic(err)
	}
    //x509 通用证书格式 序列号 签名算法 颁发者 有效时间 持有者 公钥
	// PKCS RSA 实验室与其他安全系统开发商 为了促进公钥密码的发展而做出一系列的标准
	x509.MarshalPKCS1PrivateKey(privKey)


}
