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
	common v0.0.0-00010101000000-000000000000 // indirect
	github.com/facebookgo/clock v0.0.0-20150410010913-600d898af40a // indirect
	github.com/facebookgo/grace v0.0.0-20180706040059-75cf19382434
	github.com/facebookgo/httpdown v0.0.0-20180706035922-5979d39b15c2 // indirect
	github.com/facebookgo/stats v0.0.0-20151006221625-1b76add642e4 // indirect
	github.com/go-redis/redis/v7 v7.4.1 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/google/logger v1.1.1
	github.com/gookit/config/v2 v2.0.27
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/jinzhu/gorm v1.9.16 // indirect
	google.golang.org/grpc v1.43.0
	modules/log v0.0.0-00010101000000-000000000000
	modules/user v0.0.0-00010101000000-000000000000
	modules/utility v0.0.0-00010101000000-000000000000
	proto/user/info v0.0.0-00010101000000-000000000000
	proto/user/profile v0.0.0-00010101000000-000000000000
)
