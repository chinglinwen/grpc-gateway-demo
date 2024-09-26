# api demo

## run server

```
make build && ./demo 
```

## run grpc client

```
go run cmd/client/main.go 
2024/09/22 23:41:56 servers: [name:"s1" name:"s2"]
```

## run http client

```
$ curl  http://localhost:10000/v1/servers
{"servers":[{"name":"s1", "theme":""}, {"name":"s2", "theme":""}], "nextPageToken":""}
```

## install

```bash
wget https://github.com/protocolbuffers/protobuf/releases/download/v26.1/protoc-26.1-linux-x86_64.zip

	#https://grpc.io/docs/protoc-installation/
	unzip protoc-26.1-linux-x86_64.zip -d $HOME/.local
	echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
	source ~/.bashrc
	protoc --version

git clone --depth=1 https://github.com/googleapis/googleapis.git ~/.local/googleapis

#https://grpc.io/docs/languages/go/quickstart/
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest

go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
export PATH="$PATH:$(go env GOPATH)/bin"
```