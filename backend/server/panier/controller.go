package panier

import (
	"backend/db"
	"backend/ent"
	"backend/server/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type Controller struct {
	db            *db.Client
	panierService *Service
}

func (c *Controller) CreatePanier(ctx *gin.Context) {
	logger := common.GetRequestLogger(ctx)

	newPanier, err := c.panierService.CreatePanier(ctx)
	if err != nil {
		logger.Error().Err(err).Msg("failed to create panier")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create panier"})
		return
	}

	logger.Info().Str("uid", newPanier.UID.String()).Msg("panier created successfully")
	ctx.JSON(http.StatusOK, gin.H{"uid": newPanier.UID.String()})
}

type PanierResponse struct {
	UID                string                   `json:"uid"`
	PanierInstances    []ent.Instances          `json:"instances"`
	PanierBlockStorage []ent.PanierBlockStorage `json:"panierBlockStorage"`
}

func (c *Controller) GetPanier(ctx *gin.Context) {
	logger := common.GetRequestLogger(ctx)

	logger.Info().Msg("get panier")

	uid := ctx.Query("uid")
	uidUUID, err := uuid.Parse(uid)
	if err != nil {
		logger.Error().Err(err).Msg("failed to parse UUID")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse UUID"})
		return
	}

	instances, err := c.panierService.GetAllFromPanier(ctx, uidUUID)
	if err != nil {
		logger.Error().Err(err).Msg("failed to get all panier")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get all panier"})
		return
	}

	response := PanierResponse{
		UID:             uid,
		PanierInstances: instances,
		// ... autres champs
	}

	ctx.JSON(http.StatusOK, response)
}
