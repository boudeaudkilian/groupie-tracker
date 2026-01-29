package groupie

import (
	"io"
	"net/http"
)

type Groupe struct {
	Id           string
	Name         string
	Image        string
	Members      []string
	CreationDate string
	FirstAlbum   string
	Locations    string
	ConcertDates string
	Relations    string
	Isload       int
}

type Data struct {
	NbGroup   int
	ListGroup []Groupe
}

var (
	url = "https://groupietrackers.herokuapp.com/api/artists"
)

func Itoa(nb int) string {
	if nb <= 0 {
		return ""
	}

	ret := ""
	len := 1
	for ; nb/len != 0; len *= 10 {
	}
	len /= 10
	for ; len != 0; len /= 10 {
		ret += string(rune((nb / len) + '0'))
		nb = nb % len
	}
	return ret
}

func LoadGroup(ids int) *Groupe {
	var Group1 Groupe
	id := Itoa(ids)

	req, _ := http.NewRequest("GET", url+"/"+id, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	Group1.Id = GetToken(string(body), "id")
	if Group1.Id != id {
		return nil
	}
	Group1.Image = GetToken(string(body), "image")
	Group1.Name = GetToken(string(body), "name")
	Group1.Members = GetMultiToken(string(body), "members")
	Group1.CreationDate = GetToken(string(body), "creationDate")
	Group1.FirstAlbum = GetToken(string(body), "firstAlbum")
	Group1.Locations = GetToken(string(body), "locations")
	Group1.ConcertDates = GetToken(string(body), "concertDates")
	Group1.Relations = GetToken(string(body), "relations")
	Group1.Isload = 1
	return &Group1
}

func LoadGroupResum() *Data {
	var data Data
	data.NbGroup = 0

	for i := 1; ; i++ {
		var group Groupe
		id := Itoa(i)
		req, _ := http.NewRequest("GET", url+"/"+id, nil)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil
		}
		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)
		group.Id = GetToken(string(body), "id")
		if group.Id != id {
			break
		}
		group.Image = GetToken(string(body), "image")
		group.Name = GetToken(string(body), "name")
		data.ListGroup = append(data.ListGroup, group)
		data.NbGroup++
	}
	return &data
}

func LoadAll() *Data {
	var data Data
	data.NbGroup = 0

	for i := 1; i < 999; i++ {
		groupe := LoadGroup(i)
		if groupe == nil {
			break
		}
		data.ListGroup = append(data.ListGroup, *groupe)
		data.NbGroup++
	}
	return &data
}

// func PrintListName(data Data) {
// 	for i := 0; i < data.NbGroup; i++ {
// 		fmt.Println(data.ListGroup[i].Name)
// 	}
// 	fmt.Printf("\n")
// }
// func Maintest() {
// 	data := LoadGroupResum()

// 	println(data.nbGroup)
// 	println(data.listGroup[4].name)
// }
