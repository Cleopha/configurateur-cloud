package panier

import (
	"backend/db"
	"github.com/gin-gonic/gin"
)

func AddEndpoints(group gin.RouterGroup, dbClient *db.Client) {
	panierController := Controller{db: dbClient, panierService: &Service{db: dbClient}}

	paniers := group.Group("/paniers")
	{
		paniers.GET("", panierController.GetPanier)
		paniers.POST("", panierController.CreatePanier)
	}
}
