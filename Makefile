PROTO_FILES=$(shell find ./proto/template7 -name *.proto)

.PHONY: proto
proto:
	protoc --proto_path=. \
 	       --go_out=paths=source_relative:./gen \
 	       --go-grpc_out=paths=source_relative:./gen \
	       --openapi_out=fq_schema_naming=true,default_response=false:. \
	       $(PROTO_FILES)
