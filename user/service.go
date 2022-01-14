package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// melakukan mapping dari struct input ke struct user
//Service punya depedensi (membutuhkan bantuan dari) Repository
type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
}

//mapping input ke struct user
//simpan struct User melalui repository
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	//parsing repository dari struct service
	//parsing parameter repository pada object service
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	//membuat object dari User struct
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	//Generate password to hash
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	//menuju repository.go ke dalam fungsi Save
	newUser, err := s.repository.Save(user)
	if err != nil {
		return user, err
	}
	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("No user found on that email")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil
}
