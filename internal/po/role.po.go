package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey; type:int; autoIncrement; not null"`
	RoleName string `gorm:"unique; column:role_name; type:varchar(255); not null"`
	RoleNote string `gorm:"column:role_note; type:text"`
}

func (r *Role) TableName() string {
	return "go_db_role"
}
