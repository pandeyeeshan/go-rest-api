package bootstrap

import (
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"
)

// Session holds our connection to Cassandra
var Session *gocql.Session

func init() {

	cluster := gocql.NewCluster("127.0.0.1:9042")
	cluster.Keyspace = "eeshan"
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: "test", Password: "testpwd"}
	cluster.Port=9042
	cluster.Consistency = gocql.One

	cluster.RetryPolicy = &gocql.SimpleRetryPolicy{
		NumRetries: 3,
	}
	cluster.Consistency = gocql.One
	var session *gocql.Session
	var err error
	for {
		session, err = cluster.CreateSession()
		if err == nil {
			fmt.Println("session", session, "err", err)
			break
		}
		log.Printf("CreateSession: %v", err)
		time.Sleep(time.Second)
	}
	fmt.Println("session", session==nil, "err", err)
	log.Printf("Connected OK\n")
	Session=session
	// defer session.Close()
}