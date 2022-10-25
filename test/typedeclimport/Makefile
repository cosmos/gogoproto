regenerate:
	protoc-min-version --version="3.0.0" --gogo_out=Mtest/typedeclimport/subpkg/subpkg.proto=github.com/cosmos/gogoproto/test/typedeclimport/subpkg:.  --proto_path=../../:../../protobuf/:. typedeclimport.proto
	protoc-min-version --version="3.0.0" --gogo_out=. --proto_path=../../:../../protobuf/:. ./subpkg/subpkg.proto