# Golang String Internal

Golang 在内存中使用UTF-8表示unicode, 所以
string 'Go語言', []byte長度為8, 中文(unicode)使用3個byte表示

而類型rune(unit32別名)是完整的unicode(碼元code point)類型

unicode/utf8 處理rune和UTF-8之間的轉換操作, 參數都是[]byte(8bits as UTF-8)
unicode/utf16 處理rune和UTF-16之間的轉換操作, 參數則是[]uint16

擴充套件 golang.org/x/text 可以處理不同編碼的轉換，國際化，本地化:
UTF-8(GO internal) ---encode---> other encoding
other encoding --->decode---> UTF-8(GO internal)

Unicode (codepoint aka rune) is internally stored as UTF-8 format.
So,
* one ASCII char is stored one byte
* one rune like '中' unicode(code point) is stored 3 bytes

From GO's point of view:
* encode: UTF-8 to other charset
* decode: other charset to UTF-8