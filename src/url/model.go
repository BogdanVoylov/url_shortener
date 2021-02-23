package url

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type UrlModel struct {
	db *sqlx.DB
}

func NewUrlModel(db *sqlx.DB) *UrlModel {
	return &UrlModel{db}
}

func (this *UrlModel) AddUrl(url string) int {
	this.db.MustExec(`INSERT INTO urls (url) VALUES ($1)`,url)
	var res int
	this.db.Get(&res, `SELECT currval('urls_id_seq')`)
	log.Printf("New url %s",url)
	return res
}

func (this *UrlModel) GetUrl(id int) string {
	log.Println(id)
	var res string
	this.db.Get(&res, "SELECT url FROM urls WHERE id=$1",id)
	return res
}
