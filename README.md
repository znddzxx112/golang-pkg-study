### release
- go mod And gen main cmd:
```
docker run --rm -v ~/gopath:/gopath -v ~/workspace/znddzxx112/fortest:/workspace centos7_golang:1.12.1 /go/bin/go build -o /workspace/main /workspace/fortest.go 
```
- build app's image And run:
```
docker build -t fortest:latest .
docker run -d --net=host --name fortest_con fortest:latest
```

### dev or debug:
- run server:
```
docker run -it --net=host --name fortest_debug --rm -v ~/gopath:/gopath -v ~/workspace/znddzxx112/fortest:/workspace centos7_golang:1.12.1 /go/bin/go run /workspace/fortest.go
```

- build tcp client:
```
docker run -it --net=host --name fortest_debug --rm -v ~/gopath:/gopath -v ~/workspace/znddzxx112/fortest:/workspace centos7_golang:1.12.1 /go/bin/go build /workspace/tcpcli.go
```