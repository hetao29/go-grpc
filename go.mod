module main

go 1.12

replace models/user => ./models/user

replace proto/user => ./proto/proto_src/user

require (
	google.golang.org/grpc v1.24.0
	models/user v0.0.0-00010101000000-000000000000
	proto/user v0.0.0-00010101000000-000000000000
)
