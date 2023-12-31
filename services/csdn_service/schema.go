package csdn_service

import (
	"github.com/anjude/bean-sdk-go/beansdk"
	"github.com/anjude/bean-sdk-go/beansdk/response"
	"net/http"
)

type HotArticleCommentRequest struct {
	Openid    string `json:"openid"`
	UserName  string `json:"user_name"`
	UserToken string `json:"user_token"`
}

func (h HotArticleCommentRequest) GetURL() string {
	return HotArticleCommentAPI
}

func (h HotArticleCommentRequest) GetMethod() string {
	return http.MethodPost
}

func (h HotArticleCommentRequest) GetHeaders() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
	}
}

func (h HotArticleCommentRequest) GetData() interface{} {
	return h
}

func (h HotArticleCommentRequest) GetVersion() string {
	return beansdk.Version
}

type HotArticleCommentResponse struct {
	response.HttpResponse
	Data string `json:"data"`
}
