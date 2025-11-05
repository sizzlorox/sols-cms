package database

import (
	"context"
	"fmt"
	"time"

	"github.com/sizzlorox/sols-cms/pkg/providers/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/plugin/prometheus"
)

type IRepository interface {
	Close() error
	FindOne(dest any, conds ...any) any
	FindMany(dest any, limit int, offset int, conds ...any) any
	InsertOne(value any) error
	InsertMany(values []any) error
	UpdateOne(value any, conds ...any) error
	UpdateMany(values []any, conds ...any) error
	DeleteOne(value any, conds ...any) error
	DeleteMany(value []any, conds ...any) error
}

type DatabaseProvider struct {
	DB     *gorm.DB
	Config *config.ConfigProvider
}

func NewDatabaseProvider(dsn string, config *config.ConfigProvider) (*DatabaseProvider, error) {
	dsn = config.GetDSN()
	db := &gorm.DB{}
	var err error
	switch config.DB_TYPE {
	case "postgres":
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			CreateBatchSize: 100,
		})
	case "mysql":
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			CreateBatchSize: 100,
		})
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(config.DB_NAME+".db"), &gorm.Config{
			CreateBatchSize: 100,
		})
	default:
		return nil, fmt.Errorf("unsupported database: %s", config.Getenv("DB_TYPE"))
	}
	if err != nil {
		return nil, err
	}

	if config.PROMETHEUS_ENABLED {
		db.Use(prometheus.New(prometheus.Config{
			DBName:           config.DB_NAME,
			RefreshInterval:  15,
			PushAddr:         config.PROMETHEUS_ADDR,
			StartServer:      true,
			HTTPServerPort:   uint32(config.PROMETHEUS_PORT),
			MetricsCollector: []prometheus.MetricsCollector{
				// &prometheus.MySQL{
				// 	VariableNames: []string{"Threads_running"},
				// },
			},
		}))
	}

	return &DatabaseProvider{DB: db, Config: config}, nil
}

func (p *DatabaseProvider) Close() error {
	sqlDB, err := p.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (p *DatabaseProvider) FindOne(dest any, conds ...any) any {
	ctx := context.Background()
	result := map[string]any{}
	p.DB.WithContext(ctx).Model(&dest).First(result, conds...)
	return result
}

func (p *DatabaseProvider) FindMany(dest any, limit int, offset int, conds ...any) any {
	if limit <= 0 {
		limit = 10
	}
	ctx := context.Background()
	result := []map[string]any{}
	p.DB.WithContext(ctx).Model(&dest).Limit(limit).Offset(offset).Find(dest, conds...)
	return result
}

func (p *DatabaseProvider) InsertOne(value any) error {
	result := p.DB.Create(value)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *DatabaseProvider) InsertMany(values []any) error {
	result := p.DB.Create(values)
	return result.Error
}

// TODO: Maybe go genereics API so we can return rows affected and error
func (p *DatabaseProvider) UpdateOne(value any, conds ...any) error {
	// TODO: Doc doesnt say anything about result.Error, but check if it exists for this
	result := p.DB.Model(value).Where(conds[0], conds[1:]...).Updates(value)
	return result.Error
}

// TODO: Maybe go genereics API so we can return rows affected and error
func (p *DatabaseProvider) UpdateMany(values []any, conds ...any) error {
	result := p.DB.Model(values).Where(conds[0], conds[1:]...).Updates(values)
	return result.Error
}

func (p *DatabaseProvider) DeleteOne(value any, conds ...any) error {
	now := time.Now()
	newValue := map[string]any{
		"is_active":  false,
		"deleted_at": &now,
	}

	err := p.UpdateOne(newValue, conds...)
	return err
}

func (p *DatabaseProvider) DeleteMany(values []any, conds ...any) any {
	now := time.Now()
	for i := range values {
		values[i] = map[string]any{
			"is_active":  false,
			"deleted_at": &now,
		}
	}

	err := p.UpdateMany(values, conds...)
	return err
}

// TODO: Add hard delete to a cron job that runs periodically to clean up soft deleted records
// Preferably older than 2 months
