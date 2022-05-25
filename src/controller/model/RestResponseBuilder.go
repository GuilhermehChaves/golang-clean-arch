package model

import "time"

type RestResponseBuilder struct {
	restResponse *RestResponse
}

func NewRestResponseBuilder() *RestResponseBuilder {
	return &RestResponseBuilder{
		restResponse: new(RestResponse),
	}
}

func (r *RestResponseBuilder) WithStatus(status int) *RestResponseBuilder {
	r.restResponse.Status = status
	return r
}

func (r *RestResponseBuilder) WithTimestamps() *RestResponseBuilder {
	r.restResponse.Timestamps = time.Now()
	return r
}

func (r *RestResponseBuilder) WithData(data interface{}) *RestResponseBuilder {
	r.restResponse.Data = data
	return r
}

func (r *RestResponseBuilder) Build() *RestResponse {
	return r.restResponse
}
