module main

go 1.13

replace modules/user => ./modules/user

replace common => ./proto/common

replace modules/user/info => ./modules/user/info

replace modules/log => ./modules/log

replace modules/utility => ./modules/utility

replace proto/user/info => ./proto/user/info

replace proto/order/info => ./proto/order/info

replace proto/user/profile => ./proto/user/profile

require (
	common v0.0.0-00010101000000-000000000000 // indirect
	github.com/Masterminds/goutils v1.1.0 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Masterminds/sprig v2.22.0+incompatible // indirect
	github.com/facebookgo/clock v0.0.0-20150410010913-600d898af40a // indirect
	github.com/facebookgo/grace v0.0.0-20180706040059-75cf19382434
	github.com/facebookgo/httpdown v0.0.0-20180706035922-5979d39b15c2 // indirect
	github.com/facebookgo/stats v0.0.0-20151006221625-1b76add642e4 // indirect
	github.com/go-redis/redis/v7 v7.0.0-beta.4 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/google/logger v1.0.1
	github.com/google/uuid v1.1.1 // indirect
	github.com/gookit/config/v2 v2.0.12
	github.com/grpc-ecosystem/grpc-gateway v1.12.1
	github.com/huandu/xstrings v1.2.0 // indirect
	github.com/imdario/mergo v0.3.8 // indirect
	github.com/jinzhu/gorm v1.9.11 // indirect
	github.com/lyft/protoc-gen-validate v0.1.0 // indirect
	github.com/mitchellh/copystructure v1.0.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.1 // indirect
	github.com/mwitkow/go-proto-validators v0.2.0 // indirect
	github.com/pseudomuto/protoc-gen-doc v1.3.0 // indirect
	github.com/pseudomuto/protokit v0.2.0 // indirect
	golang.org/x/crypto v0.0.0-20191112222119-e1110fd1c708 // indirect
	google.golang.org/genproto v0.0.0-20191108220845-16a3f7862a1a // indirect
	google.golang.org/grpc v1.25.0
	gopkg.in/yaml.v2 v2.2.5 // indirect
	modules/log v0.0.0-00010101000000-000000000000
	modules/user v0.0.0-00010101000000-000000000000
	modules/utility v0.0.0-00010101000000-000000000000
	proto/user/info v0.0.0-00010101000000-000000000000
	proto/user/profile v0.0.0-00010101000000-000000000000
)
