package common

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"

	"github.com/astaxie/beego"
)

func EncodeMessageMd5(msg string) string {
	h := md5.New()
	coding := base64.NewEncoding(beego.AppConfig.String("base64key"))
	h.Write([]byte(msg)) // 需要加密的字符串为 123456
	key := []byte(beego.AppConfig.String("md5key"))
	cipherStr := h.Sum([]byte(key))

	return coding.EncodeToString(cipherStr)
}

func getKey() []byte {
	strKey := "zheshiyonggolangxiedeyigecmsxitongjichugongneng"
	keyLen := len(strKey)
	if keyLen < 16 {
		panic("res key 长度不能小于16")
	}
	arrKey := []byte(strKey)
	if keyLen >= 32 {
		//取前32个字节
		return arrKey[:32]
	}
	if keyLen >= 24 {
		//取前24个字节
		return arrKey[:24]
	}
	//取前16个字节
	return arrKey[:16]
}

//加密字符串
func EncryptAes(strMesg string) string {
	key := getKey()
	var iv = []byte(key)[:aes.BlockSize]
	encrypted := make([]byte, len(strMesg))
	aesBlockEncrypter, err := aes.NewCipher(key)
	if err != nil {
		beego.Debug("AES 加密失败")
		return ""
	}
	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncrypter, iv)
	aesEncrypter.XORKeyStream(encrypted, []byte(strMesg))
	return base64.StdEncoding.EncodeToString(encrypted)
}

//解密字符串
func DecryptAes(srcstr string) (strDesc string, err error) {
	src, _ := base64.StdEncoding.DecodeString(srcstr)
	defer func() {
		//错误处理
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	key := getKey()
	var iv = []byte(key)[:aes.BlockSize]
	decrypted := make([]byte, len(src))
	var aesBlockDecrypter cipher.Block
	aesBlockDecrypter, err = aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	aesDecrypter := cipher.NewCFBDecrypter(aesBlockDecrypter, iv)
	aesDecrypter.XORKeyStream(decrypted, src)
	return string(decrypted), nil
}
