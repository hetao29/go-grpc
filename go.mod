module main

go 1.17

replace modules/user => ./modules/user

replace common => ./proto/common

replace modules/user/info => ./modules/user/info

replace modules/log => ./modules/log

replace modules/utility => ./modules/utility

replace proto/user/info => ./proto/user/info

replace proto/order/info => ./proto/order/info

replace proto/user/profile => ./proto/user/profile

require (
	github.com/facebookgo/grace v0.0.0-20180706040059-75cf19382434
	github.com/golang/protobuf v1.5.2
	github.com/gookit/config/v2 v2.0.27
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	golang.org/x/net v0.0.0-20211216030914-fe4d6282115f
	google.golang.org/genproto v0.0.0-20211223182754-3ac035c7e7cb
	google.golang.org/grpc v1.43.0
	modules/log v0.0.0-00010101000000-000000000000
	modules/user v0.0.0-00010101000000-000000000000
	modules/utility v0.0.0-00010101000000-000000000000
	proto/user/info v0.0.0-00010101000000-000000000000
)

require (
	common v0.0.0-00010101000000-000000000000 // indirect
	github.com/facebookgo/ensure v0.0.0-20200202191622-63f1cf65ac4c // indirect
	github.com/facebookgo/freeport v0.0.0-20150612182905-d4adf43b75b9 // indirect
	github.com/facebookgo/stack v0.0.0-20160209184415-751773369052 // indirect
	github.com/facebookgo/subset v0.0.0-20200203212716-c811ad88dec4 // indirect
	github.com/go-redis/redis/v7 v7.4.1 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/google/logger v1.1.1 // indirect
	github.com/gookit/goutil v0.3.15 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	proto/user/profile v0.0.0-00010101000000-000000000000 // indirect
)
