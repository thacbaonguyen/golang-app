package user

import (
	"errors"
	"go-ginapp/data/request"
	"go-ginapp/data/response"
	"go-ginapp/repositories"
	"go-ginapp/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	userRepo repositories.UserRepository
	roleRepo repositories.RoleRepository
}

func NewUserRepositoryImpl(userRepo repositories.UserRepository,
	roleRepo repositories.RoleRepository) UserService {
	return &UserServiceImpl{userRepo: userRepo, roleRepo: roleRepo}
}

func (u *UserServiceImpl) GetAllUser() ([]response.UserResponse, error) {
	users, err := u.userRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var usersResponses []response.UserResponse
	for _, usr := range users {
		usrResponse, err := utils.ToUserResponse(usr)
		if err != nil {
			return nil, err
		}
		usersResponses = append(usersResponses, usrResponse)
	}
	return usersResponses, nil
}

func (u *UserServiceImpl) GetUserByID(id uint) (response.UserResponse, error) {
	user, err := u.userRepo.FindById(id)
	if err != nil {
		return response.UserResponse{}, err
	}
	userResponse, err := utils.ToUserResponse(user)
	if err != nil {
		return response.UserResponse{}, nil
	}
	return userResponse, nil
}

func (u *UserServiceImpl) UpdateUser(id uint, userUpdateReq request.UserUpdateReq) (response.UserResponse, error) {
	existingUser, err := u.userRepo.FindById(id)
	if err != nil {
		return response.UserResponse{}, err
	}
	existingUser.FullName = userUpdateReq.FullName

	userUpdated, err := u.userRepo.Update(existingUser)
	if err != nil {
		return response.UserResponse{}, err
	}
	userResponse, err := utils.ToUserResponse(userUpdated)
	if err != nil {
		return response.UserResponse{}, err
	}
	return userResponse, nil
}

func (u *UserServiceImpl) UpdatePassword(id uint, passwordRequest request.UpdatePasswordRequest) error {
	user, err := u.userRepo.FindById(id)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordRequest.OldPassword))
	if err != nil {
		return err
	}
	if passwordRequest.Password != passwordRequest.RetypePassword {
		return errors.New("password does not match")
	}

	user.Password = passwordRequest.Password
	return u.userRepo.UpdatePassword(user)
}

func (u *UserServiceImpl) DeleteUser(id uint) error {
	return u.userRepo.Delete(id)
}
