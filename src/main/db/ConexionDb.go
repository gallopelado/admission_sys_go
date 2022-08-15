package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
)

var (
	host     = ""
	port     = ""
	user     = ""
	password = ""
	dbname   = ""
)

func getProperties() {
	pwd, _ := os.Getwd()
	yfile, err := ioutil.ReadFile(pwd + "/src/resources/application.properties.yaml")
	if err != nil {
		log.Fatal(err)
	}
	data := make(map[interface{}]interface{})
	err_unmarshal := yaml.Unmarshal(yfile, &data)
	if err_unmarshal != nil {
		log.Fatal(err_unmarshal)
	}
	host = fmt.Sprint(data["host"])
	port = fmt.Sprint(data["port"])
	user = fmt.Sprint(data["user"])
	password = fmt.Sprint(data["password"])
	dbname = fmt.Sprint(data["dbname"])
}

func createConnectionString() string {
	getProperties()
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	return psqlInfo
}

func OpenConnectionDatabase() *sql.DB {
	db, err := sql.Open("postgres", createConnectionString())
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conexion Exitosa")
	return db
}
