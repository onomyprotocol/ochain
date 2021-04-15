module github.com/onomyprotocol/ochain

go 1.15

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/MakeNowJust/heredoc/v2 v2.0.2-0.20200911102834-82b7f958b059
	github.com/cosmos/go-bip39 v1.0.0
	github.com/gogo/protobuf v1.3.3
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/onomyprotocol/onomy-sdk v0.42.32
	github.com/sethvargo/go-password v0.1.4-0.20200617175411-ca45e1f26826
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1 // indirect
	github.com/tendermint/tendermint v0.34.8
	github.com/tendermint/tm-db v0.6.4
	github.com/yelinaung/go-haikunator v0.0.0-20150320004105-1249cae259af
	google.golang.org/genproto v0.0.0-20210207032614-bba0dbe2a9ea
	google.golang.org/grpc v1.36.1
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776
)

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
