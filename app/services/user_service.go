package services

import (
	"fmt"
	"project1/app/models"
	"project1/app/repositories"

	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (svc *UserService) Add(user *models.User) (any, error) {
	add := svc.DB.Create(user)
	if add.Error != nil {
		return nil, add.Error
	}

	return add, nil
}

func (svc *UserService) Update(id *int, user *models.User) (any, error) {
	fmt.Println("masuk update")
	update := svc.DB.Model(user).Where("id = ?", id).Updates(user)
	if update.Error != nil {
		return nil, update.Error
	}

	return update, nil
}

func (svc *UserService) Delete(id *int, user *models.User) error {
	return svc.DB.Where("id = ?", id).Delete(user).Error
}

func (svc *UserService) Detail(id *int, user *models.User) (*gorm.DB, error) {
	data := svc.DB.Preload("Status").Where("id = ?", id).First(user)
	if data.Error != nil {
		return nil, data.Error
	}
	print("detail", &user.Status, data)
	return data, nil
}

// actually return []any must use generic, but next time i will. just temporary
func (svc *UserService) Search(params *map[string]any, user *models.User) (*[]models.User, error) {
	db := svc.DB.Model(user).Preload("Status")
	users := []models.User{}

	if (*params)["email"] != nil && (*params)["email"] != "" {
		db.Where("email = ? ", (*params)["email"])
	}

	if (*params)["created_at"] != nil && (*params)["created_at"] != "" {
		db.Where("created_at = ? ", (*params)["created_at"])
	}

	err := db.Find(&users).Error // can directly without Rows and Next
	// data, err := db.Rows()
	if err != nil {
		return nil, err
	}
	// defer data.Close()

	// for data.Next() {
	// 	temp := models.User{}
	// 	fmt.Println("sini1")
	// 	fmt.Println(temp)
	// 	if err = data.Scan(&temp); err != nil {
	// 		return nil, err
	// 	}
	// 	fmt.Println("sini")

	// 	users = append(users, temp)
	// }

	// convert it to any
	// result := make([]any, len(users))
	// for i, u := range users {
	// 	result[i] = u
	// }

	// return users, nil

	return &users, nil
}

func (svc *UserService) SearchByEmail(email string) (*models.User, error) {
	user := new(models.User)

	err := svc.DB.Model(&user).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil

}

func NewUserService(db *gorm.DB) repositories.UserRepository {
	return &UserService{
		DB: db,
	}
}
