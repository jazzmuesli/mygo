package main

import (
	"fmt"
	"testing"

	"gorm.io/gorm"
)

func prepare_test_database() *gorm.DB {
	db := prepare_database("test.db")
	return db
}
func TestCreatePerson(t *testing.T) {

	db := prepare_test_database()
	result := create_person(db)
	fmt.Print(result.Error)
	assertEqual(t, 1, result.RowsAffected, "")
}

func TestFindById(t *testing.T) {
	db := prepare_test_database()
	result := create_person(db)
	var persons []Person
	db.Find(&persons)
	fmt.Printf("Found %v persons\n", len(persons))
	assertEqual(t, 1, result.RowsAffected, "")
	person := find_person_by_id(db, int(persons[len(persons)-1].ID))
	assertEqual(t, "John", person.FirstName, "")
}

func TestFindByName(t *testing.T) {
	db := prepare_test_database()
	result := create_person(db)
	assertEqual(t, 1, result.RowsAffected, "")
	person := find_person_by_fname(db, "John")
	assertEqual(t, "John", person.FirstName, "")
}
