package model

import (
	"fmt"
	"time"

	"github.com/joomcode/errorx"
	"gorm.io/gorm"
	"pacstall.dev/webserver/types/repository"
)

type ShortenedLink struct {
	gorm.Model
	LinkId string `gorm:"unique"`
	Link   string
	Visits uint `gorm:"default:0"`
}

func (e *ShortenedLink) GetLinkId() string {
	return e.LinkId
}
func (e *ShortenedLink) SetLinkId(v string) {
	e.LinkId = v
}
func (e *ShortenedLink) GetLink() string {
	return e.Link
}
func (e *ShortenedLink) SetLink(v string) {
	e.Link = v
}
func (e *ShortenedLink) GetVisits() uint {
	return e.Visits
}
func (e *ShortenedLink) SetVisits(v uint) {
	e.Visits = v
}
func (e *ShortenedLink) GetCreatedAt() time.Time {
	return e.CreatedAt
}
func (e *ShortenedLink) SetCreatedAt(v time.Time) {
	e.CreatedAt = v
}
func (e *ShortenedLink) GetUpdatedAt() time.Time {
	return e.UpdatedAt
}
func (e *ShortenedLink) SetUpdatedAt(v time.Time) {
	e.UpdatedAt = v
}
func (e *ShortenedLink) GetID() uint {
	return e.ID
}
func (e *ShortenedLink) SetID(v uint) {
	e.ID = v
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

type ShortenedLinkRepository struct {
	database *gorm.DB
}

func InitShortenedLinkRepository(r *ShortenedLinkRepository, db *gorm.DB) *ShortenedLinkRepository {
	if r == nil {
		r = &ShortenedLinkRepository{}
	}

	r.database = db
	return r
}

func (r *ShortenedLinkRepository) FindOneByLinkId(linkId string) (repository.ShortenedLink, error) {
	var shortenedLink ShortenedLink
	if result := r.database.Where(ShortenedLink{LinkId: linkId}).First(&shortenedLink); result.Error != nil {
		return nil, errorx.Decorate(result.Error, "not found")
	}

	return &shortenedLink, nil
}

func (r *ShortenedLinkRepository) IncrementVisits(id uint) error {
	incrementVisitsExpression := gorm.Expr(
		fmt.Sprintf("%v + ?", ShortenedLinkColumns.Visits),
		1,
	)

	entity := ShortenedLink{}
	entity.ID = id
	if err := r.database.Where(&entity).Update(ShortenedLinkColumns.Visits, incrementVisitsExpression).Error; err != nil {
		return errorx.Decorate(err, "failed to increment visits")
	}

	return nil
}
