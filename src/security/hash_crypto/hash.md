# HASH Algorithms
[[hash-example][hash-go]]
1. MD5
1. SHA
1. HMAC
1. Advanced HASH

HASH ALGO | OUTPUT LENGTH | BYTES
---|---|---
MD5 | 128 bits | 16 Bytes
SHA256 | 256 bits | 32 Bytes

## MD5
```go
import (
    "crypto/md5"
	"crypto/sha256"
)
// h io.Writer, via io.WriteString()
h := md5.New()
str := "原始密碼"
io.WriteString(h, str)
hashed := h.Sum(nil)
// 原始密碼 => de5ad5ae91ff3a3923e027e8c8821069, len()=16
fmt.Printf("%v => %x, len()=%d\n", str, hashed, len(hashed))
// or append directly
hh := md5.New()
hh.Write([]byte("原始密碼"))
fmt.Printf("%x\n", hh.Sum(nil))
// Sum() appends bytes to the start of result
hhh := md5.New()
hhh.Write([]byte("原始密碼"))
fmt.Printf("%x\n", hhh.Sum([]byte{0, 0}))
```

## SHA256
```go
import (
    "crypto/sha256"
)

str := "一二三四五六七八九十"
h := sha256.New()
h.Write([]byte(str))
hashed := h.Sum(nil)
// 一二三四五六七八九十 => c3c377ab78a968b1c417802c5d703356788951a84376639168beda68973b8bd1, len()=32
fmt.Printf("%v => %x, len()=%d\n", str, hashed, len(hashed))
```

## HMAC
Based on `MD5` or `SHA256`
```go
import (
    "crypto/hmac"
	"crypto/sha256"
)
str := "一二三四五六七八九十"
secret := "!@#$%^"
hash := hmac.New(sha256.New, []byte(secret))
io.WriteString(hash, str)
hashed := hash.Sum(nil)
// HMAC: 一二三四五六七八九十 => f9664968e593c0e7c2aaa2c6791df7307c3ee8910a21c1e5dcccd1820caf5dc9, len()=32
fmt.Printf("HMAC: %v => %x, len()=%d\n", str, hashed, len(hashed))
```

## Advanced HASH
Thoughts: hash(salt1 + username + salt2 + hash(passwd))

```go
passwd := "原始密碼"
username := "用戶賬號"

hash := md5.New()
io.WriteString(hash, passwd)
pwmd5 := hex.EncodeToString(hash.Sum(nil))
//
pwmd5 := fmt.Sprintf("%x", h.Sum(nil))

salt1 := "!@#"
salt2 := "$%^"

// salt1 + username + salt2 + hash(passwd)
io.WriteString(hash, salt1)
io.WriteString(hash, username)
io.WriteString(hash, salt2)
io.WriteString(hash, pwmd5)

hashHex := hex.EncodeToString(hash.Sum(nil))
```

[hash-go]: https://willh.gitbook.io/build-web-application-with-golang-zhtw/09.0/09.5