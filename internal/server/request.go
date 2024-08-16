package server

import "github.com/google/uuid"

type CreateRootKeyRequest struct {
	Name        string    `json:"name"`
	WorkspaceId uuid.UUID `json:"workspaceId"`
	Permissions []string  `json:"permissions"`
}

type CreateApiRequest struct {
	ApiName string `json:"apiName"`
}

type CreateApiKeyRequest struct {
	ApiId  uuid.UUID `json:"apiId"`
	Name   string    `json:"name"`
	Prefix string    `json:"prefix"`
	Roles  []string  `json:"roles"`
}