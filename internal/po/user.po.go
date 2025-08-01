package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"primaryKey; type:char(36); index:idx_uuid"`
	Username string    `gorm:"unique"`
	IsActive bool      `gorm:"default:true"`
	Roles    []Role    `gorm:"many2many:user_roles;"`
}

func (u *User) TableName() string {
	return "go_db_user"
}
