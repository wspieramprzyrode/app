package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (cat *category) InventoryObjectCategoriesHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := cat.list(c.Request.Context())
		if err != nil {
			c.AbortWithError(500, err)
		}
		c.JSON(http.StatusOK, result)

	}
}

func (cat *category) InventoryCreateCategoryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data Category
		if err := c.ShouldBindJSON(&data); err == nil {
			id, err := cat.CreateInventoryCategory(c.Request.Context(), &data)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
			}
			c.JSON(http.StatusOK, gin.H{"id": id})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}
}
