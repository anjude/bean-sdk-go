package csdn_service

import (
	"github.com/anjude/bean-sdk-go/beansdk"
	"github.com/anjude/bean-sdk-go/services/enum"
)

type Service struct {
	ServiceName enum.ServiceName
	client      *beansdk.Client
	UserName    string
	UserToken   string
}

func NewCsdnService(client *beansdk.Client, userName, userToken string) *Service {
	return &Service{
		ServiceName: enum.CsdnService,
		client:      client,
		UserName:    userName,
		UserToken:   userToken,
	}
}

func (s *Service) HotArticleComment() (*HotArticleCommentResponse, error) {
	resp := &HotArticleCommentResponse{}
	err := s.client.Send(HotArticleCommentRequest{
		Openid:    s.client.Credential.GetSecretId(),
		UserName:  s.UserName,
		UserToken: s.UserToken,
	}, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
