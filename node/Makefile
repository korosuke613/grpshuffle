# Makefile
OUTPUT=gen
NPM_BIN=$(shell npm bin)

GRPC_TOOL=$(NPM_BIN)/grpc_tools_node_protoc
TYPESCRIPT_PLUGIN=protoc-gen-ts=$(NPM_BIN)/protoc-gen-ts
INCLUDE_PROTO_DIR=..
PROTO_PATH=../*.proto

.PHONY: all
all: protogen

../health.proto:
	curl -fsL -o ../health.proto https://raw.githubusercontent.com/grpc/grpc/master/src/proto/grpc/health/v1/health.proto

.PHONY: protogen
protogen: ../health.proto
	mkdir -p gen
	# generate js and .d.ts codes via grpc-tools
	$(GRPC_TOOL) \
    	--plugin=${TYPESCRIPT_PLUGIN} \
    	--js_out=import_style=commonjs,binary:$(OUTPUT) \
    	--grpc_out=grpc_js:$(OUTPUT) \
    	--ts_out=grpc_js:$(OUTPUT) \
    	-I $(INCLUDE_PROTO_DIR) \
    	$(PROTO_PATH)

pack: protogen
	npm run compile
	mkdir -p dist/gen
	cp ./gen/* ./dist/gen
	npm pack

clean:
	rm -rf $(OUTPUT) && mkdir -p $(OUTPUT)
