package user
import (
	context "context"
	"google.golang.org/grpc"
	"proto/user"
	status "google.golang.org/grpc/status"
	codes "google.golang.org/grpc/codes"
)

func init() {
}
type Info struct {
	user.UnimplementedInfoServer
}
func Register(s *grpc.Server, srv user.InfoServer){
	user.RegisterInfoServer(s, srv)
}
func (info*Info) Login(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (info*Info) Logout(ctx context.Context, req *user.LogoutRequest) (*user.LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
