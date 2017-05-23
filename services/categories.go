package services

import (
	"encoding/json"
	"strings"
	"time"
)

// Categories : categories information.
type Categories struct {
	Id          int64
	Category    string
	Description *string
	Deleted_at  *time.Time
}

// AddNewCategory : add new category.
func AddNewCategory(category string, description string) []byte {
	sess := SetupDB()
	var m = Categories{}
	m.Category = category
	m.Description = &description

	_, err := sess.InsertInto("categories").
		Columns("category", "description").
		Record(m).
		Exec()
	CheckErr(err)

	//get id of inserted category ...
	lastInsertedId, err := sess.Select("MAX(id)").
		From("categories").
		ReturnInt64()

	m.Id = lastInsertedId
	b, err := json.Marshal(m)
	CheckErr(err)
	return b
}

// EditCategoryInfo : edit category.
func EditCategoryInfo(categoryId int, category string, description string) {
	sess := SetupDB()

	_, err := sess.Update("categories").
		Set("category", category).
		Set("description", description).
		Set("modified_at", "NOW()").
		Where("id = ?", categoryId).
		Exec()
	CheckErr(err)
}

// DeleteCategory : remove category.
func DeleteCategory(CategoryIds string) {
	ids := strings.Split(CategoryIds, ",")
	sess := SetupDB()

	//deleting multiple categories ======
	for i := 0; i < len(ids); i += 1 {
		_, err := sess.Update("categories").
			Set("deleted_at", "NOW()").
			Where("id = ?", ids[i]).
			Exec()
		CheckErr(err)
	}
	//==============================

}

// DisplayCategories : list down all categories.
func DisplayCategories(allCategories string) []byte { // Display one category ..
	sess := SetupDB()
	categories := []Categories{}
	query := sess.Select("id, category, description").
		From("categories")

	//display all categories or active categories only ...
	if allCategories == "false" {
		query.Where("deleted_at IS NULL").
			LoadStruct(&categories)
	} else {
		query.LoadStruct(&categories)
	}

	b, err := json.Marshal(categories)
	CheckErr(err)
	return b
}
