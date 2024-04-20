package authenticationUseCase

import (
	"BE-shop/models/constants"
	"BE-shop/models/dto/authenticationDto"
	"BE-shop/pkg/utils"
	"BE-shop/src/authentication"
	"errors"
	"github.com/google/uuid"
	"time"
)

type authenticationUC struct {
	authenticationRepository authentication.AuthenticationRepository
}

func NewAuthenticationUseCase(authentication authentication.AuthenticationRepository) authentication.AuthenticationUseCase {
	return &authenticationUC{authentication}
}

func (a authenticationUC) RegisterUsers(req authenticationDto.RegisterReq) (authenticationDto.RegisterRes, error) {
	emailExists, err := a.authenticationRepository.CheckEmailExists(req.Email)
	if err != nil {
		return authenticationDto.RegisterRes{}, err
	}

	if emailExists {
		return authenticationDto.RegisterRes{}, errors.New("01")
	}

	usrID, err := uuid.NewRandom()
	if err != nil {
		return authenticationDto.RegisterRes{}, err
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	usrRole := constants.DefaultRole
	currentTime := time.Now()

	usrData := authenticationDto.Register{
		UsersID:   usrID,
		Email:     req.Email,
		Password:  hashedPassword,
		Roles:     usrRole,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	err = a.authenticationRepository.RegisterUsers(usrData)
	if err != nil {
		return authenticationDto.RegisterRes{}, err
	}

	usrRes := authenticationDto.RegisterRes{
		UsersID:   usrID,
		Email:     req.Email,
		CreatedAt: currentTime,
	}

	return usrRes, nil
}

func (a authenticationUC) LoginUsers(req authenticationDto.LoginReq) (token string, err error) {
	usr, err := a.authenticationRepository.RetrieveUsers(req.Email)
	if err != nil {
		// 01 email not registered
		if err.Error() == "01" {
			return "", errors.New("01")
		}
		return "", err
	}

	if err := utils.VerifyPassword(usr.Password, req.Password); err != nil {
		// 02 password didn't match
		return "", errors.New("02")
	}

	token, err = utils.GenerateToken(usr.UsersID, usr.Roles)
	if err != nil {
		return "", err
	}

	return token, nil
}
