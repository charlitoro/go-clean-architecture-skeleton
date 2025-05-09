package entities

import "time"

type PostType struct {
	Title       string
	Description string
	CreatedAt   time.Time
	IsPublished bool
	UserId      string
}

// Constructor function
func Post(title, description string, createdAt time.Time, userId string, isPublished bool) *PostType {
	return &PostType{
		Title:       title,
		Description: description,
		CreatedAt:   createdAt,
		IsPublished: isPublished,
		UserId:      userId,
	}
}

// Getter methods
func (p *PostType) GetTitle() string        { return p.Title }
func (p *PostType) GetDescription() string  { return p.Description }
func (p *PostType) GetCreatedAt() time.Time { return p.CreatedAt }
func (p *PostType) GetIsPublished() bool       { return p.IsPublished }
func (p *PostType) GetUserId() string       { return p.UserId }
