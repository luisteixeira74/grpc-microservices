PROTO_FILES := $(wildcard proto/*.proto)
PB_GO_FILES := $(PROTO_FILES:.proto=.pb.go)

all: $(PB_GO_FILES)

%.pb.go: %.proto
	protoc \
		--proto_path=proto \
		--go_out=proto --go_opt=paths=source_relative \
		--go-grpc_out=proto --go-grpc_opt=paths=source_relative \
		$<

clean:
	rm -f proto/*.pb.go
