package storage

import (
	"context"
	"github.com/dsxack/go/v2/kitparser/session"
	"github.com/dsxack/go/v2/safe"
	"github.com/reactivex/rxgo/v2"
	"gorm.io/gorm"
)

type GORMDriver struct {
	DB *gorm.DB
}

func NewSQLDriver(dialector gorm.Dialector, config *gorm.Config) (*GORMDriver, error) {
	db, err := gorm.Open(dialector, config)
	if err != nil {
		return nil, err
	}
	return &GORMDriver{DB: db}, nil
}

func (d GORMDriver) Migrate(bucketName string, value interface{}) error {
	migrator := d.DB.Table(bucketName).Migrator()
	if migrator.HasTable(bucketName) {
		return nil
	}
	return migrator.AutoMigrate(value)
}

func (d GORMDriver) Put(ctx context.Context, bucketName string, value interface{}) error {
	if valueModel, ok := value.(model); ok {
		valueModel.setSession(session.From(ctx))
	}

	return d.DB.WithContext(ctx).Table(bucketName).Save(value).Error
}

func (d GORMDriver) Get(ctx context.Context, bucketName string, values interface{}, where ...interface{}) error {
	return d.DB.WithContext(ctx).Table(bucketName).Find(values, where...).Error
}

func (d GORMDriver) GetIterable(ctx context.Context, bucketName string, value interface{}, where ...interface{}) rxgo.Observable {
	var (
		items      = make(chan rxgo.Item)
		observable = rxgo.FromChannel(items)
	)

	query := d.DB.WithContext(ctx).
		Table(bucketName).
		Model(value)

	for i := 0; i < len(where); i += 2 {
		if i+1 >= len(where) {
			query = query.Where(where[i])
		} else {
			query = query.Where(where[i], where[i+1])
		}
	}

	rows, err := query.Rows()
	if err != nil {
		return rxgo.Just(err)()
	}

	safe.Go(func() {
		defer close(items)
		defer rows.Close()

		for rows.Next() {
			err = d.DB.ScanRows(rows, value)
			if err != nil {
				items <- rxgo.Error(err)
			}

			items <- rxgo.Of(value)
		}
	})

	return observable
}
