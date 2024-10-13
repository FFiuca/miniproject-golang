package repositories

import (
	"project1/app/models"
	"project1/app/repositories/base"
)

type EventUserPivotRepository interface {
	base.AddBase[models.EventUserPivot]
	// base.UpdateBase[int, models.EventUserPivot]
	base.DeleteBase[int, models.EventUserPivot]
	base.DetailBase[int, models.EventUserPivot]
	// base.SearchBase[map[string]any, models.EventUserPivot]
	DeleteByEventID(eventId int, model models.EventUserPivot) error
}
