package version1

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lucas59356/gedis/cli/gedis-api/requtils"
	"github.com/lucas59356/gedis/core"
)

// Bool Represents web ready wrappers for bool processing
type Bool struct {
	DB *core.Thread
}

// Get GET /api/bool/{key}
func (o *Bool) Get(w http.ResponseWriter, r *http.Request) {
	requtils.RequestLogger(w, r)
	vars := mux.Vars(r)
	k := vars["key"]
	v, tp, err := o.DB.Get(k)
	if requtils.CommonGetErrHandling(w, r, err) {
		return
	}
	if requtils.CommonGetIfNotType(w, r, tp, core.TypeBool) {
		return
	}
	requtils.Return(w, r, http.StatusOK, v, core.Types[tp])
}

// Set POST/PUT /api/bool/{key}/{value}
func (o *Bool) Set(w http.ResponseWriter, r *http.Request) {
	requtils.RequestLogger(w, r)
	vars := mux.Vars(r)
	k := vars["key"]
	v := vars["value"]
	b, err := strconv.ParseBool(v)
	if requtils.CommonSetBadRequestIfNotParseable(w, r, err) {
		return
	}
	_, tp, err := o.DB.Set(k, b)
	if requtils.CommonSetIfNotType(w, r, tp, core.TypeBool) {
		return
	}
	if requtils.CommonSetErrHandling(w, r, err) {
		return
	}
	requtils.Return(w, r, http.StatusOK, k, b, core.Types[tp])
}
