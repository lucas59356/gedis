package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucas59356/gedis/api/requtils"
	"github.com/lucas59356/gedis/api/v1"
	"github.com/lucas59356/gedis/core"
)

const bind = ":80"

// DB Gedis database
var DB = core.NewThread()

var (
	// Router Http base router
	Router = mux.NewRouter().PathPrefix("/api").Subrouter()
)

func main() {
	log.Println("Iniciando...")
	Router.HandleFunc("/", requtils.RequestLogger).Methods(http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete)
	// 		REGISTRO DE ROTAS
	objDeleter := &version1.Deleter{DB}
	log.Println("Registrando as rotas...")
	v1 := Router.NewRoute().PathPrefix("/v1/").Subrouter()
	// 					BOOL
	objBool := &version1.Bool{DB}
	v1.HandleFunc("/bool/{key}", objBool.Get).Methods(http.MethodGet)                          // GET
	v1.HandleFunc("/bool/{key}/{value}", objBool.Set).Methods(http.MethodPost, http.MethodPut) // SET
	v1.HandleFunc("/bool/{key}", objDeleter.Del).Methods(http.MethodDelete)                    // DEL
	//					NUMBER
	objNumber := &version1.Number{DB}
	v1.HandleFunc("/number/{key}", objNumber.Get).Methods(http.MethodGet)                          // GET
	v1.HandleFunc("/number/{key}/{value}", objNumber.Set).Methods(http.MethodPost, http.MethodPut) // SET
	v1.HandleFunc("/number/{key}", objDeleter.Del).Methods(http.MethodDelete)                      // DEL
	//					STRING
	objString := &version1.String{DB}
	v1.HandleFunc("/string/{key}", objString.Get).Methods(http.MethodGet)                          // GET
	v1.HandleFunc("/string/{key}/{value}", objString.Set).Methods(http.MethodPost, http.MethodPut) // SET
	v1.HandleFunc("/string/{key}", objDeleter.Del).Methods(http.MethodDelete)                      // DEL
	//					COMMON
	objCommon := &version1.Common{DB}
	v1.HandleFunc("/com/{key}", objCommon.Get).Methods(http.MethodGet)                          // GET
	v1.HandleFunc("/com/{key}/{value}", objCommon.Set).Methods(http.MethodPost, http.MethodPut) // SET
	v1.HandleFunc("/com/{key}", objDeleter.Del).Methods(http.MethodDelete)                      // DEL
	log.Printf("Escutando em %s...", bind)
	err := http.ListenAndServe(bind, Router)
	panic(err)
}
