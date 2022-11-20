package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	FirstName string
	LastName  string
}

func prepare_database(fname string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(fname), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Person{})
	return db
}
func create_person(db *gorm.DB) *gorm.DB {
	// Create
	return db.Create(&Person{FirstName: "John", LastName: "Doe"})
}

func find_person_by_id(db *gorm.DB) Person {
	var person Person
	db.First(&person, 1) // find product with integer primary key

	return person
}

func find_person_by_fname(db *gorm.DB, firstName string) Person {
	var person Person
	db.First(&person, "FirstName = ?", "John") // find product with code D42
	return person
}

func crud_person() {

	db := prepare_database("main.db")
	create_person(db)
	// Read
	person_by_id := find_person_by_id(db)
	person_by_fname := find_person_by_fname(db, "John")

	// Update - update product's price to 200
	db.Model(&person_by_fname).Update("LastName", "Smith")
	// Update - update multiple fields
	db.Model(&person_by_id).Updates(Person{FirstName: "Joe", LastName: "Schmidt"}) // non-zero fields
	person_by_fname = find_person_by_fname(db, "Joe")
	db.Model(&person_by_fname).Updates(map[string]interface{}{"FirstName": "Todd", "LastName": "Roh"})

	// Delete - delete product
	db.Delete(&person_by_fname, 1)
	person_by_fname = find_person_by_fname(db, "Santa")
	fmt.Print(person_by_fname)

}
