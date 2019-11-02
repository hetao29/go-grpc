package profile
import (
	context "context"
	//"fmt"
	"google.golang.org/grpc"
	"proto/user/profile"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	//status "google.golang.org/grpc/status"
	//codes "google.golang.org/grpc/codes"
)

func init() {
}
type Profile struct {
	profile.UnimplementedProfileServer
}
func Register(s *grpc.Server){
	profile.RegisterProfileServer(s, &Profile{})
}
func RegisterHttp(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) () {
	profile.RegisterProfileHandlerFromEndpoint(ctx , mux , endpoint , opts ) ;
}
