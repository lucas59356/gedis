package version1

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucas59356/gedis/api/requtils"
	"github.com/lucas59356/gedis/core"
)

// Deleter Avoid code repetition on other wrappers because Delete method is the same
type Deleter struct {
	DB *core.Thread
}

// Del DELETE /api/[datatype]]/{key}
func (o *Deleter) Del(w http.ResponseWriter, r *http.Request) {
	requtils.RequestLogger(w, r)
	vars := mux.Vars(r)
	k := vars["key"]
	err := o.DB.Del(k)
	if requtils.CommonSetErrHandling(w, r, err) {
		return
	}
	requtils.Return(w, r, http.StatusOK, "Sucess")
}
