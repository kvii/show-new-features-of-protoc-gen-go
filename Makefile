.PHONY: init
# 下载 protoc 相关插件。注意：protoc-gen-go 会下载最新版本。
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest

.PHONY: pb
pb:
	protoc --proto_path=./pb \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./pb \
		   --openapi_out=fq_schema_naming=true,default_response=false:. \
			pb/*.proto