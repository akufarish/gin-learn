package siswa

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type siswa struct {
	Id string `json:"id"`
	Nama string `json:"nama"`
	Kelas string `json:"kelas"`
	Jurusan string `json:"jurusan"`
}

var siswas = []siswa{
	{
		Id: "1",
		Nama: "Farish",
		Kelas: "12",
		Jurusan: "PPLG",
	},
	{
		Id: "2",
		Nama: "lana",
		Kelas: "12",
		Jurusan: "PPLG",
	},
	{
		Id: "3",
		Nama: "Rifki",
		Kelas: "12",
		Jurusan: "TKJ",
	},
}

func IndexSiswa (c *gin.Context) {
	c.IndentedJSON(http.StatusOK, siswas)
}

func AddSiswa (c *gin.Context) {
	var newSiswa siswa

	if err := c.BindJSON(&newSiswa); err != nil {
		return	
	}

	siswas = append(siswas, newSiswa)
	c.IndentedJSON(http.StatusCreated, newSiswa)
}

