package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-memdb"
)

var schema *memdb.DBSchema
var db *memdb.MemDB

func main() {
	// Database initialization
	schema = &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"user": &memdb.TableSchema{
				Name: "user",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "UserName"},
					},
				},
			},
		},
	}

	var err error
	db, err = memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}

	txn := db.Txn(true)
	users := []*User{
		&User{"aaa", "aaa@example.com", true},
		&User{"bbb", "bbb@example.com", false},
		&User{"ccc", "ccc@example.com", true},
	}

	for _, u := range users {
		if err := txn.Insert("user", u); err != nil {
			txn.Abort()
			panic(err)
		}
	}
	txn.Commit()

	// Routes declaration
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	userRoutes := api.PathPrefix("/user").Subrouter()
	userRoutes.HandleFunc("", getAll).Methods(http.MethodGet)
	userRoutes.HandleFunc("/{userID}", getUser).Methods(http.MethodGet)
	userRoutes.HandleFunc("", createUser).Methods(http.MethodPost)
	http.ListenAndServe(":8080", r)
}
