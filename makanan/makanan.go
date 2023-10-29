package makanan

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type menu struct {
	Id string `json:"id"`
	Tipe string `json:"tipe"`
	Nama string `json:"nama"`
	Harga string `json:"harga"`
}

var menus = []menu{
	{
		Id: "1",
		Tipe: "makanan",
		Nama: "Nasi goreng",
		Harga: "Rp.10.000",
	},
	{
		Id: "2",
		Tipe: "minuman",
		Nama: "Teh es",
		Harga: "Rp.3.000",
	},
	{
		Id: "3",
		Tipe: "makanan",
		Nama: "Bakso",
		Harga: "Rp.15.000",
	},
}

func IndexMenu (c *gin.Context) {
	c.IndentedJSON(http.StatusOK, menus)
}

func StoreMenu (c *gin.Context) {
	var newMenu menu

	if err := c.BindJSON(&newMenu); err != nil {
		return
	}

	menus = append(menus, newMenu)
	c.IndentedJSON(http.StatusCreated, newMenu)
}