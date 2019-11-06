package info

import (
	context "context"
	"fmt"
	"google.golang.org/grpc"
	"proto/user/info"
	//"proto/user/profile"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

func init() {
}

// Info info
type Info struct {
	info.UnimplementedInfoServer
}

// Register r
func Register(s *grpc.Server) {
	info.RegisterInfoServer(s, &Info{})
	//profile.RegisterProfileServer(s, &Profile{})
}

// RegisterHTTP r
func RegisterHTTP(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	info.RegisterInfoHandlerFromEndpoint(ctx, mux, endpoint, opts)
	//profile.RegisterProfileHandlerFromEndpoint(ctx , mux , endpoint , opts ) ;
	return nil
}

// Login login
func (i *Info) Login(ctx context.Context, req *info.LoginRequest) (*info.LoginResponse, error) {
	user := GetByNameAndPwd(req.Name, req.Password)
	fmt.Println(user.ID)
	if user.ID != 0 {
		response := &info.LoginResponse{}
		info := &info.UserInfo{}
		info.Uid= (int64)(user.ID)
		info.Name= user.Name
		response.Info=info 
		return response, nil // status.Errorf(codes.Unimplemented, "Logined ")
	}
	return nil, status.Errorf(codes.Unimplemented, "user not exists or pwd error!")
}

// Logout logout
func (i *Info) Logout(ctx context.Context, req *info.LogoutRequest) (*info.LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "Logoutout ")
}

/*
type Profile struct {
	profile.UnimplementedProfileServer
}
func (profile *Profile) Update(ctx context.Context, req *profile.UpdateRequest) (*profile.UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "Update")
}
*/
