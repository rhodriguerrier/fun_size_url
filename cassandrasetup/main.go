package cassandrasetup

import (
	"log"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

func init() {
	var err error
	cluster := gocql.NewCluster("cassandra")
	Session, err = cluster.CreateSession()
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Cassandra session successfully started")
	}
}
