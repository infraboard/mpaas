package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mpaas/apps/trigger"
)

func TestHandleGitlabEvent(t *testing.T) {
	req := httptest.NewRequest("POST", "/webhook", strings.NewReader(`{"event_type":"push","ref":"refs/heads/feature","project_id":1}`))
	req.Header.Set(trigger.GITLAB_HEADER_EVENT_NAME, "Push Hook")
	req.Header.Set(trigger.GITLAB_HEADER_EVENT_UUID, "1234")
	resp := httptest.NewRecorder()

	impl.HandleGitlabEvent(restful.NewRequest(req), restful.NewResponse(resp))

	if resp.Code != http.StatusOK {
		t.Errorf("Expected response code %d. Got %d.", http.StatusOK, resp.Code)
	}

	var ins []interface{}
	err := json.Unmarshal(resp.Body.Bytes(), &ins)
	if err != nil {
		t.Errorf("Failed to unmarshal response body. Error: %v", err)
	}

	if len(ins) != 1 {
		t.Errorf("Expected 1 instance to be returned from service. Got %d.", len(ins))
	}
}
