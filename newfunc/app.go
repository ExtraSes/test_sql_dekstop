package main

import (
	"context"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// usergovnas структура для таблицы usergovnas
type usergovnas struct {
	ID       uint   `gorm:"primarykey" json:"id"`
	Name     string `gorm:"size:255" json:"name"`
	Sex      string `gorm:"size:10" json:"sex"`
	Sumgavna uint   `json:"sumgavna"`
}

// App struct
type App struct {
	ctx context.Context
	db  *gorm.DB
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	
	// Подключение к базе данных
	dsn := "demid:Gandon345@tcp(192.168.1.92:3306)/testgovna?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("❌ FAILED TO CONNECT DATABASE: %v", err)
		return
	}
	
	a.db = db
	log.Println("✅ DATABASE CONNECTED SUCCESSFULLY")
	
	// Автоматическая миграция
	err = db.AutoMigrate(&usergovnas{})
	if err != nil {
		log.Printf("❌ FAILED TO MIGRATE DATABASE: %v", err)
		return
	}
	log.Println("✅ DATABASE MIGRATED SUCCESSFULLY")
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// GetAllUsers возвращает всех пользователей из базы данных
func (a *App) GetAllUsers() ([]usergovnas, error) {
	var users []usergovnas
	result := a.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// AddUser добавляет нового пользователя в базу данных
func (a *App) AddUser(name, sex string, sumgavna uint) (*usergovnas, error) {
	user := usergovnas{
		Name:     name,
		Sex:      sex,
		Sumgavna: sumgavna,
	}
	
	result := a.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	
	return &user, nil
}
