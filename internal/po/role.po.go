package po

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID       uint   `gorm:"column:id; type:int; primaryKey; autoIncrement;"`
	RoleName string `gorm:"column:role_name; type:varchar(255); not null; unique;"`
	RoleDesc string `gorm:"column:role_desc; type:text;"`
}

func (r *Role) TableName() string {
	return "roles"
}
