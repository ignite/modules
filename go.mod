module github.com/ignite/modules

go 1.19

require (
	cosmossdk.io/errors v1.0.0-beta.7
	cosmossdk.io/math v1.0.0
	github.com/cosmos/cosmos-proto v1.0.0-beta.2
	github.com/cosmos/cosmos-sdk v0.47.1
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.3
	github.com/golangci/golangci-lint v1.50.1
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/spf13/cast v1.5.0
	github.com/spf13/cobra v1.6.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.8.2
	github.com/cometbft/cometbft v0.37.0
	github.com/cometbft/cometbft-db v0.7.0
	golang.org/x/tools v0.6.0
	golang.org/x/vuln v0.0.0-20221122171214-05fb7250142c
	google.golang.org/genproto v0.0.0-20230216225411-c8e22ba71e44
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.29.1
	gopkg.in/yaml.v2 v2.4.0
	mvdan.cc/gofumpt v0.4.0
)

replace (
	github.com/syndtr/goleveldb => github.com/syndtr/goleveldb v1.0.1-0.20210819022825-2ae1ddf74ef7
)
