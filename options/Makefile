.PHONY: proto-gen test

all: regenerate test

regenerate:
	protoc --gogoslick_out=Mgoogle/protobuf/descriptor.proto=github.com/cosmos/gogoproto/protoc-gen-gogo/descriptor,paths=source_relative:. cosmos.proto


test:
	protoc -I=. --gocosmos_out=plugins=interfacetype,Mgoogle/protobuf/any.proto=github.com/cosmos/gogoproto/types:. test/abc.proto
	go test ./test
