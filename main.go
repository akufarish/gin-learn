package main

import (
	"belajar/buku"
	"belajar/siswa"
	"net/http"

	"github.com/gin-gonic/gin"
)


type murid struct {
	Id string `json:"id"`
	Nama string `json:"nama"`
	Kelas string `json:"kelas"`
}
	
var murids = []murid{
	{
		Id: "1",
		Nama: "Farish",
		Kelas: "12",
	},
	{
		Id: "2",
		Nama: "lana",
		Kelas: "12",
	},
	{
		Id: "3",
		Nama: "Yazhid",
		Kelas: "11",
	},
}

func indexMurid (c *gin.Context) {
	c.IndentedJSON(http.StatusOK, murids)
}

func addMurid (c *gin.Context) {
	var newMurid murid

	if err := c.BindJSON(&newMurid); err != nil {
		return
	}

	murids = append(murids, newMurid)
	c.IndentedJSON(http.StatusCreated, newMurid)
}

func main() {
	r := gin.Default()

	r.GET("/murid", indexMurid)
	r.POST("/murid", addMurid)

	r.GET("/siswa", siswa.IndexSiswa)
	r.POST("/siswa", siswa.AddSiswa)

	r.GET("/buku", buku.IndexBuku)
	r.POST("/buku", buku.StoreBuku)
	r.Run() // listen and serve on 0.0.0.0:8080
}
