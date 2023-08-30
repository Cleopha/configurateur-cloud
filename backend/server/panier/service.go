package panier

import (
	"backend/db"
	"backend/ent"
	"backend/ent/panier"
	"backend/ent/panierinstances"
	"backend/server/common"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Service struct {
	db *db.Client
}

func (s *Service) CreatePanier(ctx context.Context) (*ent.Panier, error) {
	return s.db.Panier.
		Create().
		SetUID(uuid.New()).
		Save(ctx)
}

func (s *Service) GetPanier(ctx context.Context, uid uuid.UUID) (*ent.Panier, error) {
	return s.db.Panier.
		Query().
		Where(panier.UID(uid)).
		WithPanierInstances().
		WithPanierBlockStorage().
		Only(ctx)
}

func (s *Service) GetInstancesFromPanier(ctx context.Context, panierInstanceIDs []int) ([]ent.Instances, error) {
	var instances []ent.Instances
	for _, id := range panierInstanceIDs {
		instance, err := s.db.PanierInstances.
			Query().
			Where(panierinstances.ID(id)).
			QueryInstance().
			Only(ctx)
		if err != nil {
			return nil, err
		}
		instances = append(instances, *instance)
	}
	return instances, nil
}

func (s *Service) GetAllFromPanier(ctx *gin.Context, uid uuid.UUID) ([]ent.Instances, error) {
	logger := common.GetRequestLogger(ctx)

	p, err := s.GetPanier(ctx, uid)
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch panier")
		return nil, err
	}

	var panierInstanceIDs []int
	for _, panierInstance := range p.Edges.PanierInstances {
		panierInstanceIDs = append(panierInstanceIDs, panierInstance.ID)
	}

	instances, err := s.GetInstancesFromPanier(ctx, panierInstanceIDs)
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch instances")
		return nil, err
	}

	return instances, nil
}
