module github.com/cosmos/gogoproto

go 1.26.0

require (
	github.com/golang/protobuf v1.5.4
	github.com/google/go-cmp v0.7.0
	github.com/tendermint/go-amino v0.16.0
	google.golang.org/grpc v1.85.0
	google.golang.org/protobuf v1.36.12
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	golang.org/x/net v0.60.0 // indirect
	golang.org/x/sys v0.48.0 // indirect
	golang.org/x/text v0.42.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260817212433-ac3dfec99bb1 // indirect
)

// API changed in an incompatible way
retract v1.4.8
