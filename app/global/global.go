package global

import (
	"database/sql"

	"gorm.io/gorm"
)

// DB
var GormDB *gorm.DB
var Db *sql.DB
