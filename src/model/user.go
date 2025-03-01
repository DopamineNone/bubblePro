package model

type User struct {
	ID       int64  `gorm:"primaryKey"`
	Username string `gorm:"unique; not null"`
	Password string `gorm:"type:char(96); not null"`
}
