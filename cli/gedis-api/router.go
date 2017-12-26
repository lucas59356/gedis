package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucas59356/gedis/cli/gedis-api/requtils"
	"github.com/lucas59356/gedis/cli/gedis-api/v1"
	"github.com/lucas59356/gedis/cli/gedis-api/v2"
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
	log.Println("Registrando as rotas...")
	v1 := Router.NewRoute().PathPrefix("/v1/").Subrouter()
	regv1(v1)

	log.Printf("Escutando em %s...", bind)
	v2 := Router.NewRoute().PathPrefix("/v2/").Subrouter()
	regv2(v2)
	err := http.ListenAndServe(bind, Router)
	panic(err)
}

func regv1(r *mux.Router) {
	objDeleter := &version1.Deleter{DB}
	// 					BOOL
	objBool := &version1.Bool{DB}
	r.HandleFunc("/bool/{key}", objBool.Get).Methods(http.MethodGet)                          // GET
	r.HandleFunc("/bool/{key}/{value}", objBool.Set).Methods(http.MethodPost, http.MethodPut) // SET
	r.HandleFunc("/bool/{key}", objDeleter.Del).Methods(http.MethodDelete)                    // DEL
	//					NUMBER
	objNumber := &version1.Number{DB}
	r.HandleFunc("/number/{key}", objNumber.Get).Methods(http.MethodGet)                          // GET
	r.HandleFunc("/number/{key}/{value}", objNumber.Set).Methods(http.MethodPost, http.MethodPut) // SET
	r.HandleFunc("/number/{key}", objDeleter.Del).Methods(http.MethodDelete)                      // DEL
	//					STRING
	objString := &version1.String{DB}
	r.HandleFunc("/string/{key}", objString.Get).Methods(http.MethodGet)                          // GET
	r.HandleFunc("/string/{key}/{value}", objString.Set).Methods(http.MethodPost, http.MethodPut) // SET
	r.HandleFunc("/string/{key}", objDeleter.Del).Methods(http.MethodDelete)                      // DEL
	//					COMMON
	objCommon := &version1.Common{DB}
	r.HandleFunc("/com/{key}", objCommon.Get).Methods(http.MethodGet)                          // GET
	r.HandleFunc("/com/{key}/{value}", objCommon.Set).Methods(http.MethodPost, http.MethodPut) // SET
	r.HandleFunc("/com/{key}", objDeleter.Del).Methods(http.MethodDelete)                      // DEL
}

func regv2(r *mux.Router) {
	obj := &version2.Handler{DB}
	r.HandleFunc("/{type}/{key}", obj.Get).Methods(http.MethodGet)                          // GET
	r.HandleFunc("/{type}/{key}/{value}", obj.Set).Methods(http.MethodPost, http.MethodPut) // SET
	r.HandleFunc("/{type}/{key}", obj.Del).Methods(http.MethodDelete)                       // DEL
}
