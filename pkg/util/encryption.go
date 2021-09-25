package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// Encrypt text with the passphrase
func Encrypt(text string, passphrase string) string {
	salt := []byte("9YMwN4KeNSZ9bqc9uQvM") // salt must 20 chars

	key, iv := deriveKeyAndIv(passphrase, string(salt))

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	pad := pKCS7Padding([]byte(text), block.BlockSize())
	ecb := cipher.NewCBCEncrypter(block, []byte(iv))
	encrypted := make([]byte, len(pad))
	ecb.CryptBlocks(encrypted, pad)

	return base64.StdEncoding.EncodeToString([]byte("sAlTeD__" + string(salt) + string(encrypted)))
}

// Decrypt encrypted text with the passphrase
func Decrypt(encrypted string, passphrase string) string {
	ct, _ := base64.StdEncoding.DecodeString(encrypted)
	if len(ct) < 28 || string(ct[:8]) != "sAlTeD__" {
		return ""
	}

	salt := ct[8:28]
	ct = ct[28:]
	key, iv := deriveKeyAndIv(passphrase, string(salt))

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	cbc := cipher.NewCBCDecrypter(block, []byte(iv))
	dst := make([]byte, len(ct))
	cbc.CryptBlocks(dst, ct)

	return string(pKCS7Trimming(dst))
}

func pKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pKCS7Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

func deriveKeyAndIv(passphrase string, salt string) (string, string) {
	salted := ""
	dI := ""

	for len(salted) < 48 {
		md := md5.New()
		md.Write([]byte(dI + passphrase + salt))
		dM := md.Sum(nil)
		dI = string(dM[:16])
		salted = salted + dI
	}

	key := salted[0:32]
	iv := salted[32:48]

	return key, iv
}

// GenerateSignedKey used to generate signed key to validate order data upon creating order
func GenerateSignedKey(keys []string, secret string) (signedKey string) {
	combined := secret
	for _, v := range keys {
		if v != "" {
			combined = fmt.Sprintf("%s-%s", combined, v)
		}
	}
	hasher := md5.New()
	hasher.Write([]byte(combined))
	signedKey = hex.EncodeToString(hasher.Sum(nil))

	return
}
