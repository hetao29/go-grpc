package user

func init() {
}
func NewUserInfo() *UserInfo {
	return new(UserInfo)
}

type UserInfo struct {
	name string
}

func (this *UserInfo) Set(v string) {
	this.name = v
}
func (this *UserInfo) Get() string {
	return this.name
}
