package global

import (
	"database/sql"

	"gorm.io/gorm"
)

var GormDB *gorm.DB // Gorm 用グローバル変数
var Db *sql.DB      // MySQL 用グローバル変数
