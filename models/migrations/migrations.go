package migrations

import (
	"errors"
	"github.com/go-xorm/xorm"
)

type migration func(*xorm.Engine) error

// The version table. Should have only one row with id==1
type Version struct {
	Id      int64 `xorm:"pk"`
	Version int64
}

// This is a sequence of migrations. Add new migrations to the bottom of the list.
// If you want to "retire" a migration, replace it with "expiredMigration"
var migrations = []migration{}

// Migrate database to current version
func Migrate(x *xorm.Engine) error {
	x.Sync(new(Version))

	currentVersion := &Version{Id: 1}
	has, err := x.Get(currentVersion)
	if err != nil {
		return err
	}
	if !has {
		_, err = x.InsertOne(currentVersion)
	}

	v := currentVersion.Version

	for i, migration := range migrations[v:] {
		if err = migration(x); err != nil {
			return err
		}
		currentVersion.Version = v + int64(i) + 1
		x.Id(1).Update(currentVersion)
	}
	return nil
}

func expiredMigration(x *xorm.Engine) error {
	return errors.New("You are migrating from a too old gogs version")
}
