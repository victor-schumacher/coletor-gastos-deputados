package main

import (
	"coletor-gastos-deputados/cron"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)



func main() {


	cron.Start()
	r := mux.NewRouter()
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(port, r))
}

