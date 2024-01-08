package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Rest struct {
	server *http.Server

	ErrorChan chan error
}

func New() *Rest {
	return &Rest{
		ErrorChan: make(chan error, 1),
	}
}

func (o *Rest) Start(lAddr string) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	o.SetRouters(r)

	o.server = &http.Server{
		Addr:    lAddr,
		Handler: r,
	}

	go func() {
		defer func() {
			if rec := recover(); rec != nil {
				fmt.Println("Panic")
			}
		}()

		err := o.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			o.ErrorChan <- err
		}
	}()
}

func (o *Rest) Stop(ctx context.Context) error {
	defer close(o.ErrorChan)

	err := o.server.Shutdown(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (o *Rest) SetRouters(r *gin.Engine) {
	apiGroup := r.Group("/api")
	userGroup := apiGroup.Group("/user")

	userGroup.POST("/register", o.Register)
	userGroup.POST("/login", o.Login)
	userGroup.POST("/orders", o.LoadOrder)
	userGroup.GET("/orders", o.ListUsersOrders)
	userGroup.GET("/balance", o.Balance)
	userGroup.POST("/balance/withdraw", o.Withdraw)
	userGroup.POST("/withdrawals", o.Withdrawals)
}
