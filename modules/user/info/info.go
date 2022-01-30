package info

import (
	context "context"
	//"log"
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
	User_Info.InfoServer
}

// Register r
func Register(s *grpc.Server) {
	User_Info.RegisterInfoServer(s, &Info{})
	//profile.RegisterProfileServer(s, &Profile{})
}

// RegisterHTTP r
func RegisterHTTP(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	User_Info.RegisterInfoHandlerFromEndpoint(ctx, mux, endpoint, opts)
	//profile.RegisterProfileHandlerFromEndpoint(ctx , mux , endpoint , opts ) ;
	return nil
}

// Login login
func (i *Info) Login(ctx context.Context, req *User_Info.LoginRequest) (*User_Info.LoginResponse, error) {
	user := GetByNameAndPwd(req.Name, req.Password)
	if user.ID != 0 {
		response := &User_Info.LoginResponse{}
		info := &User_Info.UserInfo{}
		info.Uid= (int64)(user.ID)
		info.Name= user.Name
		response.Info=info 
		return response, nil // status.Errorf(codes.Unimplemented, "Logined ")
	}
	return nil, status.Errorf(codes.Unimplemented, "user not exists or pwd error!")
}

// Logout logout
func (i *Info) Logout(ctx context.Context, req *User_Info.LogoutRequest) (*User_Info.LogoutResponse, error) {
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
