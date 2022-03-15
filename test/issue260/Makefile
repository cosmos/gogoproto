regenerate:
	go install github.com/cosmos/gogoproto/protoc-gen-gogo
	protoc-min-version --version="3.0.0" --gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/cosmos/gogoproto/types:. --proto_path=../../../../../:../../protobuf/:. issue260.proto
