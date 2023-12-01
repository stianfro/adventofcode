package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	session, _ := os.ReadFile("~/.config/aocd/token")

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://adventofcode.com/2023/day/"+os.Args[1]+"/input", nil)
	req.Header.Set("Cookie", string(session))
	res, _ := client.Do(req)

	bodyBytes, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(bodyBytes))
}
