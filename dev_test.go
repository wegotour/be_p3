package be_p3

import (
	"fmt"
	"testing"

	model "github.com/wegotour/be_p3/model"
	modul "github.com/wegotour/be_p3/modul"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mconn = SetConnection("MONGOSTRING", "wegotour")

// user
func TestRegister(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "wegotour")
	var userdata User
	userdata.Email = "dapskuy@gmail.com"
	userdata.Username = "dapskuy"
	userdata.Role = "admin"
	userdata.Password = "kepoah"

	nama := InsertUser(mconn, "user", userdata)
	fmt.Println(nama)
}

// test login
func TestLogIn(t *testing.T) {
	var userdata model.User
	userdata.Username = "dapskuy"
	userdata.Password = "kepoah"
	user, status, err := modul.LogIn(mconn, "user", userdata)
	fmt.Println("Status", status)
	if err != nil {
		t.Errorf("Error logging in user: %v", err)
	} else {
		fmt.Println("Login success", user)
	}
}

// test change password
func TestChangePassword(t *testing.T) {
	username := "dapskuy"
	oldpassword := "kepoah"
	newpassword := "kepo"

	var userdata model.User
	userdata.Username = username
	userdata.Password = newpassword

	userdata, status, err := modul.ChangePassword(mconn, "user", username, oldpassword, newpassword)
	fmt.Println("Status", status)
	if err != nil {
		t.Errorf("Error changing password: %v", err)
	} else {
		fmt.Println("Password change success for user", userdata)
	}
}

// test delete user
func TestDeleteUser(t *testing.T) {
	username := "dap_skuy"

	err := modul.DeleteUser(mconn, "user", username)
	if err != nil {
		t.Errorf("Error deleting user: %v", err)
	} else {
		fmt.Println("Delete user success")
	}

	_, err = modul.GetUserFromUsername(mconn, "user", username)
	if err == nil {
		fmt.Println("Data masih ada")
	}
}

func TestGetUserFromID(t *testing.T) {
	id, _ := primitive.ObjectIDFromHex("")
	anu, _ := modul.GetUserFromID(mconn, "user", id)
	fmt.Println(anu)
}

func TestGetUserFromUsername(t *testing.T) {
	anu, err := modul.GetUserFromUsername(mconn, "user", "dapskuy")
	if err != nil {
		t.Errorf("Error getting user: %v", err)
		return
	}
	fmt.Println(anu)
}

func TestGetAllUser(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "wegotour")
	anu := modul.GetAllUser(mconn, "user")
	fmt.Println(anu)
}

// ticket
func TestInsertTicket(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "wegotour")
	var ticketdata model.Ticket
	ticketdata.Title = "Perjalanan"
	ticketdata.Description = "pergi ke bali"
	ticketdata.IsDone = true

	nama, err := modul.InsertTicket(mconn, "ticket", ticketdata)
	if err != nil {
		t.Errorf("Error inserting ticket: %v", err)
	}
	fmt.Println(nama)
}

func TestGetTicketFromID(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "wegotour")
	id, _ := primitive.ObjectIDFromHex("653e02ab28597c2c37171d44")
	anu := modul.GetTicketFromID(mconn, "ticket", id)
	fmt.Println(anu)
}

func TestGetTicketList(t *testing.T) {
	anu := modul.GetTicketList(mconn, "user")
	fmt.Println(anu)
}
