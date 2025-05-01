package model

type RoleUser struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	RoleId uint `json:"role_id" gorm:"not nul"`
	UserId uint `json:"user_id" gorm:"not nul"`
	Role   Role `gorm:"foreignKey:RoleId" json:"role"`
	// User   User `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
}

func (RoleUser) TableName() string {
	return "role_users"
}
