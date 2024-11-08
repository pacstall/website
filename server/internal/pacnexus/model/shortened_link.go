package model

import "gorm.io/gorm"

type ShortenedLink struct {
	gorm.Model
	LinkId string `gorm:"unique"`
	Link   string
	Visits uint `gorm:"default:0"`
}

var ShortenedLinkColumns = struct {
	LinkId    string
	Link      string
	Visits    string
	CreatedAt string
	UpdatedAt string
	ID        string
}{
	LinkId:    "link_id",
	Link:      "link",
	Visits:    "visits",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	ID:        "id",
}
