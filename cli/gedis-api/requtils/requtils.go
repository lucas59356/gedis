package requtils

import (
	"encoding/json"
	"log"
	"net/http"
)

// RequestLogger Register all requests on stdout
func RequestLogger(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s (%s)\n", r.Method, r.RequestURI, r.RemoteAddr)
}

// Return process and return data provided by http handlers
func Return(w http.ResponseWriter, r *http.Request, httpCode int, datas ...interface{}) error {
	errs := []string{}
	values := []interface{}{}
	for _, data := range datas {
		switch proc := data.(type) {
		case error:
			errs = append(errs, proc.Error())
		default:
			values = append(values, proc)
		}
	}
	v2r := map[string]interface{}{}
	if len(errs) != 0 {
		v2r["errors"] = errs
	}
	if len(values) != 0 {
		v2r["data"] = values
	}
	d, err := json.Marshal(v2r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	w.WriteHeader(httpCode)
	_, err = w.Write(d)
	return err
}

// CommonGetErrHandling Evitar repetição no processo de administrar os erros vindos do DB
func CommonGetErrHandling(w http.ResponseWriter, r *http.Request, err error) bool {
	if err != nil {
		Return(w, r, http.StatusNotFound, err)
		return true
	}
	return false
}

// CommonGetIfNotType Evitar repetição no processo de ver se o conteúdo é compatível
func CommonGetIfNotType(w http.ResponseWriter, r *http.Request, dataType int8, neededType int8) bool {
	if dataType != neededType {
		Return(w, r, http.StatusNoContent)
		return true
	}
	return false
}

func CommonSetBadRequestIfNotParseable(w http.ResponseWriter, r *http.Request, err error) bool {
	if err != nil {
		Return(w, r, http.StatusBadRequest, err)
		return true
	}
	return false
}

func CommonSetIfNotType(w http.ResponseWriter, r *http.Request, dataType int8, neededType int8) bool {
	if dataType != neededType {
		Return(w, r, http.StatusInternalServerError)
		return true
	}
	return false
}

func CommonSetErrHandling(w http.ResponseWriter, r *http.Request, err error) bool {
	if err != nil {
		Return(w, r, http.StatusInternalServerError, err)
		return true
	}
	return false
}
