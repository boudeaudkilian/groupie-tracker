package groupie


func Listegroupes() []Groupe {
	var groups Data
	groups = *LoadGroupResum()
	return groups.ListGroup
}

func Triealpha(groups []Groupe, sortby string) []Groupe {
	n := len(groups)
	sorted := make([]Groupe, n)
	used := make([]bool, n)

	for i := 0; i < n; i++ {
		minIndex := -1
		for j := 0; j < n; j++ {
			if used[j] {
				continue
			}
			if minIndex == -1 {
				minIndex = j
			} else {
				if sortby == "name" {
					if groups[j].Name < groups[minIndex].Name {
						minIndex = j
					}
				} else if sortby == "date" {
					if groups[j].CreationDate < groups[minIndex].CreationDate {
						minIndex = j
					}
				}
			}
		}
		sorted[i] = groups[minIndex]
		used[minIndex] = true
	}
	return sorted
}

func Triealpharivers(groups []Groupe, sortby string) []Groupe {
	n := len(groups)
	sorted := make([]Groupe, n)
	used := make([]bool, n)

	for i := 0; i < n; i++ {
		minIndex := -1
		for j := 0; j < n; j++ {
			if used[j] {
				continue
			}
			if minIndex == -1 {
				minIndex = j
			} else {
				if sortby == "name" {
					if groups[j].Name > groups[minIndex].Name {
						minIndex = j
					}
				} else if sortby == "date" {
					if groups[j].CreationDate > groups[minIndex].CreationDate {
						minIndex = j
					}
				}
			}
		}
		sorted[i] = groups[minIndex]
		used[minIndex] = true
	}
	return sorted
}