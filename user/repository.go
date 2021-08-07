package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
}

// huruf kecil pada repository, bisa menandakan encapsulation. tidak bisa di akses dari package lain
type repository struct {
	db *gorm.DB
}

// buat function NewRepository supaya ketika akses repository huruf kecil sudah ada isi datanya.
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
