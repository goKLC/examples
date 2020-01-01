package Models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/goKLC/goKLC"
	"time"
)

type Blog struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	UserID    uint      `json:"user_id"gorm:"not null"`
	Title     string    `json:"title" gorm:"type:varchar(255);not null"`
	Slug      string    `json:"slug" gorm:"type:varchar(255);not null"`
	Content   string    `json:"content" gorm:"type:text"`
	Tags      string    `json:"tags" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func init() {
	app := goKLC.GetApp()
	app.DB().AutoMigrate(&Blog{})
}
