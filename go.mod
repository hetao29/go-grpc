module main

go 1.12

replace models/user => ./models/user

replace proto/user/info => ./proto/user/info

replace proto/order/info => ./proto/order/info

replace proto/user/profile => ./proto/user/profile

require (
	google.golang.org/grpc v1.24.0
	models/user v0.0.0-00010101000000-000000000000
	proto/user/info v0.0.0-00010101000000-000000000000
	proto/user/profile v0.0.0-00010101000000-000000000000
)
