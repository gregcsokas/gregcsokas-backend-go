package newsletter

import "time"

type Subscription struct {
	UUID      string    `gorm:"primary_key" json:"uuid"`
	Email     string    `gorm:"unique" json:"email"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt time.Time `gorm:"default:NULL" json:"deleted_at"`
}
