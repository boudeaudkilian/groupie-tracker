package main

import (
	"fmt"
	"io"
	"net/http"
)

func gettoken(buffer, arg string) string {
	box := ""
	ret := ""
	isword := 0
	isgoodword := 0

	for _, v := range buffer {
		if v == '"' {
			isword = (isword + 1) % 2
			box = ""
			if isgoodword == 1 {
				break
			}
			if isgoodword > 1 {
				isgoodword -= 1
			}
		} else if isword == 1 {
			if isgoodword == 1 {
				ret += string(v)
			} else {
				box += string(v)
				if box == arg {//&& buffer[i + 1] == '"' {
					isgoodword = 3
				}
			}
		} 
	}
	return ret
}

func main() {
	url := "https://groupietrackers.herokuapp.com/api/artists"

	req, _ := http.NewRequest("GET", url + "/1", nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	img := gettoken(string(body), "locations")

	fmt.Println(img)
}
