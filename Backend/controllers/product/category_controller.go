package controllers

import (
	"net/http"
	"strconv"

	service "proyecto/ArquiSoftware/services/product"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetCategoryById(c *gin.Context) {
	log.Debug("Category id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	categoryDto, err := service.CategoryService.GetCategoryById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, categoryDto)
		return
	}

	c.JSON(http.StatusOK, categoryDto)
}

func GetCategories(c *gin.Context) {
	categoriesDto, err := service.CategoryService.GetCategories()

	if err != nil {
		c.JSON(http.StatusBadRequest, categoriesDto)
		return
	}

	c.JSON(http.StatusOK, categoriesDto)
}
