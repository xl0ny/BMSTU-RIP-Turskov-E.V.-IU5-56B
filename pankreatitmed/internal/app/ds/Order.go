package ds

import "time"

type Order struct {
	ID             uint      `gorm:"primaryKey"`
	Status         string    `gorm:"size:16;not null"`
	CreatedAt      time.Time `gorm:"not null;autoCreateTime"`
	CreatorID      uint      `gorm:"not null"`
	FormedAt       *time.Time
	FinishedAt     *time.Time
	ModeratorID    *uint
	ComputedResult *string
}
