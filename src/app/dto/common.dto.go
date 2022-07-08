package dto

type ResponseErr struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type BadReqErrResponse struct {
	Message     string      `json:"message"`
	FailedField string      `json:"failed_field"`
	Value       interface{} `json:"value"`
}

// For docs

type ResponseBadRequestErr struct {
	StatusCode int                 `json:"status_code" example:"400"`
	Message    string              `json:"message" example:"Invalid request body"`
	Data       []BadReqErrResponse `json:"data"`
}

type ResponseUnauthorizedErr struct {
	StatusCode int         `json:"status_code" example:"401"`
	Message    string      `json:"message" example:"Invalid token"`
	Data       interface{} `json:"data"`
}

type ResponseForbiddenErr struct {
	StatusCode int         `json:"status_code" example:"403"`
	Message    string      `json:"message" example:"Insufficiency permission"`
	Data       interface{} `json:"data"`
}

type ResponseNotfoundErr struct {
	StatusCode int         `json:"status_code" example:"404"`
	Message    string      `json:"message" example:"Not found"`
	Data       interface{} `json:"data"`
}

type ResponseInternalErr struct {
	StatusCode int         `json:"status_code" example:"500"`
	Message    string      `json:"message" example:"Internal service error"`
	Data       interface{} `json:"data"`
}

type ResponseServiceDownErr struct {
	StatusCode int         `json:"status_code" example:"503"`
	Message    string      `json:"message" example:"Service is down"`
	Data       interface{} `json:"data"`
}
