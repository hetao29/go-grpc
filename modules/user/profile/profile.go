package profile

import (
	context "context"
	//"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"proto/user/profile"
	//status "google.golang.org/grpc/status"
	//codes "google.golang.org/grpc/codes"
)

func init() {
}

// Profile p
type Profile struct {
	profile.ProfileServer
}

// Register r
func Register(s *grpc.Server) {
	profile.RegisterProfileServer(s, &Profile{})
}

// RegisterHTTP r
func RegisterHTTP(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) {
	profile.RegisterProfileHandlerFromEndpoint(ctx, mux, endpoint, opts)
}
func (i *Profile) Update(ctx context.Context, req *profile.UpdateRequest) (*profile.UpdateResponse, error) {
	response := &profile.UpdateResponse{}
	return response, nil // status.Errorf(codes.Unimplemented, "Logined ")
	//return nil, status.Errorf(codes.Unimplemented, "user not exists or pwd error!")
}
