package groupie

func GetToken(buffer, arg string) string {
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

func GetMultiToken(buffer, arg string) []string {
	box := ""
	var ret []string
	isword := 0
	isgoodword := 0

	for _, v := range buffer {
		if isgoodword == 1 {
			if v == ']'{
				break
			}
			if v == ',' {
				ret = append(ret, box)
				box = ""
			} else if v != '"' {
				box += string(v)
			}
		} else if isgoodword == 2 && v == '[' {
			isgoodword = 1
			box = ""
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

//test main for fct
// func Maintest() {
// 	url := "https://groupietrackers.herokuapp.com/api/artists"

// 	req, _ := http.NewRequest("GET", url+"/1", nil)
// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return
// 	}
// 	defer res.Body.Close()

// 	body, _ := io.ReadAll(res.Body)

// 	tok := GetToken(string(body), "name")
// 	fmt.Println(tok)

// 	tok2 := GetMultiToken(string(body), "members")
// 	fmt.Println(tok2)
// }
