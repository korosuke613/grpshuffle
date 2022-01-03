# Tool versoins
PROTOC_VERSION = 3.17.3
PROTOC_GEN_DOC_VERSION = 1.4.1
PROTOC_GEN_GO_VERSION = 1.26.0
PROTOC_GEN_GO_GRPC_VERSION = 1.1.0

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

PLATFORM = osx
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	PLATFORM = linux
endif
ifeq ($(UNAME_S),Darwin)
	PLATFORM = osx
endif

.PHONY: all
all: doc grpshuffle_server grpshuffle_client

.PHONY: clean
clean:
	rm -f $(DOC_DIR)/index.html $(DOC_DIR)/grpshuffle.md grpshuffle_server grpshuffle_client

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
$(DOC_MD): $(wildcard *.proto) $(DOC_DIR) $(PROTOC) $(PROTOC_GEN_DOC)
	$(RUN_PROTOC) --doc_out=$(DOC_DIR) --doc_opt=markdown,$@ grpshuffle.proto

# generate html specification
# $(DOC_HTML): grpshuffle.proto healthcheck.proto $(DOC_DIR) $(PROTOC) $(PROTOC_GEN_DOC)
$(DOC_HTML): $(wildcard *.proto) $(DOC_DIR) $(PROTOC) $(PROTOC_GEN_DOC)
	$(RUN_PROTOC) --doc_out=$(DOC_DIR) grpshuffle.proto

go/grpshuffle/%.pb.go: %.proto $(PROTOC) $(PROTOC_GEN_GO)
	$(RUN_PROTOC) --go_out=module=$(MODULE):. $<

go/grpshuffle/%_grpc.pb.go: %.proto $(PROTOC) $(PROTOC_GEN_GO_GRPC)
	$(RUN_PROTOC) --go-grpc_out=module=$(MODULE):. $<

grpshuffle_server: go/grpshuffle/grpshuffle_grpc.pb.go go/grpshuffle/grpshuffle.pb.go $(wildcard go/grpshuffle_server/*.go)
	go build -o $@ ./go/grpshuffle_server
	chmod +x $@

grpshuffle_client: go/grpshuffle/grpshuffle_grpc.pb.go go/grpshuffle/grpshuffle.pb.go $(wildcard go/grpshuffle_client/*.go)
	go build -o $@ ./go/grpshuffle_client
	chmod +x $@

doc: $(DOC_DIR)/grpshuffle.md $(DOC_DIR)/index.html
