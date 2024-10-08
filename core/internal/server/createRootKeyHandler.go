package server

import (
	"apikeyper/internal/database"
	"apikeyper/internal/database/utils"
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (r *CreateRootKeyRequest) Valid(ctx context.Context) (problems map[string]string) {
	problems = make(map[string]string)

	if r.Name == "" {
		problems["name"] = "name is required"
	}

	return
}

func (s *Server) CreateRootKeyHandler(w http.ResponseWriter, r *http.Request) {
	// Validate the request body
	decodedJson, problems, err := decodeValid[*CreateRootKeyRequest](r)

	if err != nil {
		encode(w, r, http.StatusBadRequest, problems)
		return
	}

	rootKey := utils.GenerateRandomId("apikeyper_")

	var rootKeyRow = &database.RootKey{
		ID:            uuid.New(),
		WorkspaceId:   decodedJson.WorkspaceId,
		RootHashedKey: utils.HashString(rootKey),
		RootKeyName:   &decodedJson.Name,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	_, err = s.Db.CreateRootKey(rootKeyRow)
	if err != nil {
		encode(w, r, http.StatusInternalServerError, err)
		return
	}

	respBody := CreateRootKeyResponse{
		RootKey: rootKey,
	}

	encode(w, r, http.StatusCreated, respBody)
}
