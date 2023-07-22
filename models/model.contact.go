package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ModelContact struct {
	ID        string    `json:"id"         gorm:"primary_key"`
	Realname  string    `json:"real_name"  gorm:"type:varchar; not null"`
	Whatsapp  uint64    `json:"whatsapp"   gorm:"type:bigint; unique; not null"`
	Email     string    `json:"email"      gorm:"type:text; not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *ModelContact) BeforeCreate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.CreatedAt = time.Now()
	return nil
}

func (m *ModelContact) BeforeUpdate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.UpdatedAt = time.Now()
	return nil
}
