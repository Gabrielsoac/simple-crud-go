package main

import (
	"log"
	"net/http"

	"github.com/gabrielsoac/simple-go-crud/config"
)

func main() {
	DbConnection := config.SetupDataBase();
	defer DbConnection.Close();
	log.Fatal(http.ListenAndServe(":8000", nil));
}

