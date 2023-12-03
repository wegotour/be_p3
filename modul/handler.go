package modul

import (
	"encoding/json"
	"net/http"
	"os"

	model "github.com/wegotour/be_p3/model"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	Responsed  model.Credential
	Response   model.TicketResponse
	datauser   model.User
	dataticket model.Ticket
)

// user
func GCFHandlerGetAllUser(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	Responsed.Status = false

	userlist, err := GetAllUser(mconn, collectionname)
	if err != nil {
		Responsed.Message = err.Error()
		return GCFReturnStruct(Responsed)
	}

	Responsed.Status = true
	Responsed.Message = "Get User Success"
	Responsed.Data = userlist

	return GCFReturnStruct(Responsed)
}

func GCFHandlerGetUserByUsername(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	Responsed.Status = false

	username := r.URL.Query().Get("username")
	if username == "" {
		Responsed.Message = "Missing 'username' parameter in the URL"
		return GCFReturnStruct(Responsed)
	}

	datauser.Username = username

	user, err := GetUserFromUsername(mconn, collectionname, username)
	if err != nil {
		Responsed.Message = "Error retrieving user data: " + err.Error()
		return GCFReturnStruct(Responsed)
	}

	Responsed.Status = true
	Responsed.Message = "Hello user"
	Responsed.Data = []model.User{user}

	return GCFReturnStruct(Responsed)
}

func GCFHandlerRegister(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	Responsed.Status = false

	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		Responsed.Message = "error parsing application/json: " + err.Error()
	}

	err = Register(mconn, collectionname, datauser)
	if err != nil {
		Responsed.Message = err.Error()
		return GCFReturnStruct(Responsed)
	}

	Responsed.Status = true
	Responsed.Message = "Register success"

	return GCFReturnStruct(Responsed)
}

func GCFHandlerLogIn(PASETOPRIVATEKEYENV, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	Responsed.Status = false

	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		Responsed.Message = "error parsing application/json: " + err.Error()
	}

	user, _, err := LogIn(mconn, collectionname, datauser)
	if err != nil {
		Responsed.Message = err.Error()
		return GCFReturnStruct(Responsed)
	}

	Responsed.Status = true
	tokenstring, err := watoken.Encode(user.UID, os.Getenv(PASETOPRIVATEKEYENV))
	if err != nil {
		Responsed.Message = "Gagal Encode Token :" + err.Error()

	} else {
		Responsed.Message = "Selamat Datang " + user.Username
		Responsed.Token = tokenstring
	}

	return GCFReturnStruct(Responsed)
}

func GCFHandlerUpdateUser(PASETOPUBLICKEY, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	Responsed.Status = false

	token := r.Header.Get("Authorization")
	if token == "" {
		Responsed.Message = "error parsing application/json1:"
		return GCFReturnStruct(Responsed)
	}

	_, err := watoken.Decode(os.Getenv(PASETOPUBLICKEY), token)
	if err != nil {
		Responsed.Message = "error parsing application/json2:" + err.Error() + ";" + token
		return GCFReturnStruct(Responsed)
	}

	id := r.URL.Query().Get("_id")
	if id == "" {
		Responsed.Message = "Missing '_id' parameter in the URL"
		return GCFReturnStruct(Responsed)
	}

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		Responsed.Message = "Invalid '_id' parameter in the URL"
		return GCFReturnStruct(Responsed)
	}

	datauser.ID = ID

	err = json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		Responsed.Message = "error parsing application/json: " + err.Error()
	}

	user, _, err := UpdateUser(mconn, collectionname, datauser)
	if err != nil {
		Responsed.Message = "Error updating user data: " + err.Error()
		return GCFReturnStruct(Responsed)
	}

	Responsed.Status = true
	Responsed.Message = "Update success " + user.Username
	Responsed.Data = []model.User{user}

	return GCFReturnStruct(Responsed)
}

func GCFHandlerChangePassword(PASETOPUBLICKEY, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	Responsed.Status = false

	token := r.Header.Get("Authorization")
	if token == "" {
		Responsed.Message = "error parsing application/json1:"
		return GCFReturnStruct(Responsed)
	}

	_, err := watoken.Decode(os.Getenv(PASETOPUBLICKEY), token)
	if err != nil {
		Responsed.Message = "error parsing application/json2:" + err.Error() + ";" + token
		return GCFReturnStruct(Responsed)
	}

	username := r.URL.Query().Get("username")
	if username == "" {
		Responsed.Message = "Missing 'username' parameter in the URL"
		return GCFReturnStruct(Responsed)
	}

	datauser.Username = username

	err = json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		Responsed.Message = "error parsing application/json: " + err.Error()
	}

	user, _, err := ChangePassword(mconn, collectionname, datauser)
	if err != nil {
		Responsed.Message = "Error changing password: " + err.Error()
		return GCFReturnStruct(Responsed)
	}

	Responsed.Status = true
	Responsed.Message = "Password change success for user " + user.Username
	Responsed.Data = []model.User{user}

	return GCFReturnStruct(Responsed)
}

func GCFHandlerDeleteUser(PASETOPUBLICKEY, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	Responsed.Status = false

	token := r.Header.Get("Authorization")
	if token == "" {
		Responsed.Message = "error parsing application/json1:"
		return GCFReturnStruct(Responsed)
	}

	_, err := watoken.Decode(os.Getenv(PASETOPUBLICKEY), token)
	if err != nil {
		Responsed.Message = "error parsing application/json2:" + err.Error() + ";" + token
		return GCFReturnStruct(Responsed)
	}

	username := r.URL.Query().Get("username")
	if username == "" {
		Responsed.Message = "Missing 'username' parameter in the URL"
		return GCFReturnStruct(Responsed)
	}

	datauser.Username = username

	err = json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		Responsed.Message = "error parsing application/json: " + err.Error()
	}

	_, err = DeleteUser(mconn, collectionname, datauser)
	if err != nil {
		Responsed.Message = "Error deleting user data: " + err.Error()
		return GCFReturnStruct(Responsed)
	}

	Responsed.Status = true
	Responsed.Message = "Delete user " + datauser.Username + " success"

	return GCFReturnStruct(Responsed)
}

// ticket
func GCFHandlerGetTicket(PASETOPUBLICKEY, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	Response.Status = false

	token := r.Header.Get("Authorization")
	if token == "" {
		Response.Message = "error parsing application/json1:"
		return GCFReturnStruct(Response)
	}

	_, err := watoken.Decode(os.Getenv(PASETOPUBLICKEY), token)
	if err != nil {
		Response.Message = "error parsing application/json2:" + err.Error() + ";" + token
		return GCFReturnStruct(Response)
	}

	id := r.URL.Query().Get("_id")
	if id == "" {
		Response.Message = "Missing '_id' parameter in the URL"
		return GCFReturnStruct(Response)
	}

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		Response.Message = "Invalid '_id' parameter in the URL"
		return GCFReturnStruct(Response)
	}

	// err = json.NewDecoder(r.Body).Decode(&dataticket)
	// if err != nil {
	// 	Response.Message = "error parsing application/json3: " + err.Error()
	// 	return GCFReturnStruct(Response)
	// }

	ticket, err := GetTicketFromID(mconn, collectionname, ID)
	if err != nil {
		Response.Message = err.Error()
		return GCFReturnStruct(Response)
	}

	Response.Status = true
	Response.Message = "Get ticket success"
	Response.Data = []model.Ticket{ticket}

	return GCFReturnStruct(Response)
}

func GCFHandlerInsertTicket(PASETOPUBLICKEY, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	Response.Status = false

	token := r.Header.Get("Authorization")
	if token == "" {
		Response.Message = "error parsing application/json1:"
		return GCFReturnStruct(Response)
	}

	userInfo, err := watoken.Decode(os.Getenv(PASETOPUBLICKEY), token)
	if err != nil {
		Response.Message = "error parsing application/json2:" + err.Error() + ";" + token
		return GCFReturnStruct(Response)
	}

	err = json.NewDecoder(r.Body).Decode(&dataticket)
	if err != nil {
		Response.Message = "error parsing application/json3: " + err.Error()
		return GCFReturnStruct(Response)
	}

	_, err = InsertTicket(mconn, collectionname, dataticket, userInfo.Id)
	if err != nil {
		Response.Message = err.Error()
		return GCFReturnStruct(Response)
	}

	Response.Status = true
	Response.Message = "Insert ticket success for " + dataticket.Title
	Response.Data = []model.Ticket{dataticket}

	return GCFReturnStruct(Response)
}

func GCFHandlerUpdateTicket(PASETOPUBLICKEY, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	Response.Status = false

	token := r.Header.Get("Authorization")
	if token == "" {
		Response.Message = "error parsing application/json1:"
		return GCFReturnStruct(Response)
	}

	_, err := watoken.Decode(os.Getenv(PASETOPUBLICKEY), token)
	if err != nil {
		Response.Message = "error parsing application/json2:" + err.Error() + ";" + token
		return GCFReturnStruct(Response)
	}

	id := r.URL.Query().Get("_id")
	if id == "" {
		Response.Message = "Missing '_id' parameter in the URL"
		return GCFReturnStruct(Response)
	}

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		Response.Message = "Invalid '_id' parameter in the URL"
		return GCFReturnStruct(Response)
	}

	dataticket.ID = ID

	err = json.NewDecoder(r.Body).Decode(&dataticket)
	if err != nil {
		Response.Message = "error parsing application/json3: " + err.Error()
		return GCFReturnStruct(Response)
	}

	ticket, _, err := UpdateTicket(mconn, collectionname, dataticket)
	if err != nil {
		Response.Message = err.Error()
		return GCFReturnStruct(Response)
	}

	Response.Status = true
	Response.Message = "Update ticket success"
	Response.Data = []model.Ticket{ticket}

	return GCFReturnStruct(Response)
}

func GCFHandlerDeleteTicket(PASETOPUBLICKEY, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	Response.Status = false

	token := r.Header.Get("Authorization")
	if token == "" {
		Response.Message = "error parsing application/json1:"
		return GCFReturnStruct(Response)
	}

	_, err := watoken.Decode(os.Getenv(PASETOPUBLICKEY), token)
	if err != nil {
		Response.Message = "error parsing application/json2:" + err.Error() + ";" + token
		return GCFReturnStruct(Response)
	}

	id := r.URL.Query().Get("_id")
	if id == "" {
		Response.Message = "Missing '_id' parameter in the URL"
		return GCFReturnStruct(Response)
	}

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		Response.Message = "Invalid '_id' parameter in the URL"
		return GCFReturnStruct(Response)
	}

	// err = json.NewDecoder(r.Body).Decode(&dataticket)
	// if err != nil {
	// 	Response.Message = "error parsing application/json3: " + err.Error()
	// 	return GCFReturnStruct(Response)
	// }

	_, err = DeleteTicket(mconn, collectionname, ID)
	if err != nil {
		Response.Message = err.Error()
		return GCFReturnStruct(Response)
	}

	Response.Status = true
	Response.Message = "Delete ticket success"

	return GCFReturnStruct(Response)
}

func GCFReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}
