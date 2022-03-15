regenerate:
	go install github.com/cosmos/gogoproto/protoc-min-version
	go install github.com/cosmos/gogoproto/protoc-gen-gogoslick
	protoc-min-version --version="3.0.0" --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/cosmos/gogoproto/protobuf/ \
	--gogoslick_out=Mgoogle/protobuf/timestamp.proto=github.com/cosmos/gogoproto/types:. timefail.proto 
