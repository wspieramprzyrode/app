package main

import (
	"context"
	"log"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/wspieramprzyrode/mobile/inventory/category"
	"github.com/wspieramprzyrode/mobile/inventory/objects"
	"go.uber.org/zap"
	_ "gocloud.dev/docstore/memdocstore"
)

func main() {
	ctx := context.Background()
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalln(err)
	}
	category, err := category.New(ctx)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer category.Close()
	objects, err := objects.New(ctx)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer objects.Close()
	router := gin.New()
	router.Use(ginzap.RecoveryWithZap(logger, true))
	categoriesRoutes := router.Group("/categories")
	{
		categoriesRoutes.Use(ginzap.Ginzap(logger, time.RFC3339, true))
		categoriesRoutes.GET("", category.InventoryObjectCategoriesHandler())
		categoriesRoutes.POST("", category.InventoryCreateCategoryHandler())
	}
	objectsRoutes := router.Group("/objects")
	{
		objectsRoutes.Use(ginzap.Ginzap(logger, time.RFC3339, true))
		objectsRoutes.GET("", objects.InventoryObjectsListHandler())
		objectsRoutes.POST("", objects.InventoryCreateObjectHandler())
	}

	router.GET("/health", HealthGET)
	router.Run(":8080")
}
func HealthGET(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "UP",
	})
}
