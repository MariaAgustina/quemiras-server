package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "time"
    "github.com/gorilla/mux"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)


var client *mongo.Client

func main() {

    fmt.Println("Starting the application...")
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, _ = mongo.Connect(ctx, clientOptions)

    router := mux.NewRouter()

    router.HandleFunc("/users", GetUsersArrayEndpoint).Methods("GET")
    router.HandleFunc("/user/{id}", GetUserEndpoint).Methods("GET")
    router.HandleFunc("/user/{id}", CreateUserEndpoint).Methods("POST")
    //TODO
    // router.HandleFunc("/user/{id}", DeleteUserEndpoint).Methods("DELETE")
    router.HandleFunc("/movie", getMovie).Methods("GET")

    log.Fatal(http.ListenAndServe(":12345", router))
}