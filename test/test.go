package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "net/http"
)

func makeAPICalls() {
    // Example 1: GET request
    resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1&apiKey=12345")
    if err != nil {
        fmt.Println("Error making GET request:", err)
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading GET response body:", err)
        return
    }
    fmt.Println("GET Response:", string(body))

    // Example 2: POST request
    postData := []byte(`{"title":"foo","body":"bar","userId":1}`)
    resp, err = http.Post("https://jsonplaceholder.typicode.com/posts&apikey=asda32asd", "application/json", bytes.NewBuffer(postData))
    if err != nil {
        fmt.Println("Error making POST request:", err)
        return
    }
    defer resp.Body.Close()

    body, err = ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading POST response body:", err)
        return
    }
    fmt.Println("POST Response:", string(body))
}