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
