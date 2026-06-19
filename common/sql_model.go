package common

import (
	"time"
	"github.com/google/uuid"
)

type SqlModel struct {
	Id 			*uuid.UUID `json:"id" gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4();->"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamp;not null"`
	CreatedAt   *time.Time `json:"created_at" gorm:"column:created_at;type:timestamp;not null"`
}