package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/Jonathansoufer/go-bookstore/pkg/routes"
)

func main(){
	r := mux.NewRouter()
}