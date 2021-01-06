# Basic Types
* array
* slice
* map
* string

---
## `[]byte` converted from/to `string`
Memory copy options (`copy`/`memmove`) will kick in when:
* string concatenation
* string([]byte) `runtime.slicebytetostring`
* []byte(string) `runtime.stringtoslicebyte`

