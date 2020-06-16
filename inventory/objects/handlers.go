package objects

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (o *objects) InventoryObjectsListHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := o.list(c.Request.Context())
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusOK, result)

	}
}

func (o *objects) InventoryCreateObjectHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data InventoryObject
		if err := c.ShouldBindJSON(&data); err == nil {
			id, err := o.CreateInventoryObject(c.Request.Context(), &data)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
			}
			c.JSON(http.StatusOK, gin.H{"id": id})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}
}
