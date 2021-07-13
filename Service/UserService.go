package Service

// IUserService
type IUserService interface {
	GetName(userid int) string
}

// UserService
type UserService struct {

}

// GetName
func (obj UserService) GetName(userid int) string {
	if userid == 200 {
		return "200-Name"
	}
	return "404-Name"
}
