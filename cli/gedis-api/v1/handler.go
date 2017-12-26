package version1

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/lucas59356/gedis/cli/gedis-api/requtils"

	"strconv"

	"github.com/lucas59356/gedis/core"
)

var (
	// TypeString Represents string
	TypeString = core.Types[core.TypeString]
	// TypeInt Represents int
	TypeInt = core.Types[core.TypeInt]
	// TypeBool Represents bool
	TypeBool = core.Types[core.TypeBool]
)

// Handler Api functions
type Handler struct {
	DB *core.Thread
}

// Get GET /api/{type}/{key}
func (o *Handler) Get(w http.ResponseWriter, r *http.Request) {
	requtils.RequestLogger(w, r)
	vars := mux.Vars(r)
	k := vars["key"]
	t := vars["type"]
	v, tp, err := o.DB.Get(k)
	if core.Types[tp] != t {
		requtils.Return(w, r, http.StatusAccepted, v, core.Types[tp], "Found but not of this type", err)
		return
	}
	requtils.Return(w, r, http.StatusOK, v, core.Types[tp], err)
}

// Set POST/PUT /api/{type}/{key}/{value}
func (o *Handler) Set(w http.ResponseWriter, r *http.Request) {
	requtils.RequestLogger(w, r)
	vars := mux.Vars(r)
	k := vars["key"]
	t := vars["type"]
	v := vars["value"]
	pv, err := o.convertData(v, t)
	if err != nil {
		requtils.Return(w, r, http.StatusInternalServerError, nil, t, err)
		return
	}
	rv, tp, err := o.DB.Set(k, pv)
	var rcode int
	if err != nil {
		rcode = http.StatusInternalServerError
	}
	rcode = http.StatusOK
	requtils.Return(w, r, rcode, rv, core.Types[tp], err)
}

// Del DELETE /api/{type}/{key}
func (o *Handler) Del(w http.ResponseWriter, r *http.Request) {
	requtils.RequestLogger(w, r)
	vars := mux.Vars(r)
	k := vars["key"]
	err := o.DB.Del(k)
	if err != nil {
		requtils.Return(w, r, http.StatusInternalServerError, err)
		return
	}
	requtils.Return(w, r, http.StatusOK)
}

func (o *Handler) convertData(v string, tp string) (interface{}, error) {
	switch tp {
	case TypeInt:
		pv, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		return pv, nil
	case TypeBool:
		pv, err := strconv.ParseBool(v)
		if err != nil {
			return nil, err
		}
		return pv, nil
	case TypeString:
		return v, nil
	default:
		iv, err := strconv.Atoi(v)
		if err == nil {
			return iv, nil
		}
		bv, err := strconv.ParseBool(v)
		if err == nil {
			return bv, nil
		}
		return v, nil
	}
}
