package buku

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type buku struct {
	Id string `json:"id"`
	Judul string `json:"judul"`
	Penulis string `json:"penulis"`
	Date string `json:"date"`
}

var bukus = []buku{
	{
		Id: "1",
		Judul: "Belajar Go",
		Penulis: "Farish",
		Date: "13-06-2006",
	},
	{
		Id: "2",
		Judul: "Belajar Typescript",
		Penulis: "Farish",
		Date: "6-12-2020",
	},
}

func IndexBuku (c *gin.Context) {
	c.IndentedJSON(http.StatusOK, bukus)
}

func StoreBuku (c *gin.Context) {
	var newBuku buku

	if err := c.BindJSON(&newBuku); err != nil {
		return
	}

	bukus = append(bukus, newBuku)
	c.IndentedJSON(http.StatusCreated, newBuku)
}

func SingleBuku (c *gin.Context) {
	id := c.Param("id")

	for _, dataBuku := range bukus {
		if dataBuku.Id == id {
			c.IndentedJSON(http.StatusOK, dataBuku)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, "Data tidak ditemukan")
}