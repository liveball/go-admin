# go-admin
后台管理系统

### 运行模式
gopath & go mod

### 运行命令

```shell

 go run ./cmd/main.go -conf ./cmd/admin.toml 
 
```

```curl
curl -X POST \
  http://127.0.0.1:8080/x/admin/user/add \
  -H 'Accept: */*' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Host: 127.0.0.1:8080' \
  -H 'Postman-Token: 72bf5c7c-5d0a-4ed8-a88c-0b31d8df4912,316aaef1-86d6-4c7f-afea-edc03394c775' \
  -H 'User-Agent: PostmanRuntime/7.15.0' \
  -H 'accept-encoding: gzip, deflate' \
  -H 'cache-control: no-cache' \
  -H 'content-length: 270' \
  -H 'content-type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW' \
  -F name=test \
  -F nick=sdsd
```