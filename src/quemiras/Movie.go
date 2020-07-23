package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

func getMovie(response1 http.ResponseWriter, request *http.Request) {
    fmt.Println("Movie request...")
    response, err := http.Get("https://httpbin.org/ip")
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
        response1.Write([]byte(data))
    }
    fmt.Println("Terminating the application...")
}

func getMovieEndpoint(response http.ResponseWriter, request *http.Request){
    // response.Header().Set("content-type", "application/json")
    // params := mux.Vars(request)
    // // id, _ := primitive.ObjectIDFromHex(params["id"])
    // var user User
    // collection := client.Database("quemiras").Collection("users")
    // ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
    // err := collection.FindOne(ctx, User{ID: id}).Decode(&user)
    // if err != nil {
    //     response.WriteHeader(http.StatusInternalServerError)
    //     response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
    //     return
    // }
    // json.NewEncoder(response).Encode(user)

    fmt.Println("Movie request...")
    response, err := http.Get("https://httpbin.org/ip")
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
        response1.Write([]byte(data))
    }
    fmt.Println("Terminating the application...")
}



func getMovie1(response1 http.ResponseWriter, request *http.Request) {
    fmt.Println("Movie request...")
    response, err := http.Get("https://httpbin.org/ip")
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }
    jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
    jsonValue, _ := json.Marshal(jsonData)
    response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }
    fmt.Println("Terminating the application...")
}

