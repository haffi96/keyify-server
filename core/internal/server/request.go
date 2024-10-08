package server

import (
	"github.com/google/uuid"
)

type CreateWorkspaceRequest struct {
	UserGithubId string `json:"userGithubId"`
	Name         string `json:"name"`
}

type CreateRootKeyRequest struct {
	Name        string    `json:"name"`
	WorkspaceId uuid.UUID `json:"workspaceId"`
	Permissions []string  `json:"permissions"`
}

type CreateApiRequest struct {
	ApiName string `json:"apiName"`
}

type ApiKeyRateLimitConfigRequest struct {
	Limit         int    `json:"limit"`
	LimitPeriod   string `json:"period"`
	CounterWindow string `json:"window"`
}

type CreateApiKeyRequest struct {
	ApiId     uuid.UUID                    `json:"apiId"`
	Name      string                       `json:"name"`
	Prefix    string                       `json:"prefix"`
	Roles     []string                     `json:"roles"`
	RateLimit ApiKeyRateLimitConfigRequest `json:"rateLimit"`
}

type VerifyApiKeyRequest struct {
	ApiKey string
	ApiId  uuid.UUID
}

type RevokeApiKeyRequest struct {
	ApiKey string
	ApiId  uuid.UUID
}
