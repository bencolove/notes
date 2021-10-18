# `net/http` Timeouts

![][server_timeouts]
1. Server timeouts:
    1. `http.Server.ReadTimeout` -> `net.Conn.SetReadDeadline`
    1. `http.Server.WriteTimeout` -> `net.Conn.SetWriteDeadline`

```go
srv := &http.Server{
    ReadTimeout: 5 * time.Second,
    WriteTimeout: 10 * time.Second,
}
log.Println(srv.ListenAndServe())
```

---

![][client_timeout]
1. Client timeouts:
    1. `http.Client.Timeout`

>By `context.Context`  
```go
// timeout for entire client
c := &http.Client{
    Timeout: 15 * time.Second,
}

// timeout for each request
// timer timeout pattern
ctx, cancel := context.WithCancel(context.Background())

TIME_OUT := 5*time.Second
timer := time.AfterFunc(TIME_OUT, func(){
    // timeout to cancel
    cancel()
})

req, err := http.NewRequest(http.GetMethod, URL, nil)
req = req.WithContext(ctx)

// or
req, err := http.NewRequestWithContext(ctx, http.GetMethod, URL, nil)

if err != nil {
    // ...
}
```

>Fine granluar controll  
```go
c := &http.Client{
    Transport: &http.Transport{
        Dial: (&net.Dialer{
                Timeout:   30 * time.Second,
                KeepAlive: 30 * time.Second,
        }).Dial,
        TLSHandshakeTimeout:   10 * time.Second,
        ResponseHeaderTimeout: 10 * time.Second,
        ExpectContinueTimeout: 1 * time.Second,
    }
}
```


[explain]: https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
[server_timeouts]: https://blog.cloudflare.com/content/images/2016/06/Timeouts-001.png
[client_timeout]: https://blog.cloudflare.com/content/images/2016/06/Timeouts-002.png