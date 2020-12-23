package mckcat

//Cat struct
type Cat struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	IsDeleted bool   `json:"isDeleted"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

var dbCat = make([]Cat, 0)

//GetAll func
func GetAll() (cat []Cat) {
	return dbCat
}

//GetByID func
func GetByID(id int) (getCat Cat) {
	for _, data := range dbCat {
		if data.ID == id {
			getCat = data
			break
		}
	}
	return getCat
}

//Create func
func Create(cat Cat) (result string) {
	var duplicate bool
	for _, data := range dbCat {
		if data.Name == cat.Name {
			duplicate = true
			break
		}
	}
	if duplicate {
		return "Cat Name must unique"
	}

	if len(dbCat) == 0 {
		cat.ID = 1
	} else {
		cat.ID = dbCat[len(dbCat)-1].ID + 1
	}
	dbCat = append(dbCat, cat)
	return "Created"
}

//Update func
func Update(cat Cat, id int) (result string) {
	getCatIndex := -1

	for index, data := range dbCat {
		if cat.Name == data.Name && id != data.ID {
			getCatIndex = -2
			break
		}
		if id == data.ID {
			getCatIndex = index
		}
	}

	if getCatIndex == -1 {
		return "Cat not found"
	} else if getCatIndex == -2 {
		return "Cat Name must unique"
	} else {
		dbCat[getCatIndex] = cat
		return "Updated"
	}
}

//Delete func
func Delete(id int) (result string) {
	getCatIndex := -1
	for index, data := range dbCat {
		if id == data.ID {
			getCatIndex = index
			break
		}
	}
	if getCatIndex == -1 {
		return "Cat not found"
	}
	dbCat = append(dbCat[:getCatIndex], dbCat[getCatIndex+1:]...)
	return "Deleted"
}
