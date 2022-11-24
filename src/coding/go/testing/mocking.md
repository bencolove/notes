# `gomock`

- `github.com/golang/mock/gomock`  
- `github.com/golang/mock/mockgen`

---

## mock an interface

```go
type DB interface {
    Get(key string) (any, error)
}
```

function to be tested, dependent on interface other than concrete structure.
```go
func GetFromDB(db DB, key string) any {
    if v, err := db.Get(key); err == nil {
        return v
    }

    return nil
}
```

>Generate mock file  

`$ mockgen -source=db.go -destination=db_mock.go -package=main`


>Test it with mocked object.
```go
func TestGetFromDB(t *testing.T) {
    ctrl := gomock.NewController(t)
    // assert on invocation of DB.Get()
    defer ctrl.Finish()

    // autogen by gomock from db_mock.go
    m := NewMockDB(ctrl)

    m.EXPECT().Get(gomock.Eq("Tom")).Return(100, errors.New("not exist"))

    if v := GetFromDB(m, "Tom"); v != -1 {
        t.Fatalf("expected %v, but got %v", -1, v)
    }

}
```
> stubs (set of params, returns, erros...)

```go
// case sets
// When_Tom.ReturnAndError
m.EXPECT().Get(gomock.Eq("Tom")).Return(0, errors.New("not exist"))
// When_Any.Return
m.EXPECT().Get(gomock.Any()).Return(630, nil)
// When_NotSam.DoNoReturn
m.EXPECT().Get(gomock.Not("Sam")).Do(func(key string) {
    print key
}) 
// When_Nil.DoAndReturn
m.EXPECT().Get(gomock.Nil()).DoAndReturn(func (key string) { return 0, errors.New("nil")}) 


// setup numer of invocations
m.EXPECT().Get().Return().Times(2)
call it twice

// setup invocation order
m1 := m.EXPECT()...
m2 := m.EXPECT()...
gomock.InOrder(m2, m1)
call it twice
```
