package user
import (
	context "context"
	"google.golang.org/grpc"
	"proto/user/info"
	"proto/user/profile"
	status "google.golang.org/grpc/status"
	codes "google.golang.org/grpc/codes"
)

func init() {
}
type Info struct {
	info.UnimplementedInfoServer
}
func Register(s *grpc.Server){
	info.RegisterInfoServer(s, &Info{})
	profile.RegisterProfileServer(s, &Profile{})
}
func (info*Info) Login(ctx context.Context, req *info.LoginRequest) (*info.LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "Logined ")
}
func (info*Info) Logout(ctx context.Context, req *info.LogoutRequest) (*info.LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "Logoutout ")
}
type Profile struct {
	profile.UnimplementedProfileServer
}
func (profile *Profile) Update(ctx context.Context, req *profile.UpdateRequest) (*profile.UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "Update")
}
