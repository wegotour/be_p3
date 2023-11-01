package be_p3

import (
	"fmt"
	"testing"

	model "github.com/wegotour/be_p3/model"
	modul "github.com/wegotour/be_p3/modul"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// user
func TestInsertUser(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "wegotour")
	var userdata User
	userdata.Email = "dapskuy@gmail.com"
	userdata.Username = "dapskuy"
	userdata.Role = "admin"
	userdata.Password = "kepoah"

	nama := InsertUser(mconn, "user", userdata)
	fmt.Println(nama)
}

func TestGetAllUserFromUsername(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "wegotour")
	anu := modul.GetUserFromUsername(mconn, "user", "dapskuy")
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
