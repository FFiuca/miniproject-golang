package repositories

import (
	"project1/app/models"
	"project1/app/repositories/base"
)

type StatusRepository interface {
	base.AddBase[models.Status]
	base.UpdateBase[int, models.Status]
	base.DeleteBase[int, models.Status]
	base.DetailBase[int, models.Status]
}
