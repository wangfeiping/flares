module github.com/wangfeiping/flares

go 1.14

require (
	github.com/cosmos/cosmos-sdk v0.40.0-rc0
	github.com/gogo/protobuf v1.3.1
	github.com/google/uuid v1.0.0
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.15.0
	github.com/nxadm/tail v1.4.5 // indirect
	github.com/onsi/ginkgo v1.14.2
	github.com/onsi/gomega v1.10.1
	github.com/regen-network/cosmos-proto v0.3.0
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.0.0
	github.com/tendermint/tendermint v0.34.0-rc4.0.20201005135527-d7d0ffea13c6
	github.com/tendermint/tm-db v0.6.2
	golang.org/x/sys v0.0.0-20201026173827-119d4633e4d1 // indirect
	google.golang.org/grpc v1.32.0
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
