package version1

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucas59356/gedis/cli/gedis-api/requtils"

	"strconv"

	"github.com/lucas59356/gedis/core"
)

// Number Represents web ready wrappers for integer number processing
type Number struct {
	DB *core.Thread
}

// Get GET /api/number/{key}
func (o *Number) Get(w http.ResponseWriter, r *http.Request) {
	requtils.RequestLogger(w, r)
	vars := mux.Vars(r)
	k := vars["key"]
	v, tp, err := o.DB.Get(k)
	if requtils.CommonGetErrHandling(w, r, err) {
		return
	}
	if requtils.CommonGetIfNotType(w, r, tp, core.TypeInt) {
		return
	}
	requtils.Return(w, r, http.StatusOK, v, core.Types[tp])
}

// Set POST/PUT /api/number/{key}/{value}
func (o *Number) Set(w http.ResponseWriter, r *http.Request) {
	requtils.RequestLogger(w, r)
	vars := mux.Vars(r)
	k := vars["key"]
	v := vars["value"]
	b, err := strconv.Atoi(v)
	if requtils.CommonSetBadRequestIfNotParseable(w, r, err) {
		return
	}
	_, tp, err := o.DB.Set(k, b)
	if requtils.CommonSetIfNotType(w, r, tp, core.TypeInt) {
		return
	}
	if requtils.CommonSetErrHandling(w, r, err) {
		return
	}
	requtils.Return(w, r, http.StatusOK, k, b, core.Types[tp])
}
