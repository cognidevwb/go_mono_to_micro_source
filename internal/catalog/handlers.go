package catalog

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes mounts the catalog endpoints.
func RegisterRoutes(r *gin.RouterGroup, svc *Service) {
	r.GET("/products", func(c *gin.Context) { listProducts(c, svc) })
	r.POST("/products", func(c *gin.Context) { createProduct(c, svc) })
}

func listProducts(c *gin.Context, svc *Service) {
	products, err := svc.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func createProduct(c *gin.Context, svc *Service) {
	var req Product
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := svc.Create(&req); err != nil {
		slog.Warn("create product failed", "sku", req.SKU, "err", err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, req)
}
