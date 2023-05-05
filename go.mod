module github.com/cosmos/gogoproto

go 1.19

require (
	github.com/google/go-cmp v0.5.9
	golang.org/x/exp v0.0.0-20230131160201-f062dba9d201
	google.golang.org/grpc v1.55.0
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230306155012-7f2fa6fef1f4 // indirect
)

// API changed in an incompatible way
retract v1.4.8
