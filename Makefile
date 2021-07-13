# Tool versoins
PROTOC_VERSION = 3.17.3
PROTOC_GEN_DOC_VERSION = 1.4.1
PROTOC_GEN_GO_VERSION = 1.26.0
PROTOC_GEN_GO_GRPC_VERSION = 1.1.0
PLATFORM = osx

MODULE := $(shell awk '/^module / {print $$2}' go.mod)
PWD := $(shell pwd)
MKDIR_P = mkdir -p

DOC_DIR := $(PWD)/doc
PROTOC = $(PWD)/bin/protoc
PROTOC_GEN_DOC = $(PWD)/bin/protoc-gen-doc
PROTOC_GEN_GO = $(PWD)/bin/protoc-gen-go
PROTOC_GEN_GO_GRPC = $(PWD)/bin/protoc-gen-go-grpc
RUN_PROTOC = PATH=$(PWD)/bin:$$PATH $(PROTOC) -I$(PWD)/include -I.
DOC_MD = $(DOC_DIR)/grpshuffle.md
DOC_HTML = $(DOC_DIR)/index.html

.PHONY: all
all: doc server client

.PHONY: clean
clean:
	rm -f grpshuffle.md server client

.PHONY: fullclean
fullclean: clean
	rm -rf bin include go/grpshuffle

$(DOC_DIR):
	$(MKDIR_P) $(DOC_DIR)

$(PROTOC):
	curl -fsL -o /tmp/protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_VERSION)/protoc-$(PROTOC_VERSION)-$(PLATFORM)-x86_64.zip
	unzip /tmp/protoc.zip 'bin/*' 'include/*'
	rm -f /tmp/protoc.zip

$(PROTOC_GEN_DOC):
	mkdir -p bin
	GOBIN=$(PWD)/bin go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@v$(PROTOC_GEN_DOC_VERSION)

$(PROTOC_GEN_GO):
	mkdir -p bin
	GOBIN=$(PWD)/bin go install google.golang.org/protobuf/cmd/protoc-gen-go@v$(PROTOC_GEN_GO_VERSION)

$(PROTOC_GEN_GO_GRPC):
	mkdir -p bin
	GOBIN=$(PWD)/bin go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v$(PROTOC_GEN_GO_GRPC_VERSION)

# generate markdown specification
$(DOC_MD): grpshuffle.proto $(DOC_DIR) $(PROTOC) $(PROTOC_GEN_DOC)
	$(RUN_PROTOC) --doc_out=$(DOC_DIR) --doc_opt=markdown,$@ $<

# generate html specification
$(DOC_HTML): grpshuffle.proto $(DOC_DIR) $(PROTOC) $(PROTOC_GEN_DOC)
	$(RUN_PROTOC) --doc_out=$(DOC_DIR) $<

go/grpshuffle/grpshuffle.pb.go: grpshuffle.proto $(PROTOC) $(PROTOC_GEN_GO)
	$(RUN_PROTOC) --go_out=module=$(MODULE):. $<

go/grpshuffle/grpshuffle_grpc.pb.go: grpshuffle.proto $(PROTOC) $(PROTOC_GEN_GO_GRPC)
	$(RUN_PROTOC) --go-grpc_out=module=$(MODULE):. $<

server: go/grpshuffle/grpshuffle_grpc.pb.go go/grpshuffle/grpshuffle.pb.go $(wildcard go/server/*.go)
	go build -o $@ ./go/server

client: go/grpshuffle/grpshuffle_grpc.pb.go go/grpshuffle/grpshuffle.pb.go $(wildcard go/client/*.go)
	go build -o $@ ./go/client

doc: $(DOC_DIR)/grpshuffle.md $(DOC_DIR)/index.html
