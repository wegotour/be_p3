package be_p3

import (
	"fmt"
	"testing"

	model "github.com/wegotour/be_p3/model"
	modul "github.com/wegotour/be_p3/modul"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mconn = modul.MongoConnect("MONGOSTRING", "wegotour")

// user
func TestRegister(t *testing.T) {
	var data model.User
	data.Email = "prisya@gmail.com"
	data.Username = "prisya"
	// data.Role = "user"
	data.Password = "secret"
	data.ConfirmPassword = "secret"

	err := modul.Register(mconn, "user", data)
	if err != nil {
		t.Errorf("Error registering user: %v", err)
	} else {
		fmt.Println("Register success", data.Username)
	}
}

// test login
func TestLogIn(t *testing.T) {
	var data model.User
	data.Username = "prisya"
	data.Password = "secret"
	data.Role = "user"

	user, status, err := modul.LogIn(mconn, "user", data)
	fmt.Println("Status", status)
	if err != nil {
		t.Errorf("Error logging in user: %v", err)
	} else {
		fmt.Println("Login success", user)
	}
}

func TestUpdateUser(t *testing.T) {
	var data model.User
	data.Email = "prisyahaura@gmail.com"
	data.Username = "prisya"

	id := "656c3f638442be4a7c185a09"
	ID, err := primitive.ObjectIDFromHex(id)
	data.ID = ID
	if err != nil {
		fmt.Printf("Data tidak berhasil diubah")
	} else {

		_, status, err := modul.UpdateUser(mconn, "user", data)
		fmt.Println("Status", status)
		if err != nil {
			t.Errorf("Error updateting document: %v", err)
		} else {
			fmt.Printf("Data berhasil diubah untuk id: %s\n", id)
		}
	}
}

// test change password
func TestChangePassword(t *testing.T) {
	var data model.User
	data.Password = "secrets"
	data.ConfirmPassword = "secrets"

	username := "prisya"
	data.Username = username

	_, status, err := modul.ChangePassword(mconn, "user", data)
	fmt.Println("Status", status)
	if err != nil {
		t.Errorf("Error updateting document: %v", err)
	} else {
		fmt.Println("Password berhasil diubah dengan username:", username)
	}
}

// test delete user
func TestDeleteUser(t *testing.T) {
	var data model.User
	data.Username = "prisya"

	status, err := modul.DeleteUser(mconn, "user", data)
	fmt.Println("Status", status)
	if err != nil {
		t.Errorf("Error deleting document: %v", err)
	} else {
		fmt.Println("Delete user" + data.Username + "success")
	}
}

func TestGetUserFromID(t *testing.T) {
	id := "656bf30c733cf24a0f73d0a8"
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Errorf("Error converting id to ObjectID: %v", err)
		return
	}

	anu, err := modul.GetUserFromID(mconn, "user", ID)
	if err != nil {
		t.Errorf("Error getting user: %v", err)
		return
	}
	fmt.Println(anu)
}

func TestGetUserFromUsername(t *testing.T) {
	anu, err := modul.GetUserFromUsername(mconn, "user", "prisya")
	if err != nil {
		t.Errorf("Error getting user: %v", err)
		return
	}
	fmt.Println(anu)
}

func TestGetUserFromEmail(t *testing.T) {
	anu, _ := modul.GetUserFromEmail(mconn, "user", "prisya@gmail.com")
	fmt.Println(anu)
}

func TestGetAllUser(t *testing.T) {
	anu, err := modul.GetAllUser(mconn, "user")
	if err != nil {
		t.Errorf("Error getting user: %v", err)
		return
	}
	fmt.Println(anu)
}

// ticket
func TestInsertTicket(t *testing.T) {
	var data model.Ticket
	data.Title = "Artapela"
	data.Description = "mendaki gunung"
	data.Deadline = "12/04/2023"
	// data.IsDone = false

	uid := "0040f398-1200-4f36-8332-6752ab3e55c0"

	id, err := modul.InsertTicket(mconn, "ticket", data, uid)
	if err != nil {
		t.Errorf("Error inserting ticket: %v", err)
	}
	fmt.Println(id)
}

func TestGetTicketFromID(t *testing.T) {
	id, _ := primitive.ObjectIDFromHex("655c4408d06d3d2ddba5d1d7")
	anu, err := modul.GetTicketFromID(mconn, "ticket", id)
	if err != nil {
		t.Errorf("Error getting ticket: %v", err)
		return
	}
	fmt.Println(anu)
}

func TestGetTicketFromUsername(t *testing.T) {
	anu, err := modul.GetTicketFromUsername(mconn, "ticket", "prisya")
	if err != nil {
		t.Errorf("Error getting ticket: %v", err)
		return
	}
	fmt.Println(anu)
}

func TestUpdateTicket(t *testing.T) {
	var data model.Ticket
	data.Title = "Belajar Golang"
	data.Description = "Hari ini belajar golang"
	data.Deadline = "02/02/2021"

	id := "655c5047370b53741a9705d8"
	ID, err := primitive.ObjectIDFromHex(id)
	data.ID = ID
	if err != nil {
		fmt.Printf("Data tidak berhasil diubah")
	} else {

		_, status, err := modul.UpdateTicket(mconn, "ticket", data)
		fmt.Println("Status", status)
		if err != nil {
			t.Errorf("Error updating ticket with id: %v", err)
			return
		} else {
			fmt.Printf("Data berhasil diubah untuk id: %s\n", id)
		}
		fmt.Println(data)
	}
}

func TestDeleteTicket(t *testing.T) {
	id := "655c4408d06d3d2ddba5d1d7"
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Errorf("Error converting id to ObjectID: %v", err)
		return
	} else {

		status, err := modul.DeleteTicket(mconn, "ticket", ID)
		fmt.Println("Status", status)
		if err != nil {
			t.Errorf("Error deleting document: %v", err)
			return
		} else {
			fmt.Println("Delete success")
		}
	}
}
