package credentials

import (
	"fmt"

	"github.com/jetnoli/notion-voice-assistant/db"
)

type UserCredential struct {
	Id       int
	UserId   int
	Password string
	Salt     string
}

type Properties struct {
	UserId   int
	Password string
	Salt     string
}

func Seed() {
	fmt.Println("Starting UserCredentials Seeding")

	query := db.Connect()

	_, err := query(`
		drop table user_credentials;
	`)

	if err != nil {
		panic("Failed To Drop Users Table: " + err.Error())
	}

	_, err = query(`create table user_credentials (
		id int generated always as identity primary key,
		user_id int,
		password varchar(255),
		salt varchar(255),
		foreign key (user_id) references users(id)
	);`)

	if err != nil {
		panic("Failed to Create Users Table: " + err.Error())
	}

	fmt.Println("Users Seeded Successfully")
}

func Create(userCredential *UserCredential) (*UserCredential, error) {
	newUserCredential := &UserCredential{}

	query := db.QueryRow(fmt.Sprintf(`
		insert into user_credentials (user_id, password, salt)
		values ( '%d', '%s', '%s')
		returning id, user_id, password, salt;
		`, userCredential.UserId, userCredential.Password, userCredential.Salt))

	err := query.Scan(&newUserCredential.Id, &newUserCredential.UserId, &newUserCredential.Password, &newUserCredential.Salt)

	if err != nil {
		return &UserCredential{}, err
	}

	return newUserCredential, err
}

func GetAll() (usersCredentials []*UserCredential, err error) {
	query, err := db.Query(`
	    select * from user_credentials;
	`)

	if err != nil {
		return usersCredentials, err
	}

	usersCredentials = make([]*UserCredential, 0)

	for i := query.Next(); i; i = query.Next() {
		newUser := &UserCredential{}

		err = query.Scan(&newUser.Id, &newUser.UserId, &newUser.Password, &newUser.Salt)

		if err != nil {
			return usersCredentials, err
		}

		usersCredentials = append(usersCredentials, newUser)
	}

	query.Close()
	err = query.Err()

	return usersCredentials, err
}

func GetById(id int) (userCredentials *UserCredential, err error) {
	userCredentials = &UserCredential{}

	query := db.QueryRow(fmt.Sprintf(`
		select id, user_id, password, salt from user_credentials
		where id=%d
	`, id))

	err = query.Scan(&userCredentials.Id, &userCredentials.UserId, &userCredentials.Password, &userCredentials.Salt)

	return userCredentials, err
}

// TODO: Update Credential By Id
// func UpdateById(id int) (*UserCredential, error) {

// 	bodyMap, err := utils.StructToMap(body)

// 	if err != nil {
// 		return &UserCredential{}, err
// 	}

// 	bodyString := ""

// 	for key, value := range bodyMap {
// 		if bodyString != "" {
// 			bodyString += ", "
// 		}

// 		bodyString += fmt.Sprintf(`%s='%s'`, key, value)
// 	}

// 	query := db.QueryRow(fmt.Sprintf(`
// 		update user_credentials
// 		set %s
// 		where id=%d
// 		returning id, username, email, password
// 	`, bodyString, id))

// 	newUser := &UserCredential{}

// 	err = query.Scan(&newUser.Id, &newUser.Username, &newUser.Email, &newUser.Password)

// 	return newUser, err
// }

func DeleteById(id int) error {
	_, err := db.Query(fmt.Sprintf(`
	delete from user_credentials
	where id=%d
	`, id))

	return err
}

func DeleteAll() error {
	_, err := db.Query(`delete from user_credentials`)

	return err
}

func GetByUserId(userId int) (userCredential *UserCredential, err error) {
	userCredential = &UserCredential{}

	query := db.QueryRow(fmt.Sprintf(`
		select * from user_credentials
		where user_id=%d
	`, userId))

	err = query.Scan(&userCredential.Id, &userCredential.UserId, &userCredential.Password, &userCredential.Salt)

	if err != nil {
		return nil, err
	}

	return userCredential, nil

}
