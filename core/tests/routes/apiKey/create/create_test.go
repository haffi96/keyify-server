package tests

import (
	"apikeyper/internal/database"
	"apikeyper/internal/database/utils"
	"apikeyper/internal/events"
	ApikeyperServer "apikeyper/internal/server"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"apikeyper/tests"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateApiKey(t *testing.T) {
	// Create a new service
	s := &ApikeyperServer.Server{
		Db:      database.New(),
		Message: events.New(),
	}

	server := httptest.NewServer(
		http.HandlerFunc(
			ApikeyperServer.Auth(
				s.Db, s.CreateApiKeyHandler,
			),
		),
	)

	defer server.Close()

	// Create a workspace
	workspaceId, _ := s.Db.CreateWorkspace(&database.Workspace{
		ID:            uuid.New(),
		WorkspaceName: "test-workspace",
	})

	// Create root key
	rootKey := "test-root-key"
	s.Db.CreateRootKey(&database.RootKey{
		ID:            uuid.New(),
		WorkspaceId:   workspaceId,
		RootHashedKey: utils.HashString(rootKey),
	})

	// Create an API
	apiId := uuid.New()
	s.Db.CreateApi(&database.Api{
		ID:          apiId,
		WorkspaceId: workspaceId,
	})

	createApiKeyReq := ApikeyperServer.CreateApiKeyRequest{
		ApiId: apiId,
	}
	var buf bytes.Buffer
	_ = json.NewEncoder(&buf).Encode(createApiKeyReq)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", server.URL, &buf)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", rootKey))

	resp, err := client.Do(req)

	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}
	defer resp.Body.Close()

	// Assertions
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	// Cleanup db
	defer tests.CleanupDb()
}
