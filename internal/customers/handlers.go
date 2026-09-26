package customers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes mounts the customer endpoints.
func RegisterRoutes(r *gin.RouterGroup, svc *Service) {
	r.GET("/customers/:id", func(c *gin.Context) { getCustomer(c, svc) })
	r.POST("/customers", func(c *gin.Context) { registerCustomer(c, svc) })
}

func getCustomer(c *gin.Context, svc *Service) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad id"})
		return
	}
	cust, err := svc.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, cust)
}

func registerCustomer(c *gin.Context, svc *Service) {
	var req Customer
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := svc.Register(&req); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, req)
}
