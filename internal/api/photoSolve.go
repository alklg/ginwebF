package api

import (
	"fmt"
	"ginweb/internal/db"
	"ginweb/internal/models"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func GetPhotoSolve(c *gin.Context) {
	fmt.Println("get photo solve success")

	currentDir, err := os.Getwd()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": "open directory error"})
	}
	downloadPath := filepath.Join(currentDir, "././downloads") + "/output_image.png"
	fmt.Println(downloadPath)

	data, err := ioutil.ReadFile(downloadPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to read processed image"})
	}

	c.Header("Content-Type", "image/jpeg")
	c.Header("Content-Disposition", "attachment; filename=processed_image.jpg")

	c.Data(http.StatusOK, "image/jpeg", data)
	return
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

	var filename string
	for _, file := range files {
		src, err := file.Open()

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"data": "Can't open file",
			})

			return
		}

		defer src.Close()
		filename = file.Filename

		fmt.Println(filename)

		dst, err := os.Create(uploadPath + filename)
		fmt.Println("used for test----", uploadPath+filename)

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

	pythonPath := filepath.Join(currentDir, "././") + "/main.py"
	println(pythonPath)

	cmd := exec.Command("python3", pythonPath, uploadPath+filename)
	fmt.Println(cmd)

	output, err := cmd.Output()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": "Error executing Python script",
		})
		return
	}

	fmt.Println(output)
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
