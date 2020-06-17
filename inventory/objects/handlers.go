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
			return
		}
		c.JSON(http.StatusOK, result)

	}
}

func (o *objects) InventoryCreateObjectHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data InventoryObject
		if err := c.ShouldBindJSON(&data); err == nil {
			data.CreatedBy = c.Request.Header.Get("UserID")
			id, err := o.CreateInventoryObject(c.Request.Context(), &data)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			c.JSON(http.StatusCreated, gin.H{"id": id})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
}

func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}
	c.JSON(code, resp)
	c.Abort()
}
