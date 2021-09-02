# Crypto
1. DES (decprecated)
1. AES (128/256 bit)
1. RSA

## AES crypto
[[go-exmaple][aes-go]]

AES PROP | VALUES
---|---
KEY | 128 256 bits  
MODE | CBC CFB  
IV | 16bytes  
PADDING | PKCS5 PKCS7 NOPADDING


```go
import (
	"crypto/aes"
	"crypto/cipher"
)

// raw plain text
rawText := "你好,世界 ! for AES"

// key, 32 bytes, 256 bits -> AES256
aesKey := make([]byte, 32)
rand.Read(aesKey)
// initial-vector, 16 bytes, 128 bits
commonIV := make([]byte, 16)
rand.Read(commonIV)

// pick AES based on key length
c, err := aes.NewCipher([]byte(aesKey))
if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
}

rawBytes := []byte(rawText)
// encode
// raw(string) => raw([]byte) => encoded([]byte) => hex(string)
cfb := cipher.NewCFBEncrypter(c, []byte(commonIV))
encodedBytes := make([]byte, len(rawBytes))
cfb.XORKeyStream(encodedBytes, rawBytes)
encodedString := hex.EncodeToString(encodedBytes)
fmt.Printf("ENCODE[AES-CFB]: raw: %v => %v\n", rawText, encodedString)

// decode
// hex(string) => decoded([]byte) => raw([]byte) => text(string)
cfbDec := cipher.NewCFBDecrypter(c, commonIV)
decodedBytes := make([]byte, len(encodedString))
bytesToDecode, _ := hex.DecodeString(encodedString)
cfbDec.XORKeyStream(decodedBytes, bytesToDecode)
fmt.Printf("DECODE[AES-CFB]: raw: %v => %v\n", encodedString, string(decodedBytes))
```

[aes-go]: https://willh.gitbook.io/build-web-application-with-golang-zhtw/09.0/09.6