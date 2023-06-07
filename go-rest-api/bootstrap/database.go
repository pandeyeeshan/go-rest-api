package bootstrap

import (
	"be-training/go-rest-api/pkg/model"
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"
)

//Open close cassandra connections


type Signature struct {
	Signer    string `cql:"signer"`
	Comment   string `cql:"comment"`
	Timestamp int64  `cql:"timestamp"`
}

type Laptop struct {
	Id        gocql.UUID   `cql:"id"`
	Brand	  string       `cql:"brand"`
	Model 	  string       `cql:"model"`
	Year 	  int          `cql:"year"`
	Price 	  int          `cql:"price"`
}

// var Session *gocql.Session

// func init() {
// 	var err error
// 	cluster := gocql.NewCluster("127.0.0.1:9042")
// 	cluster.Keyspace = "eeshan"
// 	cluster.Port=9042
// 	cluster.ProtoVersion = 4
//     cluster.Consistency = gocql.Quorum
//     cluster.CQLVersion = "3.4.5"
//     cluster.IgnorePeerAddr = true
//     cluster.DefaultIdempotence = true
//     cluster.Timeout = time.Second * 3000000
//     cluster.ConnectTimeout = time.Second * 30000000
// 	cluster.Authenticator = gocql.PasswordAuthenticator{Username: "test", Password: "testpwd"}

// 	Session, err = cluster.CreateSession()
// 	if err != nil {
// 	  panic(err)
// 	}
// 	fmt.Println("cassandra init done")
//   }

  

func NewCassandraDatabase(env *Env) {


// 	// cassandraConfigHost := env.CassandraHost
// 	cassandraConfigPort := env.CassandraPort
// 	cassandraConfigKeyspace := env.CassandraKeyspace
// 	cassandraUser:=env.CASSANDRA_USER
// 	cassandraPass:=env.CASSANDRA_PASSWORD
// 	// cassandraConfigConsistancy := env.CassanraConsistency

// 	cluster := gocql.NewCluster("127.0.0.1:9042") 
//     cluster.Keyspace = cassandraConfigKeyspace
//     cluster.Port = cassandraConfigPort

//     cluster.DisableInitialHostLookup = true
//     cluster.Authenticator = gocql.PasswordAuthenticator{Username: cassandraUser, Password: cassandraPass}
//     cluster.ProtoVersion = 4
//     cluster.Consistency = gocql.Quorum
//     cluster.CQLVersion = "3.4.5"
//     cluster.IgnorePeerAddr = true
//     cluster.DefaultIdempotence = true
//     cluster.Timeout = time.Second * 3000000
//     cluster.ConnectTimeout = time.Second * 30000000
//     session, err := cluster.CreateSession()
//     if err != nil {
//         fmt.Println("error in session", err)
//         return
//     }
// 	if session == nil || session.Closed() {
//         session, err := cluster.CreateSession()
// 		fmt.Println("session", session, "err", err)
//     }
//     // defer session.Close()
// }

// func CloseCassandraConnection() {
// 	Session.Close() 


cluster := gocql.NewCluster("127.0.0.1:9042")
	cluster.Keyspace = "eeshan"
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: "test", Password: "testpwd"}
	cluster.Port=9042
	cluster.RetryPolicy = &gocql.SimpleRetryPolicy{
		NumRetries: 3,
	}
	cluster.Consistency = gocql.Quorum
	var session *gocql.Session
	var err error
	for {
		session, err = cluster.CreateSession()
		if err == nil {
			break
		}
		log.Printf("CreateSession: %v", err)
		time.Sleep(time.Second)
	}
	log.Printf("Connected OK\n")
	defer session.Close()
	
}


func CreateDocument(laptop *model.LaptopStruct) error {

	fmt.Println(laptop.Brand, "laptop")
	q := `
		INSERT INTO Laptop (
		    id,
		    brand,
		    model
		    year,
			price
		)
		VALUES (?, ?, ?, ?, ?)
    	`

		fmt.Println(q, "query")
	err := Session.Query(q,
		laptop.Id,
		laptop.Brand,
		laptop.Model,
		laptop.Year,
		laptop.Price).Exec()
		fmt.Println(err, "error#########################################")
	if err != nil {
		log.Printf("ERROR: fail create document, %s", err.Error())
	}

	return err
}


func Get() []model.LaptopStruct{
	var userList []model.LaptopStruct 
	m := map[string]interface{}{}
	query := "SELECT *  FROM eeshan.Laptop"
	iterable :=Session.Query(query).Iter()

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


  func GetDocumentById() model.LaptopStruct{
	var userFromDB model.LaptopStruct
	fmt.Println("inside get document by id")
	fmt.Print(Session, "session")
	if err := Session.Query("SELECT * from eeshan.Laptop LIMIT 1").WithContext(context.Background()).Scan(&userFromDB.Id, &userFromDB.Brand, &userFromDB.Model, &userFromDB.Price,&userFromDB.Year); 
	err != nil {
		fmt.Println("select error",err)
		log.Fatal(err)
	} 
	fmt.Println(userFromDB)
	return userFromDB

}

func GetDocument(company string, id gocql.UUID) (*Laptop, error) {
	// toSignatures := func(i interface{}) []Signature {
	// 	sigs := []Signature{}
	// 	sig := Signature{}
	// 	for _, s := range i.([]map[string]interface{}) {
	// 		mapstructure.Decode(s, &sig)
	// 		sigs = append(sigs, sig)
	// 	}

	// 	return sigs
	// }

	m := map[string]interface{}{}
	q := `
		SELECT * FROM documents
			WHERE company = ? AND id = ?
		LIMIT 1
    	`
	itr := Session.Query(q, company, id).Consistency(gocql.One).Iter()
	for itr.MapScan(m) {
		document := &Laptop{}
		document.Id = m["id"].(gocql.UUID)
		document.Brand = m["brand"].(string)
		document.Model = m["model"].(string)
		document.Year = m["year"].(int)
		document.Price = m["price"].(int)

		log.Printf("INFO: found document, %v", document)

		return document, nil
	}

	return nil, errors.New("document not found")
}

func UpdateDocument(company string, id gocql.UUID, name string, status string) error {
	q := `
        	UPDATE laptop
		SET name = ?, status = ?
		WHERE company = ? AND id = ?
    	`
	err := Session.Query(q, name, status, company, id).Exec()
	if err != nil {
		log.Printf("ERROR: fail update document, %s", err.Error())
		return err
	}

	return nil
}

func AddTag(company string, id gocql.UUID, tag string) error {
	q := `
		UPDATE mystiko.documents
		SET tags = tags + ?
		WHERE company = ? AND id = ?;
	`

	err := Session.Query(q, []string{tag}, company, id).Exec()
	if err != nil {
		log.Printf("ERROR: fail add tag, %s", err.Error())
		return err
	}

	return nil
}

func RemoveTag(company string, id gocql.UUID, tag string) error {
	q := `
		UPDATE mystiko.documents
		SET tags = tags - ?
		WHERE company = ? AND id = ?;
	`

	err := Session.Query(q, []string{tag}, company, id).Exec()
	if err != nil {
		log.Printf("ERROR: fail remove tag, %s", err.Error())
		return err
	}

	return nil
}

func AddSignature(company string, id gocql.UUID, signature Signature) error {
	q := `
		UPDATE mystiko.documents
		SET signatures = signatures + ?
		WHERE company = ? AND id = ?;
	`

	err := Session.Query(q, []Signature{signature}, company, id).Exec()
	if err != nil {
		log.Printf("ERROR: fail add signature, %s", err.Error())
		return err
	}

	return nil
}

func removeSignature(company string, id gocql.UUID, signature Signature) error {
	q := `
		UPDATE mystiko.documents
		SET signatures = signatures - ?
		WHERE company = ? AND id = ?;
	`

	err := Session.Query(q, []Signature{signature}, company, id).Exec()
	if err != nil {
		log.Printf("ERROR: fail remove signature, %s", err.Error())
		return err
	}

	return nil
}