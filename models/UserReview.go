package models

import "mime/multipart"

type UserReview struct {
	FirstName     string
	LastName      string
	ContactNumber string

	Service        string
	Quality        int64
	Recommendation int64
	Notes          string
	Image          *multipart.FileHeader
}
