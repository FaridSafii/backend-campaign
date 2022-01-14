package user

import "golang.org/x/crypto/bcrypt"

// melakukan mapping dari struct input ke struct user
//Service punya depedensi (membutuhkan bantuan dari) Repository
type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
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
