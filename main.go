package main

import (
	"encoding/json"
	"fmt"
	"log"

	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

type show struct {
	Name     string    `json:"name" gorethink:"name"`
	Genre    string    `json:"genre" gorethink:"genre"`
	Website  string    `json:"website" gorethink:"website"`
	Episodes []episode `json:"episodes" gorethink:"episodes"`
}

type episode struct {
	Name    string `json:"name" gorethink:"name"`
	Summary string `json:"summary" gorethink:"summary"`
}

func main() {
	log.SetFlags(0)

	rdbOpts := r.ConnectOpts{
		Address: "localhost:28015",
	}

	rconn, err := r.Connect(rdbOpts)
	checkError(err)

	for {
		result, err := r.Table("tv_shows").Changes().Run(rconn)
		checkError(err)

		var response interface{}
		for result.Next(&response) {
			printObj(response)
		}
		checkError(result.Err())
	}

}

func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func printObj(v interface{}) {
	vBytes, _ := json.Marshal(v)
	fmt.Println(string(vBytes))
}
