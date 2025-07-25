package user

import (
	context "context"
	//"fmt"
	"google.golang.org/grpc"
	//"proto/user/info"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hetao29/go-grpc/modules/user/info"
	"github.com/hetao29/go-grpc/modules/user/profile"
	//status "google.golang.org/grpc/status"
	//codes "google.golang.org/grpc/codes"
)

func init() {
}

// Register ，注册方法
func Register(s *grpc.Server) {
	info.Register(s)
	profile.Register(s)
}

// RegisterHTTP 注册 http
func RegisterHTTP(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) {
	info.RegisterHTTP(ctx, mux, endpoint, opts)
	profile.RegisterHTTP(ctx, mux, endpoint, opts)
}
