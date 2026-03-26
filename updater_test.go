package nats

import (
	"testing"
)

func TestUpdateConstants(t *testing.T) {
	tests := []struct {
		name     string
		got      string
		expected string
	}{
		{"UPDATE_ERROR", UPDATE_ERROR, "admin.update.agents.task_status_error"},
		{"UPDATE_PENDING", UPDATE_PENDING, "admin.update.agents.task_status_pending"},
		{"UPDATE_SUCCESS", UPDATE_SUCCESS, "admin.update.agents.task_status_success"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.expected {
				t.Errorf("got %q, want %q", tt.got, tt.expected)
			}
		})
	}
}

func TestTaskUpdatePossibleStatus(t *testing.T) {
	statuses := TaskUpdatePossibleStatus()

	if len(statuses) != 3 {
		t.Fatalf("expected 3 statuses, got %d", len(statuses))
	}

	expected := []string{UPDATE_ERROR, UPDATE_PENDING, UPDATE_SUCCESS}
	for i, s := range statuses {
		if s != expected[i] {
			t.Errorf("status[%d] = %q, want %q", i, s, expected[i])
		}
	}
}
