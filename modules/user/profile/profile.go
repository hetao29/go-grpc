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
	User_Profile.ProfileServer
}

// Register r
func Register(s *grpc.Server) {
	User_Profile.RegisterProfileServer(s, &Profile{})
}

// RegisterHTTP r
func RegisterHTTP(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) {
	User_Profile.RegisterProfileHandlerFromEndpoint(ctx, mux, endpoint, opts)
}
// Update u
func (i *Profile) Update(ctx context.Context, req *User_Profile.UpdateRequest) (*User_Profile.UpdateResponse, error) {
	response := &User_Profile.UpdateResponse{}
	return response, nil // status.Errorf(codes.Unimplemented, "Logined ")
	//return nil, status.Errorf(codes.Unimplemented, "user not exists or pwd error!")
}
