module github.com/cosmos/gogoproto

go 1.19

require (
	github.com/golang/protobuf v1.5.3
	github.com/google/go-cmp v0.5.9
	golang.org/x/exp v0.0.0-20230811145659-89c5cff77bcb
	google.golang.org/grpc v1.57.0
	google.golang.org/protobuf v1.31.0
)

require (
	golang.org/x/net v0.14.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
	golang.org/x/text v0.12.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230815205213-6bfd019c3878 // indirect
)

// API changed in an incompatible way
retract v1.4.8
