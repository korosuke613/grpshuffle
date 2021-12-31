# Makefile
OUTPUT=gen
NPM_BIN=$(shell npm bin)

GRPC_TOOL=$(NPM_BIN)/grpc_tools_node_protoc
TYPESCRIPT_PLUGIN=protoc-gen-ts=$(NPM_BIN)/protoc-gen-ts
INCLUDE_PROTO_DIR=..
PROTO_PATH=../*.proto

.PHONY: protogen
protogen:
	# generate js and .d.ts codes via grpc-tools
	$(GRPC_TOOL) \
    	--plugin=${TYPESCRIPT_PLUGIN} \
    	--js_out=import_style=commonjs,binary:$(OUTPUT) \
    	--grpc_out=grpc_js:$(OUTPUT) \
    	--ts_out=grpc_js:$(OUTPUT) \
    	-I $(INCLUDE_PROTO_DIR) \
    	$(PROTO_PATH)

clean:
	rm -rf $(OUTPUT) && mkdir -p $(OUTPUT)
