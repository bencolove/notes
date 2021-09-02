# En/decode Charset
1. [[transform encoding][text-transform]]
1. [[determine HTML encoding][determine-html-encoding]]

## Transform Encoding
By [[golang.org/x/text][text-transform]]

```go
import (
    "bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
    "golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

// from UTF-8 --encode--> Big5 
// 測試 --encode--> "\xb4\xfa\xb8\xd5"
func encodeBig5String(utf8String string) (string, error) {
	utf8ToBig5 := traditionalchinese.Big5.NewEncoder()
	if big5String, _, err := transform.String(utf8ToBig5, utf8String); err != nil {
		return "", err
	} else {
        fmt.Printf("Big5: %q\n", big5String)
		return big5String, nil
	}
}

// from Big5  --decode--> UTF-8 
// "\xb4\xfa\xb8\xd5" --decode--> 測試
func decodeUTF8String(big5Text string) (string, error) {
	big5ToUtf8 := traditionalchinese.Big5.NewDecoder()
	if utf8String, _, err := transform.String(big5ToUtf8, big5Text); err != nil {
		return "", nil
	} else {
		fmt.Printf("UTF8: %v\n", utf8String)
		return utf8String, nil
	}
}


// Big5 -> UTF8
// big5ToUTF8
func decodeBig5(s []byte) ([]byte, error) {
	// []byte -> io.Reader
	// io.Reader -> bufio(buffered.io.Reader) bufio.NewReader(r)
	bin := bytes.NewReader(s)
	// Big5ToUTF8
	big5ToUTF8Reader := transform.NewReader(bin, traditionalchinese.Big5.NewDecoder())
	// read all
	if utf8Bytes, err := ioutil.ReadAll(big5ToUTF8Reader); err != nil {
		return nil, err
	} else {
		return utf8Bytes, nil
	}
}

// UTF8 -> Big5
// utf8ToBig5
func encodeBig5(utf8Bytes []byte) ([]byte, error) {
	bin := bytes.NewReader(utf8Bytes)
	utf8ToBig5Reader := transform.NewReader(bin, traditionalchinese.Big5.NewEncoder())

	if big5Bytes, err := ioutil.ReadAll(utf8ToBig5Reader); err != nil {
		return nil, err
	} else {
		return big5Bytes, nil
	}
}

func testDecodeBig5() {
    // quoted string of bytes
	big5Bytes := []byte("\xb4\xfa\xb8\xd5")

	utf8Bytes, err := decodeBig5(big5Bytes)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n%s\n%v\n", utf8Bytes, utf8Bytes, string(utf8Bytes))
	}
}
```

## Determine HTML Encoding
By [[golang.org/x/net][determine-html-encoding]]

```go

func tryUTF8Reader(htmlBodyReader io.Reader) (io.Reader, error) {
    encoding, encodingName, certain, err := determineEncodingFromReade(htmlBodyReader)

	utf8Reader := transform.NewReader(htmlBodyReader, encoding.NewDecoder())

    return utf8Reader
}

func determineEncodingFromHTMLBody(reader io.Reader) (encoding.Encoding, string, bool, error) {
	if bufferReader, err := bufio.NewReader(reader).Peek(1024); err != nil {
		// github.com/pkg/errors err = errors.Wrap(err, "bufio.NewReader")
		return nil, "", false, err
	} else {
		// golang.org/x/net/html/charset
		encoding, encodingName, certain := charset.DetermineEncoding(bufferReader, "")
		return encoding, encodingName, certain, nil
	}

}
```

[text-transform]: https://openhome.cc/Gossip/Go/XText.html
[determine-html-encoding]: https://zwindr.blogspot.com/2019/11/go-goquery.html