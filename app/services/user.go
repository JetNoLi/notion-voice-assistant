package services

import (
	"fmt"

	userModel "github.com/jetnoli/notion-voice-assistant/models/user"
	"github.com/jetnoli/notion-voice-assistant/models/user/credentials"
	"github.com/jetnoli/notion-voice-assistant/utils"
)

type SignUpRequestBody struct {
	Username string
	Email    string
	Password string
}

type SignInRequestBody struct {
	Email    string
	Password string
}

func SignUp(userDetails *SignUpRequestBody) (user *userModel.User, err error) {

	// Generate Salt and Password
	password, salt, err := utils.GenerateEncodedSaltAndPasswordHash(userDetails.Username)

	if err != nil {
		return user, err
	}

	//TODO: First Check if user with given username or email already exists

	// Create User In DB
	user, err = userModel.Create(&userModel.Properties{Username: &userDetails.Username, Email: &userDetails.Email})

	if err != nil {
		return user, err
	}

	_, err = credentials.Create(&credentials.UserCredential{Password: password, Salt: salt, UserId: user.Id})

	return user, err
}

func SignIn(userDetails *SignInRequestBody) (user *userModel.User, err error) {
	user, err = GetUserByEmail(userDetails.Email)

	if err != nil {
		return nil, err
	}

	userCredentials, err := credentials.GetByUserId(user.Id)

	if err != nil {
		return nil, err
	}

	fmt.Println(userDetails)
	fmt.Println(userCredentials)

	isAuthenticated, err := utils.DecodeAndComparePasswords(userDetails.Password, userCredentials.Password, userCredentials.Salt)

	if err != nil {
		return nil, err
	}

	if !isAuthenticated {
		return nil, fmt.Errorf("password does not match")
	}

	return user, nil
}

func CreateUser(userDetails *userModel.Properties) (*userModel.User, error) {
	return userModel.Create(userDetails)
}

func GetAllUsers() (users []*userModel.User, err error) {
	return userModel.GetAll()
}

func GetUserById(id int) (user *userModel.User, err error) {
	return userModel.GetById(id)
}

func UpdateUserById(id int, updates *userModel.Properties) (*userModel.User, error) {
	return userModel.UpdateById(id, *updates)
}

func GetUserByUsername(username string) (*userModel.User, error) {
	return userModel.GetByUsername(username)
}

func GetUserByEmail(email string) (*userModel.User, error) {
	return userModel.GetByEmail(email)
}

func DeleteAllUsers() error {
	return userModel.DeleteAll()
}

func DeleteUserById(id int) error {
	return userModel.DeleteById(id)
}
