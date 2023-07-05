module main

go 1.17

replace modules/user => ./modules/user

replace common => ./proto/out/go/common

replace modules/user/info => ./modules/user/info

replace modules/utility => ./modules/utility

replace proto/user/info => ./proto/out/go/user/info

replace proto/order/info => ./proto/out/go/order/info

replace proto/user/profile => ./proto/out/go/user/profile

require (
	github.com/envoyproxy/go-control-plane v0.10.3
	github.com/golang/protobuf v1.5.2
	github.com/gookit/config/v2 v2.0.27
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/lyft/protoc-gen-star v0.6.1
	golang.org/x/net v0.7.0
	google.golang.org/api v0.103.0
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f
	google.golang.org/grpc v1.53.0
	modules/user v0.0.0-00010101000000-000000000000
	modules/utility v0.0.0-00010101000000-000000000000
	proto/user/info v0.0.0-00010101000000-000000000000
)

require (
	github.com/gookit/goutil v0.6.0 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

require (
	cloud.google.com/go/compute v1.15.1 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	common v0.0.0-00010101000000-000000000000 // indirect
	github.com/census-instrumentation/opencensus-proto v0.4.1 // indirect
	github.com/cncf/xds/go v0.0.0-20230105202645-06c439db220b // indirect
	github.com/envoyproxy/protoc-gen-validate v0.9.1 // indirect
	github.com/go-redis/redis/v7 v7.4.1 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.2.0 // indirect
	github.com/googleapis/gax-go/v2 v2.7.0 // indirect
	github.com/gookit/color v1.5.2 // indirect
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/spf13/afero v1.9.2 // indirect
	github.com/xo/terminfo v0.0.0-20210125001918-ca9a967f8778 // indirect
	go.opencensus.io v0.24.0 // indirect
	golang.org/x/oauth2 v0.4.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/term v0.5.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	proto/user/profile v0.0.0-00010101000000-000000000000 // indirect
)
