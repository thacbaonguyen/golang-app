package auth

import (
	"errors"
	"go-ginapp/data/request"
	"go-ginapp/data/response"
	"go-ginapp/models"
	"go-ginapp/repositories"
	"go-ginapp/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	userRepo   repositories.UserRepository
	roleRepo   repositories.RoleRepository
	jwtService JWTService
}

func NewAuthServiceImpl(userRepo repositories.UserRepository,
	roleRepo repositories.RoleRepository,
	jwtService JWTService) AuthService {
	return &AuthServiceImpl{userRepo: userRepo,
		roleRepo:   roleRepo,
		jwtService: jwtService,
	}
}

func (a *AuthServiceImpl) Login(req request.LoginRequest) (response.AuthResponse, error) {
	//TODO implement me
	user, err := a.userRepo.FindByUsername(req.Username)
	if err != nil {
		return response.AuthResponse{}, errors.New("user not found")
	}
	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return response.AuthResponse{}, errors.New("invalid credentials")
	}
	// gen token
	token, err := a.jwtService.GenerateToken(user)

	if err != nil {
		return response.AuthResponse{}, err
	}

	userResponse, err := utils.ToUserResponse(user)
	if err != nil {
		return response.AuthResponse{}, errors.New("cannot convert user to user response")
	}

	return response.AuthResponse{
		token,
		userResponse,
	}, nil
}

func (a *AuthServiceImpl) Register(req request.RegisterRequest) (response.AuthResponse, error) {
	_, err := a.userRepo.FindByUsername(req.Username)
	if err != nil {
		return response.AuthResponse{}, errors.New("username already exist")
	}

	_, err = a.userRepo.FindByEmail(req.Email)
	if err != nil {
		return response.AuthResponse{}, errors.New("email already exist")
	}

	defaultRole, err := a.roleRepo.FindByName("user")
	if err != nil {
		return response.AuthResponse{}, errors.New("cannot found user role")
	}

	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		FullName: req.FullName,
		RoleId:   defaultRole.ID,
		Role:     defaultRole,
	}

	user, err = a.userRepo.Create(user)
	if err != nil {
		return response.AuthResponse{}, err
	}

	token, err := a.jwtService.GenerateToken(user)
	if err != nil {
		return response.AuthResponse{}, err
	}

	userResponse, err := utils.ToUserResponse(user)
	if err != nil {
		return response.AuthResponse{}, err
	}

	return response.AuthResponse{
		token,
		userResponse,
	}, nil
}
