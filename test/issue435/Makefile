regenerate:
	go install github.com/cosmos/gogoproto/protoc-gen-gogo
	protoc --gogo_out=. --proto_path=../../../../../:../../protobuf/:. *.proto
