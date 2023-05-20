package models

type User struct {
	ID       int    `gorm:"primaryKey" json:"id_user"`
	LevelID  int    `json:"id_level"`
	Level    Level  `gorm:"foreignKey:LevelID" json:"level"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
