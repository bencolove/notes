# Convert between `[]byte` and `string`
* Base64
* Hex

## Base64
```go
import (
	"encoding/base64"
)
// demostration of BASE64
// raw text
rawText := "你好, 世界"

// encode
// NOT base64.StdEncoding
b64EncodedString := base64.URLEncoding.EncodeToString([]byte(rawText))
fmt.Printf("Raw: %v, encoded: %v\n", rawText, b64EncodedString)

// decode
if decodedBytes, err := base64.URLEncoding.DecodeString(b64EncodedString); err != nil {
    fmt.Printf("Error: %v\n", err)
} else {
    //
    fmt.Printf("DECODE: raw: %v, decoded: %v\n", b64EncodedString, string(decodedBytes))
}
```

## Hex String
```go
import (
    "crypto/rand"
)

buf := make([]byte, 16)
// random 16 bytes 
_, _ = rand.Read(buf)
// about 5X faster than the fmt.Printf
hexText := hex.EncodeToString(buf)
// slower
fmt.Printf("fmt: %x\n", buf)

hexBytes, _ := hex.DecodeString(hexText)
fmt.Printf("fmt: %x\n", hexBytes)
```
