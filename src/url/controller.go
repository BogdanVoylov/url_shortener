package url

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"
	"url_shortener/src/httpUtils"
)


type HttpApp interface {
	HandleFunc(url string, h func (http.ResponseWriter, *http.Request)) *mux.Route
	Db() *sqlx.DB
}

type UrlController struct {
	model UrlModel
}

func NewUrlController(app HttpApp) *UrlController {
	m := NewUrlModel(app.Db())
	c := &UrlController{*m}
	app.HandleFunc("/url/add",c.AddUrl).Methods("POST")
	app.HandleFunc("/url/get/{id}",c.GetUrl).Methods("GET")
	return c
}

func (this *UrlController) AddUrl(w http.ResponseWriter, r *http.Request) {
	url := new(string)

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&url); err != nil {
		httpUtils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	id := this.model.AddUrl(*url)
	httpUtils.RespondWithJSON(w,200,id)
}

func (this *UrlController) GetUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		httpUtils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	url := this.model.GetUrl(id)
	httpUtils.RespondWithJSON(w,200,url)
}