package gopkg

import (
	"encoding/base64"
	"encoding/hex"
	"testing"

	"gotest.tools/assert"
)

func TestCryptoAES(t *testing.T) {
	data := []byte("Hello World")     // 待加密的数据
	key := []byte("ABCDEFGHIJKLMNOP") // 加密的密钥
	t.Log("原文：", string(data))

	t.Log("------------------ CBC模式 --------------------")
	encrypted := AESEncryptCBC(data, key)
	t.Log("密文(hex)：", hex.EncodeToString(encrypted))
	t.Log("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted := AESDecryptCBC(encrypted, key)
	assert.Equal(t, decrypted, data)

	t.Log("------------------ ECB模式 --------------------")
	encrypted = AESEncryptECB(data, key)
	t.Log("密文(hex)：", hex.EncodeToString(encrypted))
	t.Log("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted = AESDecryptECB(encrypted, key)
	assert.Equal(t, decrypted, data)

	t.Log("------------------ CFB模式 --------------------")
	encrypted = AESEncryptCFB(data, key)
	t.Log("密文(hex)：", hex.EncodeToString(encrypted))
	t.Log("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted = AESDecryptCFB(encrypted, key)
	assert.Equal(t, decrypted, data)

	t.Log("------------------ OFB模式 --------------------")
	encrypted = AESEncryptOFB(data, key)
	t.Log("密文(hex)：", hex.EncodeToString(encrypted))
	t.Log("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted = AESDecryptOFB(encrypted, key)
	assert.Equal(t, decrypted, data)

	t.Log("------------------ GCM模式 --------------------")
	encrypted, _ = AESEncryptGCM(data, key)
	t.Log("密文(hex)：", hex.EncodeToString(encrypted))
	t.Log("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted, _ = AESDecryptGCM(encrypted, key)
	assert.Equal(t, decrypted, data)

}

func TestBase64(t *testing.T) {
	data := []byte("Hello World")     // 待加密的数据
	key := []byte("ABCDEFGHIJKLMNOP") // 加密的密钥

	gotBase := encodeBase64(data)
	got, _ := decodeBase64(gotBase)
	assert.Equal(t, got, data)

	encrypted, _ := AESEncryptGCMBase64String(string(data), string(key))
	gotData, _ := AESDecryptGCMBase64String(encrypted, string(key))
	assert.Equal(t, gotData, data)

}
