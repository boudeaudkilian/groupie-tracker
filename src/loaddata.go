package groupie

func LoadData() *Data {
	var data Data
	data.ListGroup = Listegroupes()
	data.NbGroup = len(data.ListGroup)
	return &data
}