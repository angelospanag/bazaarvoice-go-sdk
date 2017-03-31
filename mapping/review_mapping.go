package mapping

import (
	"strconv"
	"time"

	"github.com/angelospanag/bazaarvoice-go-sdk/domain"
)

type ReviewsFromBV struct {
	Limit        int    `json:"Limit"`
	Offset       int    `json:"Offset"`
	TotalResults int    `json:"TotalResults"`
	Locale       string `json:"Locale"`
	Results      []struct {
		ID                         string    `json:"Id"`
		CID                        string    `json:"CID"`
		SourceClient               string    `json:"SourceClient"`
		LastModificationTime       time.Time `json:"LastModificationTime"`
		LastModeratedTime          time.Time `json:"LastModeratedTime"`
		ProductID                  string    `json:"ProductId"`
		AuthorID                   string    `json:"AuthorId"`
		ContentLocale              string    `json:"ContentLocale"`
		IsFeatured                 bool      `json:"IsFeatured"`
		TotalClientResponseCount   int       `json:"TotalClientResponseCount"`
		TotalCommentCount          int       `json:"TotalCommentCount"`
		Rating                     int       `json:"Rating"`
		SecondaryRatingsOrder      []string  `json:"SecondaryRatingsOrder"`
		IsRatingsOnly              bool      `json:"IsRatingsOnly"`
		IsRecommended              bool      `json:"IsRecommended"`
		TotalFeedbackCount         int       `json:"TotalFeedbackCount"`
		TotalNegativeFeedbackCount int       `json:"TotalNegativeFeedbackCount"`
		TotalPositiveFeedbackCount int       `json:"TotalPositiveFeedbackCount"`
		ModerationStatus           string    `json:"ModerationStatus"`
		SubmissionID               string    `json:"SubmissionId"`
		SubmissionTime             time.Time `json:"SubmissionTime"`
		ReviewText                 string    `json:"ReviewText"`
		Title                      string    `json:"Title"`
		UserNickname               string    `json:"UserNickname"`
		SecondaryRatings           struct {
			Quality struct {
				Value       int         `json:"Value"`
				ID          string      `json:"Id"`
				MaxLabel    interface{} `json:"MaxLabel"`
				DisplayType string      `json:"DisplayType"`
				Label       string      `json:"Label"`
				MinLabel    interface{} `json:"MinLabel"`
				ValueRange  int         `json:"ValueRange"`
				ValueLabel  interface{} `json:"ValueLabel"`
			} `json:"Quality"`
			Value struct {
				Value       int         `json:"Value"`
				ID          string      `json:"Id"`
				MaxLabel    interface{} `json:"MaxLabel"`
				DisplayType string      `json:"DisplayType"`
				Label       string      `json:"Label"`
				MinLabel    interface{} `json:"MinLabel"`
				ValueRange  int         `json:"ValueRange"`
				ValueLabel  interface{} `json:"ValueLabel"`
			} `json:"Value"`
		} `json:"SecondaryRatings"`
		InappropriateFeedbackList []interface{} `json:"InappropriateFeedbackList"`
		ContextDataValuesOrder    []interface{} `json:"ContextDataValuesOrder"`
		Helpfulness               interface{}   `json:"Helpfulness"`
		Photos                    []interface{} `json:"Photos"`
		RatingRange               int           `json:"RatingRange"`
		Badges                    struct {
		} `json:"Badges"`
		CommentIds               []interface{} `json:"CommentIds"`
		UserLocation             interface{}   `json:"UserLocation"`
		BadgesOrder              []interface{} `json:"BadgesOrder"`
		ClientResponses          []interface{} `json:"ClientResponses"`
		CampaignID               interface{}   `json:"CampaignId"`
		Cons                     interface{}   `json:"Cons"`
		Videos                   []interface{} `json:"Videos"`
		ProductRecommendationIds []interface{} `json:"ProductRecommendationIds"`
		AdditionalFieldsOrder    []interface{} `json:"AdditionalFieldsOrder"`
		IsSyndicated             bool          `json:"IsSyndicated"`
		Pros                     interface{}   `json:"Pros"`
		TagDimensionsOrder       []interface{} `json:"TagDimensionsOrder"`
		ContextDataValues        struct {
		} `json:"ContextDataValues"`
		TagDimensions struct {
		} `json:"TagDimensions"`
		AdditionalFields struct {
		} `json:"AdditionalFields"`
	} `json:"Results"`
	Includes struct {
	} `json:"Includes"`
	HasErrors bool          `json:"HasErrors"`
	Errors    []interface{} `json:"Errors"`
}

func ReviewsMapping(reviewsFromBV ReviewsFromBV) []*domain.Review {

	reviews := []*domain.Review{}

	if len(reviewsFromBV.Results) > 0 {

		for _, reviewFromBV := range reviewsFromBV.Results {

			productReview := domain.Review{}

			productReview.AuthorID = reviewFromBV.AuthorID
			productReview.ReviewText = reviewFromBV.ReviewText
			productReview.SubmissionTime = reviewFromBV.SubmissionTime
			productReview.ID, _ = strconv.Atoi(reviewFromBV.ID)

			reviews = append(reviews, &productReview)
		}
	}
	return reviews
}
