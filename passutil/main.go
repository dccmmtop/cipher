package passutil

import (
	"bytes"
	"cipher/logger"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	mrand "math/rand"
	"os"
	"strings"
	"time"
)
var (
	KEY_SEP = "|+d@c#m$m+|"
	PASS_SEP = "-|=w-a-n=|-"
	PRIVATE_KEY = "./私钥/private.pem"
	PUBLIC_KEY = "./公钥/public.pem"
)


//RSA公钥私钥产生
func GenRsaKey()  {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	prvkey := pem.EncodeToMemory(block)
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		panic(err)
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	pubkey := pem.EncodeToMemory(block)
	os.Mkdir(strings.Split(PUBLIC_KEY,"/")[1],os.ModePerm)
	os.Mkdir(strings.Split(PRIVATE_KEY,"/")[1],os.ModePerm)
	ioutil.WriteFile(PUBLIC_KEY, pubkey, 0777)
	ioutil.WriteFile(PRIVATE_KEY, prvkey, 0777)
}

//签名
func RsaSignWithSha256(data []byte, keyBytes []byte) []byte {
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("private key error"))
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("ParsePKCS8PrivateKey err", err)
		panic(err)
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		fmt.Printf("Error from signing: %s\n", err)
		panic(err)
	}

	return signature
}

//验证
func RsaVerySignWithSha256(data, signData, keyBytes []byte) bool {
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("public key error"))
	}
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	hashed := sha256.Sum256(data)
	err = rsa.VerifyPKCS1v15(pubKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], signData)
	if err != nil {
		panic(err)
	}
	return true
}

// 公钥加密
func RsaEncrypt(data, keyBytes []byte) []byte {
	//解密pem格式的公钥
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("public key error"))
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, data)
	if err != nil {
		panic(err)
	}
	return ciphertext
}

// 私钥解密
func RsaDecrypt(ciphertext, keyBytes []byte) []byte {
	//获取私钥
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		CheckError(nil, "私钥读取错误")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		CheckError(err, "私钥解析错误")
	}
	// 解密
	data, err := rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
	if err != nil {
		CheckError(err, "解密失败")
		return nil
	}
	return data
}

func DecryptByRsa(cipherText string, privateKey []byte) string {
	cipherBytes, err := hex.DecodeString(cipherText)
	CheckError(err, "解密错误")
	return string(RsaDecrypt(cipherBytes, privateKey))

}

func CheckError(err error, msg string) {
	if err != nil {
		logger.Logger.Error("%s, 错误信息: %v\n", msg, err)
	}

}

//加密过程：
//  1、处理数据，对数据进行填充，采用PKCS7（当密钥长度不够时，缺几位补几个几）的方式。
//  2、对数据进行加密，采用AES加密方法中CBC加密模式
//  3、对得到的加密数据，进行base64加密，得到字符串
// 解密过程相反

//16,24,32位字符串的话，分别对应AES-128，AES-192，AES-256 加密方法
//key不能泄露
// var PwdKey = []byte("ABCDABCDABCDABCD")

//pkcs7Padding 填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	//判断缺少几位长度。最少1，最多 blockSize
	padding := blockSize - len(data)%blockSize
	//补足位数。把切片[]byte{byte(padding)}复制padding个
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

//pkcs7UnPadding 填充的反向操作
func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	}
	//获取填充的个数
	unPadding := int(data[length-1])
	return data[:(length - unPadding)], nil
}

//AesEncrypt 加密
func AesEncrypt(data []byte, key []byte) ([]byte, error) {
	//创建加密实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//判断加密快的大小
	blockSize := block.BlockSize()
	//填充
	encryptBytes := pkcs7Padding(data, blockSize)
	//初始化加密数据接收切片
	crypted := make([]byte, len(encryptBytes))
	//使用cbc加密模式
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	//执行加密
	blockMode.CryptBlocks(crypted, encryptBytes)
	return crypted, nil
}

//AesDecrypt 解密
func AesDecrypt(data []byte, key []byte) ([]byte, error) {
	//创建实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取块的大小
	blockSize := block.BlockSize()
	//使用cbc
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	//初始化解密数据接收切片
	crypted := make([]byte, len(data))
	//执行解密
	blockMode.CryptBlocks(crypted, data)
	//去除填充
	crypted, err = pkcs7UnPadding(crypted)
	if err != nil {
		return nil, err
	}
	return crypted, nil
}

//EncryptByAes Aes加密 后 base64 再加
func EncryptByAes(data []byte, key string) (string, error) {
	res, err := AesEncrypt(data, []byte(key))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(res), nil
}

//DecryptByAes Aes 解密
func DecryptByAes(data string, key string) string {
	dataByte, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		CheckError(err, "")
	}

	text, err := AesDecrypt(dataByte, []byte(key))
	CheckError(err, "")
	return string(text)
}

func Encrypt(content string, pubKeyPathList []string) string {
	pubKeySize := len(pubKeyPathList)
	if len(content) == 0 || pubKeySize == 0{
		return ""
	}

	desKey := randKey()
	// 不同用户的公钥分别加密该key
	var  desKeyRsaHexList = make([]string, 0, pubKeySize)
	for i :=0; i < pubKeySize; i ++ {
		pubKey, err := ioutil.ReadFile(pubKeyPathList[i])
		CheckError(err, "读取公钥信息错误")
		// 加密key
		desKeyRsa := RsaEncrypt([]byte(desKey), pubKey)
		desKeyRsaHexList = append(desKeyRsaHexList, hex.EncodeToString(desKeyRsa))
	}

	// 对原文加密
	secretText, err := EncryptByAes([]byte(content), desKey)
	CheckError(err, "des 加密错误")
	// 拼接加密的desKey,和密文
	result := strings.Join(desKeyRsaHexList,KEY_SEP) + PASS_SEP + secretText
	logger.Logger.Debug("密文\n",result)
	return result
}

func Decrypt(str string) (string, error) {
	if len(str) == 0 {
		logger.Logger.Error("密文为空")
		return "",errors.New("密文为空！")
	}
	ciphertextList := strings.Split(str, PASS_SEP)
	size := len(ciphertextList)
	if size < 2 {
		logger.Logger.Error("密文格式错误")
		return "",errors.New("解密失败")
	}

	prv, err := ioutil.ReadFile(PRIVATE_KEY)
	if err != nil {
		logger.Logger.Error("读取私钥错误: ", err)
		return  "",errors.New("解密失败，无法读取私钥!")
	}
	// 解密使用自己公钥加密的key,循环尝试解密
	// FIXME: 改进
	keys := strings.Split(ciphertextList[0],KEY_SEP)
	keySize := len(keys)
	var selfKey string
	for i := 0; i< keySize; i ++ {
		selfKey = DecryptByRsa(keys[i], prv)
		if selfKey != "" {
			break
		}
	}
	if selfKey == "" {
		logger.Logger.Debug("本人无法解密")
		return "",errors.New("无法解密")
	}
	text := DecryptByAes(ciphertextList[1], selfKey)
	return text,nil
}
func randKey() string {
	mrand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, 32)
	for i := range b {
		b[i] = letterRunes[mrand.Intn(len(letterRunes))]
	}
	return string(b)
}
