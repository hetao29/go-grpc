module main

go 1.17

replace modules/user => ./modules/user

replace common => ./proto/common

replace modules/user/info => ./modules/user/info

replace modules/utility => ./modules/utility

replace proto/user/info => ./proto/user/info

replace proto/order/info => ./proto/order/info

replace proto/user/profile => ./proto/user/profile

require (
	github.com/golang/protobuf v1.5.2
	github.com/gookit/config/v2 v2.0.27
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	golang.org/x/net v0.0.0-20211216030914-fe4d6282115f
	google.golang.org/genproto v0.0.0-20211223182754-3ac035c7e7cb
	google.golang.org/grpc v1.43.0
	modules/user v0.0.0-00010101000000-000000000000
	modules/utility v0.0.0-00010101000000-000000000000
	proto/user/info v0.0.0-00010101000000-000000000000
)

require (
	common v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-redis/redis/v7 v7.4.1 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jinzhu/gorm v1.9.16 // indirect
	proto/user/profile v0.0.0-00010101000000-000000000000 // indirect
)
