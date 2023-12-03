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
	data.ID = primitive.NewObjectID()
	data.Email = "dapskuy@gmail.com"
	data.Username = "dapskuy"
	data.Role = "user"
	data.Password = "kepodah"

	err := modul.Register(mconn, "user", data)
	if err != nil {
		t.Errorf("Error registering user: %v", err)
	} else {
		fmt.Println("Register success", data)
	}
}

// test login
func TestLogIn(t *testing.T) {
	var userdata model.User
	userdata.Username = "dapskuy"
	userdata.Password = "kepodah"
	user, status, err := modul.LogIn(mconn, "user", userdata)
	fmt.Println("Status", status)
	if err != nil {
		t.Errorf("Error logging in user: %v", err)
	} else {
		fmt.Println("Login success", user)
	}
}

func TestUpdateUser(t *testing.T) {
	var data model.User
	data.Email = "dapskuy@gmail.com"
	data.Username = "dapskuy"
	data.Role = "admin"

	data.Password = "kepodah" // password tidak diubah

	id, err := primitive.ObjectIDFromHex("654a6513226d8ad245cd01ff")
	data.ID = id
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
	data.Email = "dapskuy@gmail.com" // email tidak diubah
	data.Username = "dapskuy"        // username tidak diubah
	data.Role = "admin"              // role tidak diubah

	data.Password = "kepodah"

	// username := "dapskut123"

	_, status, err := modul.ChangePassword(mconn, "user", data)
	fmt.Println("Status", status)
	if err != nil {
		t.Errorf("Error updateting document: %v", err)
	} else {
		fmt.Println("Password berhasil diubah dengan username:", data.Username)
	}
}

// test delete user
func TestDeleteUser(t *testing.T) {
	var data model.User
	data.Username = "admin"

	status, err := modul.DeleteUser(mconn, "user", data)
	fmt.Println("Status", status)
	if err != nil {
		t.Errorf("Error deleting document: %v", err)
	} else {
		fmt.Println("Delete user" + data.Username + "success")
	}
}

func TestGetUserFromID(t *testing.T) {
	id, _ := primitive.ObjectIDFromHex("6539d6c46700af5da789a678")
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

func TestGetUserFromEmail(t *testing.T) {
	anu, _ := modul.GetUserFromEmail(mconn, "user", "tejo@gmail.com")
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
	data.Title = "Pergi ke sana"
	data.Description = "membeli itu ini"
	data.Deadline = "02/02/2021"
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
