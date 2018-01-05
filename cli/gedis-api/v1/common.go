package version1

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucas59356/gedis/cli/gedis-api/requtils"
	"github.com/lucas59356/gedis/core"
)

// Common Represents web ready wrappers for generic object processing
type Common struct {
	DB *core.Thread
}

// Get GET /api/com/{key}
func (o *Common) Get(w http.ResponseWriter, r *http.Request) {
	requtils.RequestLogger(w, r)
	vars := mux.Vars(r)
	k := vars["key"]
	v, tp, err := o.DB.Get(k)
	if requtils.CommonGetErrHandling(w, r, err) {
		return
	}
	requtils.Return(w, r, http.StatusOK, v, core.Types[tp])
}

// Set POST/PUT /api/bool/{key}/{value}
func (o *Common) Set(w http.ResponseWriter, r *http.Request) {
	requtils.RequestLogger(w, r)
	vars := mux.Vars(r)
	k := vars["key"]
	v := vars["value"]
	b := v
	_, tp, err := o.DB.Set(k, b)
	if requtils.CommonSetErrHandling(w, r, err) {
		return
	}
	requtils.Return(w, r, http.StatusOK, k, b, core.Types[tp])
}
