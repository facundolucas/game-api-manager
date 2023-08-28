package controllers

import (
	"fmt"
	"game-api-manager/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func SaveMoneda(c *gin.Context) {

	var input models.Moneda

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := input.SaveMoneda()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "moneda saved successfully"})
}

func UploadImagenMoneda(c *gin.Context) {

	id := c.Param("id")
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo obtener el archivo"})
		return
	}

	// Crear directorio de imágenes si no existe
	imageDir := "./images/"
	if _, err := os.Stat(imageDir); os.IsNotExist(err) {
		os.Mkdir(imageDir, 0755)
	}

	// Generar nombre de archivo único y guardarlo
	fileName := fmt.Sprintf("%s_%s", id, file.Filename)
	filePath := fmt.Sprintf("%s%s", imageDir, fileName)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el archivo"})
		return
	}

	// Actualizar la URL de la imagen en la base de datos
	moneda, err := models.GetMonedaByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	moneda.ImageURL = filePath
	_, err = moneda.UpdateMoneda()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Imagen cargada exitosamente"})
}

func UploadImagenMonedaModels(c *gin.Context) {

	id := c.Param("id")
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo obtener el archivo"})
		return
	}

	// Crear directorio de imágenes si no existe
	imageDir := "./images/"
	if _, err := os.Stat(imageDir); os.IsNotExist(err) {
		os.Mkdir(imageDir, 0755)
	}

	// Generar nombre de archivo único y guardarlo
	fileName := fmt.Sprintf("%s_%s", id, file.Filename)
	filePath := fmt.Sprintf("%s%s", imageDir, fileName)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el archivo"})
		return
	}

	// Actualizar la URL de la imagen en la base de datos
	moneda, err := models.GetMonedaByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	var model models.MonedaModel
	model.Path = filePath
	moneda.Models = append(moneda.Models, model)
	_, err = moneda.UpdateMoneda()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Imagen cargada exitosamente"})
}

func GetMonedasByTipo(c *gin.Context) {

	tipo := c.Param("tipo")

	monedas, err := models.GetMonedasOrderedByValue(tipo)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, monedas)
}
