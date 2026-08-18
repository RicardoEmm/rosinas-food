package domain

type Material struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size:50;not null;uniqueIndex" json:"name"`
}
