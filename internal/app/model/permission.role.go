package model

type PermissionRole struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	RoleID       uint       `gorm:"not null"`
	PermissionID uint       `gorm:"not null"`
	Role         Role       `gorm:"foreignKey:RoleID;constraint:OnDelete:CASCADE"`
	Permission   Permission `gorm:"foreignKey:PermissionID;constraint:OnDelete:CASCADE"`
}

func (PermissionRole) TableName() string {
	return "role_permissions"
}
