package userhandler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"encoding/json"

	"uas-pbp/models"
)

// Login godoc
// @Summary Login User
// @Description Login user sebelum CRUD di database 
// @Tags   Login
// @Accept   json
// @Produce  json
// @Param    user body models.User true "Data User"
// @Success 200 {array} models.User
// @Router   /login [post]
func Login(c *gin.Context) {
	var requestBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data tidak ditemukan"})
		return
	}

	var user models.User

	db := models.DB.Where("username = ?", requestBody.Username).First(&user)

	if db.Error != nil {
		if db.Error == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Username salah"})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Terjadi kesalahan saat login"})
		return
	}

	if user.Password != requestBody.Password {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Password salah"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login berhasil"})
}


// Index godoc
// @Summary Get User
// @Description Melihat list index user di database 
// @Tags   User
// @Accept   json
// @Produce  json
// @Param    user body models.User true "Data User"
// @Success 200 {array} models.User
// @Router   /user [get]
func Index(c *gin.Context) {
	var user []models.User

	models.DB.Find(&user)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// Show godoc
// @Summary Select User based on ID
// @Description Melihat list index user berdasarkan ID di database 
// @Tags   User
// @Accept   json
// @Produce  json
// @Param    user body models.User true "Data User"
// @Success 200 {array} models.User
// @Router   /user/:id [get]
func Show(c *gin.Context){
	var user models.User

	id := c.Param("id")
	if err := models.DB.First(&user, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ada"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	models.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// Create godoc
// @Summary Add User
// @Description Membuat user baru di database 
// @Tags   User
// @Accept   json
// @Produce  json
// @Param    user body models.User true "Data User"
// @Success 200 {array} models.User
// @Router   /user [post]
func Create(c *gin.Context){
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// Update godoc
// @Summary Edit the user
// @Description Mengubah user di database 
// @Tags   User
// @Accept   json
// @Produce  json
// @Param    user body models.User true "Data User"
// @Success 200 {array} models.User
// @Router   /user/:id [put]
func Update(c *gin.Context){
	var user models.User
	id := c.Param("id")

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&user).Where("id = ?", id).Updates(&user).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengubah data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diubah"})
}

// Delete godoc
// @Summary Deleting the User
// @Description Menghapus user di database 
// @Tags   User
// @Accept   json
// @Produce  json
// @Param    user body models.User true "Data User"
// @Success 200 {array} models.User
// @Router   /user/:id [delete]
func Delete(c *gin.Context){
	var user models.User

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&user, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
