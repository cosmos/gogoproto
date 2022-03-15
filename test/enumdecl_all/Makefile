regenerate:
	go install github.com/cosmos/gogoproto/protoc-gen-gogo
	protoc-min-version --version="3.0.0" --gogo_out=. --proto_path=../../../../../:../../protobuf/:. enumdeclall.proto
