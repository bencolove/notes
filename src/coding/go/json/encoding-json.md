# `encoding/json`
Pakage-level functions:
* `Marshal(v interface{}) ([]byte, error)`
* `MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)`
* `Unmarshal([]byte, v interface{}) error`
* `Valid([]byte) bool`

Basic usages:
* Unmarshal(parse) into struct
* Unmarshal(parse) into interface{}/map[string]interface{}
* Marshal(format) 
* Tag conventions

## Parse into Struct
When format of JSON is known in advance.

```go
type App struct {
    Id string `json:"id"`
    Title string `json:"title"`
}

data := []byte(`
    {
        "id": "k34rAT4",
        "title": "My Awesome App"
    }
`)

var app App
err := json.Unmarshal(data, &app)
```
## Tag Conventions:
1. Only exportable fields (name with first letter in uppercase) are feeded or
2. `json:"field_name"` to denote mapped JSON field
3. `omitempty` when formating (marshal/encode into JSON) will ignore field if zero-valued
4. `-` simply ignore both parsing and formating

### Tag `json:"-"`
```go
type Confidential struct {
	Id string `json:"id"`
	Password string `json:"-"`
}

func TestIgnore() {
	raw := []byte(`{
	   "id": "3xt8",
	   "password": "you will not see my password"
	}`)
	
	var data Confidential 
	err := json.Unmarshal(raw, &data)
	
	fmt.Printf("parse(unmarshal):\n  source: %s, \n  parsed: %+v, \n  err=%v \n\n", string(raw), data, err)
	
	var output = Confidential {"out", "abc"}
	encoded, err := json.MarshalIndent(output, "p", "\t")
	fmt.Printf("Format(marshal): \n  source: %+v, \n  formatted: %s, \n  err=%v \n\n", output, encoded, err)
}
```
>Output  
Parse(unmarshal):  
  source: '{
	   "id": "3xt8",
	   "password": "you will not see my password"
	}',  
  parsed: {Id:3xt8 Password:},  
  err=<nil>  
Format(marshal):  
  source: {Id:out Password:abc},  
  formatted: '{
p	"id": "out"
p}',  
  err=<nil>  




Interfaces:
* `Marshaler`
* `Unmarshaler`

> `Marshaler` and `Unmarshaler`    
```go
type Marshaler interface {
    MarshalJSON() ([]byte, error)
}
type Unmarshaler interface {
    UnmarshalJSON([]byte) error
}
```

