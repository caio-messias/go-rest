package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	userID := pathParams["userID"]

	txn := db.Txn(false)
	raw, err := txn.First("user", "id", userID)

	if raw == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf(`{"message": "could not find user %s"}`, userID)))
		return
	}
	if err != nil {
		log.Printf("error fetching user %s: %s", userID, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"message": "error fetching user %s"}`, userID)))
		return
	}

	user := raw.(*User)
	userJSON, err := json.Marshal(user)
	if err != nil {
		log.Printf("error parsing user %s: %s", userID, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"message": "could not fetch user %s"}`, userID)))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)
}

func getAll(w http.ResponseWriter, r *http.Request) {
	var users []*User

	txn := db.Txn(false)
	iter, err := txn.Get("user", "id")
	if err != nil {
		log.Println("error fetching all users")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "could not fetch users"}`))
	}

	for row := iter.Next(); row != nil; row = iter.Next() {
		u := row.(*User)
		if u.IsEnabled {
			users = append(users, u)
		}
	}

	usersJSON, err := json.Marshal(users)
	if err != nil {
		log.Println("error parsing users")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "could not fetch users"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(usersJSON)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser UserPayload
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&newUser)

	log.Printf("inserting new user %s %s %t", newUser.UserName, newUser.Email, newUser.IsEnabled)

	userModel := CreateUserFromPayload(newUser)
	txn := db.Txn(true)
	if err := txn.Insert("user", userModel); err != nil {
		txn.Abort()
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"message": "could not insert user %s"}`, newUser.UserName)))
		return
	}
	txn.Commit()

	userJSON, err := json.Marshal(userModel)
	if err != nil {
		log.Printf("error parsing user %s: %s", newUser.UserName, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"message": "error parsing user %s"}`, newUser.UserName)))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Location", fmt.Sprintf(`/user/%s`, userModel.UserName))
	w.Write(userJSON)
}
