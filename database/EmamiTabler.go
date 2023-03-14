package database

import "gorm.io/gorm/schema"

type EmamiTabler interface {
	schema.Tabler
	TableNiceName() string
}
