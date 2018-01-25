package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

func main() {
    apiUrl := "https://discordapp.com/api/v6"
    resp, err := http.Get(apiUrl)
    check(err)
    body, err := ioutil.ReadAll(resp.Body)
    check(err)
    fmt.Println(string(body))
}

func check(err error) {
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

