PROTOC_VERSION=23.3
DIRS=$(shell ls ../protobuf)

.PHONY: all
.PHONY: clean
.PHONY: distclean
.PHONY: build

all:	tools build

clean:
	rm -rf $(DIRS)

distclean: clean
	rm -rf tools

tools:
	mkdir tools \
	&& cd tools \
	&& wget -q https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip \
	&& unzip protoc-${PROTOC_VERSION}-linux-x86_64.zip \
	&& rm protoc-${PROTOC_VERSION}-linux-x86_64.zip \
	&& go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31 \
	&& go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

$(DIRS): tools
	find ../protobuf -name *.proto -exec 'tools/bin/protoc' '-I../protobuf' '-Itools/include' '--go_out=.' '--go_opt=paths=source_relative' '--go-grpc_out=.' '--go-grpc_opt=paths=source_relative' '{}' ';'

build: $(DIRS)
	go mod tidy
