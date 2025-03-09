package handlers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shevchenko-nikita/EDiary/internals/models"
	"github.com/shevchenko-nikita/EDiary/internals/services"
	"mime/multipart"
	"path/filepath"
	"strconv"
)

type Handler struct {
	Database *sql.DB
}

func NewHandler(database *sql.DB) *Handler {
	return &Handler{Database: database}
}

func GetUserFromCookie(c *gin.Context) (*models.User, bool) {
	userCookie, ok := c.Get("user")

	if !ok {
		return nil, false
	}

	user, ok := userCookie.(*models.User)

	if !ok {
		return nil, false
	}

	return user, true
}

func SaveFile(c *gin.Context, path string, file *multipart.FileHeader, userId int) (string, error) {
	imgName := GenerateFileName(filepath.Ext(file.Filename), userId)

	dstRelative := path + imgName

	if err := c.SaveUploadedFile(file, dstRelative); err != nil {
		return "", fmt.Errorf("can't save file")
	}

	return dstRelative, nil
}

const CODE_LEN = 5

func GenerateFileName(extension string, userId int) string {
	code := services.GenerateCode(CODE_LEN)

	fileName := strconv.Itoa(userId) + "_" + code + extension

	return fileName
}
