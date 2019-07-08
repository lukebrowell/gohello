package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

func getSample () string {
    response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
    res := ""
    if err != nil {
        res = fmt.Sprintf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        res = fmt.Sprintln(string(data))
    }
    return res
}

func postSample (jsonValue []byte) string {
    response, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
    res := ""
    if err != nil {
        res = fmt.Sprintf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        res = fmt.Sprintln(string(data))
    }
    return res
}

func main() {
    fmt.Println("Starting the app...")
    var sample = getSample()
    fmt.Println(sample)

    jsonData := map[string]string{"firstname": "George", "lastname": "Clooney"}
    jsonValue, _ := json.Marshal(jsonData)
    var sample2 = postSample(jsonValue)
    fmt.Println(sample2)

    fmt.Println("Terminating the app...")
}

func foo() string {
    var str = "brilliant"
    return str
}