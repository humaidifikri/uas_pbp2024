package produkhandler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"encoding/json"

	"uas-pbp/models"
)

// Index godoc
// @Summary Get Produk
// @Description Melihat list index produk di database 
// @Tags   Produk 
// @Accept   json
// @Produce  json
// @Param    produk body models.Produk true "Data Produk"
// @Success 200 {array} models.Produk
// @Router   /produk [get]
func Index(c *gin.Context) {
	var produk []models.Produk

	models.DB.Find(&produk)
	c.JSON(http.StatusOK, gin.H{"produk": produk})
}

// Show godoc
// @Summary Select Produk based on ID
// @Description Melihat list index berdasarkan ID produk di database 
// @Tags   Produk 
// @Accept   json
// @Produce  json
// @Param    produk body models.Produk true "Data Produk"
// @Success 200 {array} models.Produk
// @Router   /produk/:id [get]
func Show(c *gin.Context) {
	var produk models.Produk
	id := c.Param("id")

	if err := models.DB.First(&produk, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ada"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"produk": produk})
}

// Create godoc
// @Summary Create a new produk
// @Description Membuat produk baru di database 
// @Tags   Produk 
// @Accept   json
// @Produce  json
// @Param    produk body models.Produk true "Data Produk"
// @Success 200 {array} models.Produk
// @Router   /produk [post]
func Create(c *gin.Context) {
	var produk models.Produk

	if err := c.ShouldBindJSON(&produk); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&produk)
	c.JSON(http.StatusOK, gin.H{"product": produk})
	
}

// Update godoc
// @Summary Edit the produk
// @Description Mengubah produk di database 
// @Tags   Produk 
// @Accept   json
// @Produce  json
// @Param    produk body models.Produk true "Data Produk"
// @Success 200 {array} models.Produk
// @Router   /produk/:id [put]
func Update(c *gin.Context) {
	var produk models.Produk
	id := c.Param("id")

	if err := c.ShouldBindJSON(&produk); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&produk).Where("id = ?", id).Updates(&produk).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengubah produk"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produk berhasil diubah"})
}

// Delete godoc
// @Summary Deleting the produk
// @Description Menghapus produk di database 
// @Tags   Produk 
// @Accept   json
// @Produce  json
// @Param    produk body models.Produk true "Data Produk"
// @Success 200 {array} models.Produk
// @Router   /produk/:id [delete]
func Delete(c *gin.Context) {
	var produk models.Produk

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&produk, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus produk"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produk berhasil dihapus"})
}