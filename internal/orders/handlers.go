package orders

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/acme/shop/internal/platform"
)

// RegisterRoutes mounts the order endpoints.
func RegisterRoutes(r *gin.RouterGroup, svc *Service) {
	r.POST("/orders", func(c *gin.Context) { placeOrder(c, svc) })
	r.GET("/orders/mine", func(c *gin.Context) { myOrders(c, svc) })
}

func placeOrder(c *gin.Context, svc *Service) {
	var req PlaceOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order, err := svc.PlaceOrder(req)
	if err != nil {
		slog.Error("place order failed", "customer", req.CustomerID, "err", err)
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, order)
}

func myOrders(c *gin.Context, svc *Service) {
	userID := platform.UserID(c.Request.Context())
	list, err := svc.ForCustomer(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}
