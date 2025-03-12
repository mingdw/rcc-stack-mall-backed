package server

import (
	"context"
	"os"
	"rcc-stake-mall-backed/internal/model"
	"rcc-stake-mall-backed/pkg/log"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type MigrateServer struct {
	db  *gorm.DB
	log *log.Logger
}

func NewMigrateServer(db *gorm.DB, log *log.Logger) *MigrateServer {
	return &MigrateServer{
		db:  db,
		log: log,
	}
}
func (m *MigrateServer) Start(ctx context.Context) error {
	if err := m.db.AutoMigrate(
		&model.User{},
	); err != nil {
		m.log.Error("user migrate error", zap.Error(err))
		return err
	}
	m.log.Info("AutoMigrate success")
	os.Exit(0)
	return nil
}
func (m *MigrateServer) Stop(ctx context.Context) error {
	m.log.Info("AutoMigrate stop")
	return nil
}
