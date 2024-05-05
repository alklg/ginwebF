package api

import (
	"fmt"
	"ginweb/internal/db"
	"ginweb/internal/models"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func GetPhotoSolve(c *gin.Context) {
	fmt.Println("get photo solve success")
}

func PostPhotoSolve(c *gin.Context) {

	form, err := c.MultipartForm()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": "Send photo error",
		})

		return
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return
	}

	uploadPath := filepath.Join(currentDir, "././uploads") + "/"

	fmt.Println(uploadPath)

	files := form.File["file"]
	fmt.Println("this is", files)

	if len(files) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"data": "Can't find file",
		})

		return
	}

	for _, file := range files {
		src, err := file.Open()

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"data": "Can't open file",
			})

			return
		}

		defer src.Close()
		fmt.Println(file.Filename)

		dst, err := os.Create(uploadPath + file.Filename)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"data": "Can't create target file",
			})

			return
		}
		defer dst.Close()

		_, err = io.Copy(dst, src)
		fmt.Println(dst.Name())
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"data": "Can't save target file",
			})

			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "upload file success",
	})

	fmt.Println("123")
	return
}

func CreateDirectory(user *models.User) {
	uid := db.ParseUserMessage(user)

	err := os.MkdirAll(fmt.Sprintf("../../uploads/%d", uid), 0755)

	if err != nil {
		fmt.Println("Error in creating directory:", err)
		return
	}
}

func UploadImage() {
	return
}
