module main

go 1.12

replace models/user => ./models/user

replace proto/user/info => ./proto/proto_src/user/info

replace proto/order/info => ./proto/proto_src/order/info

replace proto/user/profile => ./proto/proto_src/user/profile

require (
	google.golang.org/grpc v1.24.0
	models/user v0.0.0-00010101000000-000000000000
	proto/user/info v0.0.0-00010101000000-000000000000
	proto/user/profile v0.0.0-00010101000000-000000000000
)
