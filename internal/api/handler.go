package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (o *Rest) Register(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

func (o *Rest) Login(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

func (o *Rest) LoadOrder(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

func (o *Rest) ListUsersOrders(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

func (o *Rest) Balance(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

func (o *Rest) Withdraw(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

func (o *Rest) Withdrawals(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
