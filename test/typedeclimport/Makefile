regenerate:
	cd .. && protoc-min-version --version="3.0.0" --gogo_out=Mtypedeclimport/subpkg/subpkg.proto=github.com/cosmos/gogoproto/test/typedeclimport/subpkg:.  --proto_path=../:../protobuf/:. typedeclimport/typedeclimport.proto
	cd .. && protoc-min-version --version="3.0.0" --gogo_out=. --proto_path=../:../protobuf/:. ./subpkg/subpkg/./subpkg/subpkg.proto