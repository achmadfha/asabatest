package authentication

import "BE-shop/models/dto/authenticationDto"

type AuthenticationRepository interface {
	RegisterUsers(req authenticationDto.Register) error
	CheckEmailExists(usrEmail string) (bool, error)
	RetrieveUsers(usrEmail string) (usr authenticationDto.Register, err error)
}

type AuthenticationUseCase interface {
	RegisterUsers(req authenticationDto.RegisterReq) (authenticationDto.RegisterRes, error)
	LoginUsers(req authenticationDto.LoginReq) (token string, err error)
}
