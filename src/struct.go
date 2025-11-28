package groupie

import (
	"io"
	"net/http"
)

type Groupe struct {
	id           string
	name         string
	image        string
	members      []string
	creationDate string
	firstAlbum   string
	locations    string
	concertDates string
	relations    string
	isload       int
}

type Data struct {
	nbGroup   int
	listGroup []Groupe
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

	Group1.id = GetToken(string(body), "id")
	if Group1.id != id {
		return nil
	}
	Group1.image = GetToken(string(body), "image")
	Group1.name = GetToken(string(body), "name")
	Group1.members = GetMultiToken(string(body), "members")
	Group1.creationDate = GetToken(string(body), "creationDate")
	Group1.firstAlbum = GetToken(string(body), "firstAlbum")
	Group1.locations = GetToken(string(body), "locations")
	Group1.concertDates = GetToken(string(body), "concertDates")
	Group1.relations = GetToken(string(body), "relations")
	Group1.isload = 1
	return &Group1
}

func LoadGroupResum() *Data {
	var data Data
	data.nbGroup = 0

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
		group.id = GetToken(string(body), "id")
		if group.id != id {
			break
		}
		group.image = GetToken(string(body), "image")
		group.name = GetToken(string(body), "name")
		data.listGroup = append(data.listGroup, group)
		data.nbGroup++
	}
	return &data
}

// func Maintest() {
// 	data := LoadGroupResum()

// 	println(data.nbGroup)
// 	println(data.listGroup[4].name)
// }
