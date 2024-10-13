package repositories

import (
	"project1/app/models"
	"project1/app/repositories/base"
)

type UserRepository interface {
	base.AddBase[models.User]
	base.UpdateBase[int, models.User]
	base.DeleteBase[int, models.User]
	base.DetailBase[int, models.User]
	base.SearchBase[map[string]any, models.User, []models.User]
}
