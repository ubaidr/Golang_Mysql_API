package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"

	//"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)
var db *gorm.DB
var err error

type user struct {

	UserId     int  `gorm:"primary_key" json:"id"`
	Name     string `gorm:"size:100;not null" json:"name;omitempty"`
	Email     string `gorm:"size:100;not null;unique" json:"email"`
	Password string   `gorm:"size:100;notnull" json:"password"`
	PhNum    string  `json:"phone"`
	CompanyName string `json:"company_name"`
	Aid int

}

type Account struct {
	AccId     int  `gorm:"primary_key" json:"id"`
	AccType     string `gorm:"size:100;not null" json:"type"`
	AccountName   string  `gorm:"size:100;not null;unique" json:"accname"`
	Users      [] user `gorm:"ForeignKey:Aid"`
}

var userobj= &user{}
var accobj= &Account{}

func main() {
	fmt.Println("Running...")

	path:=os.Getenv("DB_CONNECTION")
	db, err = gorm.Open("mysql", path)

	if err != nil {
		panic("failed to connect database")
	}

	db.LogMode(true)

	defer db.Close()
	db.AutoMigrate(&Account{}, &user{})
	db.Model(&user{}).AddForeignKey("aid", "accounts(acc_id)", "CASCADE", "CASCADE")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowOrigins: []string{"*"},
	}))

	e.POST("/players", CreateUser)
	e.GET("/players/email", Login)
	e.POST("/players/acc", CreateAccount)
	e.Logger.Fatal(e.Start(":8081"))
}

