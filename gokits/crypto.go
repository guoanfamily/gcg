package utils

import (
	"crypto/aes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// MD5String md5
func MD5String(b []byte) string {
	return fmt.Sprintf("%x", md5.Sum(b))
}

// AESEncryptData AESEncryptData
func AESEncryptData(keystr, src string) string {
	key := []byte(keystr)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	bs := block.BlockSize()
	d := []byte(src)
	d = padData(d, bs)
	r := make([]byte, len(d))
	r1 := r
	for len(d) > 0 {
		block.Encrypt(r1, d)
		d = d[bs:]
		r1 = r1[bs:]
	}
	s := hex.EncodeToString(r)

	return s
}

// AESDecryptData AESDecryptData
func AESDecryptData(keystr, hexSrc string) string {
	key := []byte(keystr)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	src, err := hex.DecodeString(hexSrc)
	if err != nil {
		panic(err)
	}
	bs := block.BlockSize()
	d := []byte(src)
	r := make([]byte, len(d))
	r1 := r
	for len(d) > 0 {
		block.Decrypt(r1, d)
		d = d[bs:]
		r1 = r1[bs:]
	}

	return string(removePad(r))
}

// HmacShal HmacShal
func HmacSha1(message, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha1.New, key)
	h.Write([]byte(message))
	return base64.URLEncoding.EncodeToString([]byte(hex.EncodeToString(h.Sum(nil))))
}

func padData(d []byte, bs int) []byte {
	padedSize := ((len(d) / bs) + 1) * bs
	pad := padedSize - len(d)
	for i := len(d); i < padedSize; i++ {
		d = append(d, byte(pad))
	}
	return d
}

func removePad(r []byte) []byte {
	l := len(r)
	last := int(r[l-1])
	pad := r[l-last : l]
	isPad := true
	for _, v := range pad {
		if int(v) != last {
			isPad = false
			break
		}
	}
	if !isPad {
		return r
	}
	return r[:l-last]
}
