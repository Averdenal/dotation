package Models

import (
	"gorm.io/gorm"
	"time"
)

type Dotation struct {
	gorm.Model
	Date time.Time
}
