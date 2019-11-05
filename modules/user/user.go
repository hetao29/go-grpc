package user
import (
	context "context"
	//"fmt"
	"google.golang.org/grpc"
	//"proto/user/info"
	"modules/user/info"
	"modules/user/profile"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	//status "google.golang.org/grpc/status"
	//codes "google.golang.org/grpc/codes"
)

func init() {
}
// Register ，注册方法
func Register(s *grpc.Server){
	info.Register(s)
	profile.Register(s)
}
// RegisterHTTP 注册 http
func RegisterHTTP(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption){
	info.RegisterHttp(ctx , mux , endpoint , opts ) ;
	profile.RegisterHttp(ctx , mux , endpoint , opts ) ;
}
