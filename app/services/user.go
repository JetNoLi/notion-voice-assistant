package services

import (
	userModel "github.com/jetnoli/notion-voice-assistant/models/user"
)

func SignUpUser(userDetails *userModel.User) (user *userModel.User, err error) {

	return userModel.Create(userDetails)
}

func GetAllUsers() (users []*userModel.User, err error) {
	return userModel.GetAll()
}

func GetUserById(id int) (user *userModel.User, err error) {
	return userModel.GetById(id)
}

func UpdateUserById(id int, updates *userModel.UserUpdateBody) (*userModel.User, error) {
	return userModel.UpdateById(id, *updates)
}

func GetUserByUsername(username string) (user *userModel.User, err error) {
	return userModel.GetByUsername(username)
}

func DeleteAllUsers() error {
	return userModel.DeleteAll()
}

func DeleteUserById(id int) error {
	return userModel.DeleteById(id)
}
