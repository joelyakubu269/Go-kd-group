package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "https://api.github.com/repos/joelyakubu269/Go-kd-group"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusOK:
		fmt.Println("repo exists")
	case http.StatusNotFound:
		fmt.Println("Repo does not exist")
	}
	fmt.Println("respStatus:", resp.StatusCode)
}
