package repositories

import (
	"project1/app/models"
	"project1/app/repositories/base"
)

type EventReminderRepository interface {
	base.AddBase[models.EventReminder]
	base.DeleteBase[int, models.EventReminder]
	base.DetailBase[int, models.EventReminder]
}
