package model

import "time"

type RestResponse struct {
	Status     int         `json:"-"`
	Timestamps time.Time   `json:"timestamps"`
	Data       interface{} `json:"data"`
}

func NewRestResponse(status int, data interface{}) *RestResponse {
	return &RestResponse{
		Status:     status,
		Timestamps: time.Now(),
		Data:       data,
	}
}
