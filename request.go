package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"time"
)


func getRequest(url string) (output string, err error) {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

    resp, err := client.Get("https://www.thecocktaildb.com/api/json/v1/1/random.php") 
    if err != nil { 
        fmt.Println(err) 
        return
    } 
    defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}