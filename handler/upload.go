package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"example.com/morethanjustlinks/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) Upload(ctx *gin.Context) {

	form, _ := ctx.MultipartForm()
	files := form.File["file"]

	id := ctx.PostForm("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing required id"})
		return
	}

	uuid, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong parsing uuid"})
		return
	}

	// Check directory exists
	// TODO store files in GCP bucket
	if err := os.MkdirAll("./nextjs-frontend/public/users/"+id+"/", 0755); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went preparing to place uploaded file"})
		return
	}

	for _, file := range files {
		// TODO store files in GCP bucket
		dst, err := filepath.Abs("./nextjs-frontend/public/users/" + id + "/profile_img.png")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong uploading..."})
			return
		}

		if err := ctx.SaveUploadedFile(file, dst); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "somethin went wrong saving file..."})
			return
		}

		// save path to db
		h.db.Model(&models.User{}).Where("id = ?", uuid).Update("profile_pic", "/users/"+id+"/profile_img.png")
	}

	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%d file uploaded!", len(files))})

}
