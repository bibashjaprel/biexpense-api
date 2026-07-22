package router

import (
	"github.com/bibashjaprel/biexpense-api/internal/modules/accounts"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/gin-gonic/gin"
)

func Setup(db *pgxpool.Pool) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "biexpense-api is running",
		})
	})

	accountRepository := accounts.NewRepository(db)
	accountService := accounts.NewService(accountRepository)
	accountHandler := accounts.NewHandler(accountService)

	api := r.Group("/api/v1")

	{
		api.GET("/accounts", accountHandler.GetAll)
	}

	return r
}
