# Hex
---

## uint8 bytes
`[]byte{0x12,0x34,0x56,0x78}`

> uint8_bytes <<=>> hex_bytes
 ```go
 //`[]byte{0x12,0x34,0xab,0xef}` =>
 //`[]byte("1234abef")`
 encoding.hex.Encode(hex_bytes, int_bytes) numEncoded 
 
 //`[]byte("1234abef")` =>
 //`[]byte{0x12,0x34,0xab,0xef}`
 encoding.hex.Decode(int_bytes, hex_bytes) numDecoded

 ```

> uint8_bytes <<=>> int
```go
//`[]byte{0x12,0x34,0xab,0xef}` =>
//`int64(43907237)`
encoding.BigEndian.Uint64(int_bytes) int64
 

//`int64(43907237)` =>
//`[]byte{0x12,0x34,0xab,0xef}`
encoding.BigEndian.PutUint64(int_bytes, 4393409)
 
```

> hex_bytes <<=>> int
```go
//`int64(43907237)` =>
//`string("1234abef")`
strconv.FormatInt(43907237, 16)

//`string("1234abef")` =>
//`int64(43907237)`
strconv.ParseInt("1234abef", 16, 64) (int64, error)

```
 