# Build Golang Docker Image
1. official
2. custom

## Official Way
Pre-requistes:
1. Go version >= 1.16
1. docker

> 1. Dockerfile  
```Dockerfile
FROM golang:1.16-alpine

# create dir inside image
# used as cwd
WORKDIR /app

# copy go.mod and go.sum from current dir where Dockerfile resides into IMAGE:cwd 
COPY go.mod .
COPY go.sum .

# download modules needed inside IMAGE
RUN go mod download

COPY *.go .

RUN go build -o ./executable

CMD [ "./executable" ]
```

> 2. build Image and tag  
`$ docker build --tag docker-go .`

> 3. view local Images  
`$ docker images`

> 4. more tags  
`$ docker tag executable:latest executable:v1.0`

## Multistage Build
```Dockerfile
##
## Build
##

FROM golang:1.16 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY *.go .

RUN go build -o /executable

##
## Deploy
##

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /executable /executable

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/executable"]
```

`$ docker build -t executable:multistage -f Dockerfile .`

## Other exmaple  
-- _`Dockerfile`_ --
```Dockerfile
FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" .


FROM scratch

WORKDIR /app

COPY --from=builder /app/executable /user/bin

ENTRYPOINT ["executable"]
```

>NOTE  
1. [[CGO_ENABLED][CGO-linking]]=0: disalbe CGO, use static linking
1. -ldflags: pass along to `go tool link`
1. -w: disable DWARF generation (for debugging)
1. -s: disable symbol table 

> info (note the SIZE column)  
`$ docker images | grep executable -B 1`

[build-go-image]: https://docs.docker.com/language/golang/build-images/
[small-go-image]: https://blog.wu-boy.com/2017/04/build-minimal-docker-container-using-multi-stage-for-go-app/
[CGO-linking]: https://johng.cn/cgo-enabled-affect-go-static-compile/