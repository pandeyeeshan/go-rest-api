// package store

// import (
// 	"be-training/go-rest-api/pkg/model"

// 	"github.com/gocql/gocql"
// )

// const (
// 	TABLE        = "eeshan.Laptoo"
// 	FIELD_ID     = "id"
// 	FIELD_BRAND  = "brand"
// 	FIELD_MODEL  = "model"
// 	FIELD_PRICE  = "price"
// 	FIELD_YEAR   = "year"
// 	FIELD_TEXT   = "text"
// 	SELECT       = "SELECT * FROM " + TABLE
// 	SELECT_BY_BRAND = "SELECT " + FIELD_BRAND + ", " + FIELD_MODEL +  ", " + FIELD_PRICE+" FROM " + TABLE + " WHERE " + FIELD_BRAND + " = ?"
// 	INSERT       = "INSERT INTO " + TABLE + " (" + FIELD_ID + ", " + FIELD_TEXT + ") VALUES (?, ?)"
// 	DELETE       = "DELETE from " + TABLE + " WHERE " + FIELD_ID + " = ?"
// 	UPDATE       = "UPDATE " + TABLE + " SET " + FIELD_TEXT + " = ? WHERE " + FIELD_ID + " = ? IF EXISTS"
// )

// func GetById(uuid gocql.UUID, session *gocql.Session) model.LaptopStruct {
// 	return getOne(uuid, session)
// }


// func GetByBrand(uuid gocql.UUID, session *gocql.Session,brand string) model.LaptopStruct {
// 		return getOne(uuid, session,brand)
// 	}

// func GetTodo(session *gocql.Session, state string) ([]Todo, string) {
// 	return findAll(session, state)
// }

// func DeleteOne(id gocql.UUID, session *gocql.Session) {
// 	deleteOne(session, id)
// }

// func PostTodo(t *Todo, session *gocql.Session) {
// 	save(session, t)
// }

// func UpdateOne(uuid gocql.UUID, todo *Todo, session *gocql.Session) {
// 	update(uuid, todo, session)
// }

// func deleteOne(session *gocql.Session, id gocql.UUID) {
// 	session.Query(DELETE, id).Exec()
// }

// func getOne(id gocql.UUID, session *gocql.Session, brand string) model.LaptopStruct {
// 	var t model.LaptopStruct
// 	session.Query(SELECT_BY_BRAND, brand).Scan(&t.ID, &t.Name)
// 	return t
// }

// func findAll(session *gocql.Session, state string) ([]Todo, string) {
// 	var ts []Todo
// 	var t Todo
// 	query := session.Query(SELECT)
// 	if state != "" {
// 		st, _ := b64.StdEncoding.DecodeString(state)
// 		query.PageState(st)
// 	}
// 	it := query.Iter()
// 	sw := it.WillSwitchPage()
// 	for !sw && it.Scan(&t.ID, &t.Name) {
// 		ts = append(ts, t)
// 		sw = it.WillSwitchPage()
// 	}
// 	if err := it.Close(); err != nil {
// 		log.Println(LOG_ERROR_FIND_ALL, err)
// 	}
// 	return ts, b64.StdEncoding.EncodeToString(it.PageState())
// }

// func save(session *gocql.Session, todo *Todo) {
// 	todo.ID = gocql.TimeUUID()
// 	if err := session.Query(INSERT, todo.ID, todo.Name).Exec(); err != nil {
// 		log.Println(LOG_ERROR, err)
// 	}
// }

// func update(uuid gocql.UUID, todo *Todo, session *gocql.Session) {
// 	session.Query(UPDATE, todo.Name, uuid).Exec()
// }