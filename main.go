package main

import (
    "encoding/json"
    "fmt"
)

type Language struct {
    Id   int
    Name string
}


type Result struct {
    Drinks []Language
}

func main() {
    text := "{\"drinks\": [{\"Id\": 100, \"Name\": \"Go\"}]}"

    bytes := []byte(text)


    var result Result
    json.Unmarshal(bytes, &result)

    fmt.Println(result)
}