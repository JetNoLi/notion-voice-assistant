package user

import (
	"fmt"

	"github.com/jetnoli/notion-voice-assistant/db"
	"github.com/jetnoli/notion-voice-assistant/utils"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"` // Hashed Value
}

type UserUpdateBody struct {
	Username *string `json:"username"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

func Seed() {
	fmt.Println("Starting User Seeding")

	query := db.Connect()

	_, err := query(`
		drop table users;
	`)

	if err != nil {
		panic("Failed To Drop Users Table: " + err.Error())
	}

	_, err = query(`create table users (
		id int generated always as identity,
		username varchar(255),
		email varchar(255),
		password varchar(255)
	);`)

	if err != nil {
		panic("Failed to Create Users Table: " + err.Error())
	}

	fmt.Println("Users Seeded Successfully")
}

func Create(user *User) (*User, error) {
	newUser := &User{}

	query := db.QueryRow(fmt.Sprintf(`
		insert into users (username, email, password)
		values ( '%s', '%s', '%s')
		returning id, username, email, password;
		`, user.Username, user.Email, user.Password))

	err := query.Scan(&newUser.Id, &newUser.Username, &newUser.Email, &newUser.Password)

	if err != nil {
		return &User{}, err
	}

	return newUser, err
}

func GetAll() (users []*User, err error) {
	query, err := db.Query(`
	    select * from users;
	`)

	if err != nil {
		return users, err
	}

	users = make([]*User, 0)

	for i := query.Next(); i; i = query.Next() {
		newUser := &User{}

		err = query.Scan(&newUser.Id, &newUser.Username, &newUser.Email, &newUser.Password)

		if err != nil {
			return users, err
		}

		users = append(users, newUser)
	}

	query.Close()
	err = query.Err()

	return users, err
}

func GetById(id int) (user *User, err error) {
	user = &User{}

	query := db.QueryRow(fmt.Sprintf(`
		select * from users
		where id=%d
	`, id))

	err = query.Scan(&user.Id, &user.Username, &user.Email, &user.Password)

	return user, err
}

func UpdateById(id int, body UserUpdateBody) (*User, error) {

	bodyMap, err := utils.StructToMap(body)

	if err != nil {
		return &User{}, err
	}

	bodyString := ""

	for key, value := range bodyMap {
		if bodyString != "" {
			bodyString += ", "
		}

		bodyString += fmt.Sprintf(`%s='%s'`, key, value)
	}

	query := db.QueryRow(fmt.Sprintf(`
		update users
		set %s
		where id=%d
		returning id, username, email, password
	`, bodyString, id))

	newUser := &User{}

	err = query.Scan(&newUser.Id, &newUser.Username, &newUser.Email, &newUser.Password)

	return newUser, err
}

func DeleteById(id int) error {
	_, err := db.Query(fmt.Sprintf(`
	delete from users
	where id=%d
	`, id))

	return err
}

func DeleteAll() error {
	_, err := db.Query(`delete from users`)

	return err
}

func GetByEmail(email string) (user *User, err error) {
	user = &User{}

	query := db.QueryRow(fmt.Sprintf(`
		select * from users
		where email='%s'
	`, email))

	err = query.Scan(&user.Id, &user.Username, &user.Email, &user.Password)

	return user, err
}

func GetByUsername(username string) (user *User, err error) {
	user = &User{}

	query := db.QueryRow(fmt.Sprintf(`
		select * from users
		where username='%s'
	`, username))

	err = query.Scan(&user.Id, &user.Username, &user.Email, &user.Password)

	return user, err
}
