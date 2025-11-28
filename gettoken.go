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
		if isgoodword == 1 {
			if v == ',' || v == '['{
				break
			}
			if v != '"' {
				ret += string(v)
			}
		} else if isgoodword == 2 && v == ':' {
			isgoodword = 1
		} else if v == '"' {
			isword = (isword + 1) % 2
			box = ""
		} else if isword == 1 {
			box += string(v)
			if box == arg {
				isgoodword = 2
			}
		}
	}
	return ret
}

func main() {
	url := "https://groupietrackers.herokuapp.com/api/artists"

	req, _ := http.NewRequest("GET", url+"/1", nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	tok := gettoken(string(body), "name")

	fmt.Println(tok)
}
