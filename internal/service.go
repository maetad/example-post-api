package internal

import (
	"errors"

	"post/internal/config"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Service struct {
	ConnDB *gorm.DB
	Router *gin.Engine
}

func New(opt config.Options) (*Service, error) {
	db, err := gorm.Open(postgres.Open(opt.DatabaseDSN))

	if err != nil {
		return nil, errors.New("cannot connect Database: " + err.Error())
	}

	db.AutoMigrate(
		&Post{},
		&Comment{},
	)

	r := gin.Default()
	r.Use(cors.New(cors.Options{
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "Origin", "X-Request-ID", "X-User-ID", "X-Employee-ID"},
		AllowedOrigins: opt.CORSAllowOrigin,
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"},
	}))

	NewRoute(r, db)

	return &Service{
		ConnDB: db,
		Router: r,
	}, nil
}
