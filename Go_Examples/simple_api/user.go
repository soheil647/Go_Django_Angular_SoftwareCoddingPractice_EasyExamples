package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var db *gorm.DB
var err error

type User struct{
	gorm.Model
	Name string
	Email string
}

func InitialMigration(){
	db, err = gorm.Open("mysql", "root:@/hellogo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Create DB")
		panic(err.Error())
	}
	defer db.Close()

	db.AutoMigrate(&User{})
}

func AllUsers(w http.ResponseWriter, r * http.Request){
	db, err = gorm.Open("mysql", "root:@/hellogo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("All Users")
		panic(err.Error())
	}
	defer db.Close()

	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func NewUser(w http.ResponseWriter, r *http.Request){
	db, err = gorm.Open("mysql", "root:@/hellogo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("All Users")
		panic(err.Error())
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{Name:name,Email:email})

	fmt.Fprintf(w,"New User Created")
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	db, err = gorm.Open("mysql", "root:@/hellogo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("All Users")
		panic(err.Error())
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w,"User Deleted")
}
func UpdateUser(w http.ResponseWriter, r *http.Request){
	db, err = gorm.Open("mysql", "root:@/hellogo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("All Users")
		panic(err.Error())
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User
	db.Where("name = ?", name).Find(&user)

	user.Email = email
	db.Save(&user)
	fmt.Fprintf(w, "Update Secsusfully")


}

