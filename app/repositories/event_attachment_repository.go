package repositories

import (
	"project1/app/models"
	"project1/app/repositories/base"
)

type EventAttachmentRepository interface {
	base.AddBase[models.EventAttachment]
	base.DeleteBase[int, models.EventAttachment]
	base.DetailBase[int, models.EventAttachment]
}
