package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"column:uuid; type:varchar(255); not null; index:idx_user_uuid; unique;"`
	Username string    `gorm:"column:user_name; type:varchar(255); not null; unique;"`
	IsActive bool      `gorm:"column:is_active; type:bool; default:true;"`
	Roles    []Role    `gorm:"many2many:user_roles;"`
}
