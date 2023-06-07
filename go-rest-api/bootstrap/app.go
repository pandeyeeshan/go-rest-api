package bootstrap

import "github.com/gocql/gocql"

//Application struct
type Application struct {
	Env   *Env
	Cassandra *gocql.Session
}


//App Env and Cassandra Initiatlization here 
func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	return *app
}

// func (app *Application) CloseDBConnection() {
// 	CloseCassandraConnection()
// }
