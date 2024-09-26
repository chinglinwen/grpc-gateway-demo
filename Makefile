
.PHONY: all
all: generate build

.PHONY: build
build:
	go build -o demo .

.PHONY: generate
generate: grpcgen genopenapiyaml

.PHONY: grpcgen
grpcgen:
	mkdir -p generated/grpcgen
	protoc -I ~/.local/googleapis \
	--go_out=generated/grpcgen \
	--go_opt=paths=source_relative \
    --go-grpc_out=generated/grpcgen \
	--go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=generated/grpcgen \
	--grpc-gateway_opt=paths=source_relative \
	--grpc-gateway_opt=generate_unbound_methods=true \
    --proto_path=protos \
	protos/server.proto

# .PHONY: genjs
# genjs:
# 	mkdir -p generated/jsgen
# 	protoc -I ~/.local/googleapis -I=protos \
# 	--js_out=import_style=commonjs:generated/jsgen \
# 	--grpc-web_out=import_style=commonjs,mode=grpcwebtext:generated/jsgen \
# 	server.proto \

# .PHONY: gents
# gents:
# 	mkdir -p generated/tsgen
# 	protoc -I ~/.local/googleapis -I=protos \
# 	--js_out=import_style=commonjs:generated/tsgen \
# 	--grpc-web_out=import_style=typescript,mode=grpcwebtext:generated/tsgen \
# 	server.proto \

.PHONY: genopenapiyaml
genopenapiyaml:
	protoc -I ~/.local/googleapis -I=protos --openapi_out=protos server.proto 

