package tests

import (
	"apikeyper/internal/database"
	"apikeyper/internal/database/utils"
	"apikeyper/internal/events"
	"apikeyper/internal/ratelimit"
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

func TestVerifyApiKeyHandler(t *testing.T) {
	// Create a new service
	s := &ApikeyperServer.Server{
		Db:          database.New(),
		Message:     events.New(),
		RateLimiter: ratelimit.New(),
	}

	server := httptest.NewServer(
		http.HandlerFunc(
			ApikeyperServer.Auth(
				s.Db, s.VerifyApiKeyHandler,
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

	// Create an API key
	generatedApiKey, err := utils.GenerateApiKey("test_")
	if err != nil {
		t.Fatalf("error generating api key. Err: %v", err)
	}
	prefix := "test_"
	name := "test-api-key"
	s.Db.CreateApiKey(&database.ApiKey{
		ID:        uuid.New(),
		ApiId:     apiId,
		HashedKey: utils.HashString(generatedApiKey),
		Name:      &name,
		Prefix:    &prefix,
	})

	createApiKeyReq := ApikeyperServer.VerifyApiKeyRequest{
		ApiKey: generatedApiKey,
		ApiId:  apiId,
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
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Cleanup db
	defer tests.CleanupDb()
}
