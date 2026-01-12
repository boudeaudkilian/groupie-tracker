package groupie

func lowcase(txt string) string {
	ret := ""

	for _, v := range txt {
		if v >= 'A' && v <= 'Z' {
			v += 'a' - 'A'
		}
		ret += string(v)
	}
	return ret
}

func ispart(groupname, txt string) bool {
	cmp := 0
	groupname = lowcase(groupname)
	txt = lowcase(txt)

	for i := 0; i < len(groupname)-(len(txt)-1); i++ {
		for j := 0; j < len(txt); j++ {
			if groupname[i+j] == txt[j] {
				cmp++
			} else {
				break
			}
		}
		if cmp == len(txt) {
			return true
		}
		cmp = 0
	}
	return false
}

func ispartmembers(groupname []string, txt string) bool {
	txt = lowcase(txt)
	if groupname == nil {
		return false
	}
	for m := 0; m < len(groupname); m++ {
		cmp := 0
		members := lowcase(groupname[m])

		for i := 0; i < len(members)-(len(txt)-1); i++ {
			for j := 0; j < len(txt); j++ {
				if members[i+j] == txt[j] {
					cmp++
				} else {
					break
				}
			}
			if cmp == len(txt) {
				return true
			}
			cmp = 0
		}
	}
	return false
}

func Search(txt string, stru Data) *Data {
	var newstru Data

	for i := 0; i < stru.NbGroup; i++ {
		if ispart(stru.ListGroup[i].Name, txt) {
			newstru.ListGroup = append(newstru.ListGroup, stru.ListGroup[i])
			newstru.NbGroup++
		} else if ispartmembers(stru.ListGroup[i].Members, txt) {
			newstru.ListGroup = append(newstru.ListGroup, stru.ListGroup[i])
			newstru.NbGroup++
		} else if ispart(stru.ListGroup[i].FirstAlbum, txt) {
			newstru.ListGroup = append(newstru.ListGroup, stru.ListGroup[i])
			newstru.NbGroup++
		} else if ispart(stru.ListGroup[i].CreationDate, txt) {
			newstru.ListGroup = append(newstru.ListGroup, stru.ListGroup[i])
			newstru.NbGroup++
		}
	}

	return &newstru
}

// func test(groupname, txt string) {
// 	if ispart(groupname, txt) {
// 		fmt.Println("1")
// 	} else {
// 		fmt.Println("0")
// 	}
// }
// func Maintest(){
// 	test("oui", "ou")
// 	test("non", "ou")
// 	test("oui", "ouii")
// 	test("oui", "ui")
// 	test("oui", "U")
// }
