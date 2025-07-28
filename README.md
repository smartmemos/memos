# smartmemos

<img height="56px" src="https://smart-memos.zeropkg.com/logo.png" alt="Memos" />

An open-source, lightweight note-taking solution.

### 安装

#### golang 依赖

```bash

#  buf 工具
go install github.com/bufbuild/buf/cmd/buf@latest

# grpcurl工具
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

# protobuf to golang
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# connectrpc golang
go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest

# 生成swager文档
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

```

#### node 依赖

```bash

# connectrpc nodejs
npm install -g @bufbuild/protoc-gen-es

```
