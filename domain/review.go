package domain

import (
	"time"

	"github.com/google/jsonapi"

	"fmt"
)

// Review A product review
type Review struct {
	ID             int       `jsonapi:"primary,reviews"`
	AuthorID       string    `jsonapi:"attr,author_id"`
	ReviewText     string    `jsonapi:"attr,review_text"`
	SubmissionTime time.Time `jsonapi:"attr,submission_time"`
}

// JSONAPILinks JSON API styled Links for Review
func (review Review) JSONAPILinks() *jsonapi.Links {
	return &jsonapi.Links{
		"self": fmt.Sprintf("https://example.com/review/%d", review.ID),
	}
}
