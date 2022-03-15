regenerate:
	go install github.com/cosmos/gogoproto/protoc-gen-gogo
	protoc --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/cosmos/gogoproto/protobuf/ \
	--gogo_out=. events.proto 
