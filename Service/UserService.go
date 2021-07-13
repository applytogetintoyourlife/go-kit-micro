package Service

import "errors"

// IUserService
type IUserService interface {
	GetName(userid int) string
	DelUser(userid int) error
}

// UserService
type UserService struct {

}

// GetName
func (obj UserService) GetName(userid int) string {
	if userid == 200 {
		return "200-Get-Name"
	}
	return "404-Get-Name"
}

// DelUser
func (obj UserService) DelUser(userid int) error {
	if userid != 200 {
		return errors.New("No-Permission")
	}
	return nil
}