package main

import (
	"log"

	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func main() {
	log.SetFlags(0)

	rdbOpts := r.ConnectOpts{
		Address: "localhost:28015",
	}

	rconn, err := r.Connect(rdbOpts)
	checkError(err)

	err = r.DB("test").TableCreate("tv_shows").Exec(rconn)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}
