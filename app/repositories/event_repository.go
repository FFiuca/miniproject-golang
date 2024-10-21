package repositories

import (
	"project1/app/models"
	"project1/app/repositories/base"
)

type EventRepository interface {
	base.AddBase[models.Event]
	base.UpdateBase[int, models.Event]
	base.DeleteBase[int, models.Event]
	base.DetailBase[int, models.Event]
	base.SearchBase[map[string]any, models.Event, []models.Event]
}
