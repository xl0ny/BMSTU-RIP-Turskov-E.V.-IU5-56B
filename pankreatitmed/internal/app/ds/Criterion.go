package ds

type Criterion struct {
	ID          uint    `gorm:"primaryKey"`
	Code        string  `gorm:"size:8;not null"`
	Name        string  `gorm:"size:120;not null"`
	Indicator   string  `gorm:"size:80;not null"`
	Duration    string  `gorm:"size:60;not null"`
	HomeVisit   bool    `gorm:"not null;default:false"`
	ImageURL    *string `gorm:"not null"`
	Description string  `gorm:"not null"`
	IsActive    bool    `gorm:"not null;default:true"`
}
