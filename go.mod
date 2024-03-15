module github.com/cosmos/gogoproto

go 1.21

toolchain go1.21.7

require (
	cosmossdk.io/x/tx v0.13.0
	github.com/google/go-cmp v0.6.0
	github.com/stretchr/testify v1.8.4
	github.com/tendermint/go-amino v0.16.0
	golang.org/x/exp v0.0.0-20231006140011-7918f672742d
	google.golang.org/grpc v1.61.1
	google.golang.org/protobuf v1.32.0
)

require (
	cosmossdk.io/api v0.7.2 // indirect
	cosmossdk.io/core v0.11.0 // indirect
	github.com/cosmos/cosmos-proto v1.0.0-beta.3 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/kisielk/errcheck v1.7.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/tools v0.17.0 // indirect
	google.golang.org/genproto v0.0.0-20231211222908-989df2bf70f3 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20231120223509-83a465c0220f // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231212172506-995d672761c0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// API changed in an incompatible way
retract v1.4.8
