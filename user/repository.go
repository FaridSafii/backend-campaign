package user

import "gorm.io/gorm"

//interface dalam go merupakan depedensi antar Repository dan Service
type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByID(ID int) (User, error)
	Update(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

//Mengambil data db yang ada dalam main
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

//Membuat function Save untuk struct repository
func (r *repository) Save(user User) (User, error) {
	//Membuat data baru dengan Create
	err := r.db.Create(&user).Error
	//jika error
	if err != nil {
		return user, err
	}
	//jika berhasil
	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User

	err := r.db.Where("email = ?", email).Find(&user).Error
	//jika error
	if err != nil {
		return user, err
	}
	//jika berhasil
	return user, nil

}

func (r *repository) FindByID(ID int) (User, error) {
	var user User
	//Select data where di gorm
	err := r.db.Where("id = ?", ID).Find(&user).Error
	//jika error
	if err != nil {
		return user, err
	}
	//jika berhasil
	return user, nil
}
func (r *repository) Update(user User) (User, error) {
	//Save di gorm
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
