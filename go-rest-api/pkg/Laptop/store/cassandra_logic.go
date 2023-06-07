package store

import (
	"be-training/go-rest-api/pkg/Laptop/model"
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"
)

const (
	TABLE        = "eeshan.Laptoo"
	FIELD_ID     = "id"
	FIELD_BRAND  = "brand"
	FIELD_MODEL  = "model"
	FIELD_PRICE  = "price"
	FIELD_YEAR   = "year"
	SELECT       = "SELECT * FROM " + TABLE
	SELECT_BY_BRAND = "SELECT " + FIELD_BRAND + ", " + FIELD_MODEL +  ", " + FIELD_PRICE+" FROM " + TABLE + " WHERE " + FIELD_BRAND + " = ?"
	INSERT       = "INSERT INTO " + TABLE + " (" + FIELD_ID + ", " + FIELD_BRAND + ") VALUES (?, ?)"
	DELETE       = "DELETE from " + TABLE + " WHERE " + FIELD_ID + " = ?"
	UPDATE       = "UPDATE " + TABLE + " SET " + FIELD_BRAND + " = ? WHERE " + FIELD_ID + " = ? IF EXISTS"
)


func GetByBrand(uuid gocql.UUID, session *gocql.Session,brand string) model.LaptopStruct {
		return getOne(uuid, session,brand)
	}

func DeleteOne(id int, session *gocql.Session) error {
	return deleteOne(session, id)
}


func UpdateOne(id int, session *gocql.Session,laptop string) error {
	 return updateOne( id,laptop,session)
}

func updateOne(uuid int, todo string, session *gocql.Session)error {
	err:= session.Query(UPDATE, todo, uuid).Exec()

	if err != nil {
		log.Printf("ERROR: fail update document, %s", err.Error())
	}

	return err
}

func CreateDocument(session *gocql.Session,laptop *model.LaptopStruct) error {
	q := `
		INSERT INTO eeshan.Laptop (
			id,
		    brand,
		    model,
			price,
			year
		)
		VALUES (?, ?, ?, ?, ?)
    	`

		fmt.Println(q, "query")
		err := session.Query(q,
			laptop.Id,
		laptop.Brand,
		laptop.Model,
		laptop.Price,
		time.Now().Year(),
		).Exec()
	if err != nil {
		log.Printf("ERROR: fail create document, %s", err.Error())
	}

	return err
}


func Get(session *gocql.Session) []model.LaptopStruct{
	var userList []model.LaptopStruct 
	m := map[string]interface{}{}
	query := "SELECT *  FROM eeshan.Laptop"
	iterable :=session.Query(query).Iter()

	for iterable.MapScan(m) {
	  userList = append(userList, model.LaptopStruct{
		Id : m["id"].(int),
		Brand:m["brand"].(string),
		Model : m["model"].(string),
		Year : m["year"].(int),
		Price : m["price"].(int),
	  })
	  m = map[string]interface{}{}
	}

	return userList
  }



func deleteOne(session *gocql.Session, id int) error {
	err:= session.Query(DELETE, id).Exec()
	if err != nil {
		log.Printf("ERROR: fail edit document, %s", err.Error())
	}

	return err
}

func getOne(id gocql.UUID, session *gocql.Session, brand string) model.LaptopStruct {
	var t model.LaptopStruct
	session.Query(SELECT_BY_BRAND, brand).Scan(&t.Id, &t.Brand)
	return t
}


func save(session *gocql.Session, todo *model.LaptopStruct) {
	todo.Id = gocql.MustRandomUUID().Variant()
	if err := session.Query(INSERT, todo.Id, todo.Year).Exec(); err != nil {
		log.Println( err)
	}
}
