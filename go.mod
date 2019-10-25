module main

go 1.12

replace models/user => ./models/user

replace proto/user/info => ./proto/user/info

replace proto/order/info => ./proto/order/info

replace proto/user/profile => ./proto/user/profile

require (
	github.com/Masterminds/goutils v1.1.0 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Masterminds/sprig v2.22.0+incompatible // indirect
	github.com/facebookgo/clock v0.0.0-20150410010913-600d898af40a // indirect
	github.com/facebookgo/grace v0.0.0-20180706040059-75cf19382434
	github.com/facebookgo/httpdown v0.0.0-20180706035922-5979d39b15c2 // indirect
	github.com/facebookgo/stats v0.0.0-20151006221625-1b76add642e4 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/gookit/config/v2 v2.0.12
	github.com/grpc-ecosystem/grpc-gateway v1.11.3
	github.com/huandu/xstrings v1.2.0 // indirect
	github.com/imdario/mergo v0.3.8 // indirect
	github.com/lyft/protoc-gen-validate v0.1.0 // indirect
	github.com/mitchellh/copystructure v1.0.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.1 // indirect
	github.com/mwitkow/go-proto-validators v0.2.0 // indirect
	github.com/pseudomuto/protoc-gen-doc v1.3.0 // indirect
	github.com/pseudomuto/protokit v0.2.0 // indirect
	golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550 // indirect
	google.golang.org/genproto v0.0.0-20191009194640-548a555dbc03 // indirect
	google.golang.org/grpc v1.24.0
	models/user v0.0.0-00010101000000-000000000000
	proto/user/info v0.0.0-00010101000000-000000000000
	proto/user/profile v0.0.0-00010101000000-000000000000
)
