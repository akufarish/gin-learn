package main

import (
	"belajar/buku"
	"belajar/makanan"
	"belajar/siswa"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/siswa", siswa.IndexSiswa)
	r.POST("/siswa", siswa.AddSiswa)
	r.GET("/siswa/:id", siswa.SingleSiswa)

	r.GET("/buku", buku.IndexBuku)
	r.POST("/buku", buku.StoreBuku)
	r.GET("/buku/:id", buku.SingleBuku)

	r.GET("/menu", makanan.IndexMenu)
	r.POST("/menu", makanan.StoreMenu)

	r.Run() // listen and serve on 0.0.0.0:8080
}
