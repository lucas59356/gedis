package version1

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucas59356/gedis/cli/gedis-api/requtils"
	"github.com/lucas59356/gedis/core"
)

// String Represents web ready wrappers for string processing
type String struct {
	DB *core.Thread
}

// Get GET /api/string/{key}
func (o *String) Get(w http.ResponseWriter, r *http.Request) {
	requtils.RequestLogger(w, r)
	vars := mux.Vars(r)
	k := vars["key"]
	v, tp, err := o.DB.Get(k)
	if requtils.CommonGetErrHandling(w, r, err) {
		return
	}
	if requtils.CommonGetIfNotType(w, r, tp, core.TypeString) {
		return
	}
	requtils.Return(w, r, http.StatusOK, v, core.Types[tp])
}

// Set POST/PUT /api/string/{key}/{value}
func (o *String) Set(w http.ResponseWriter, r *http.Request) {
	requtils.RequestLogger(w, r)
	vars := mux.Vars(r)
	k := vars["key"]
	v := vars["value"]
	b := v
	_, tp, err := o.DB.Set(k, b)
	if requtils.CommonSetIfNotType(w, r, tp, core.TypeString) {
		return
	}
	if requtils.CommonSetErrHandling(w, r, err) {
		return
	}
	requtils.Return(w, r, http.StatusOK, k, b, core.Types[tp])
}
