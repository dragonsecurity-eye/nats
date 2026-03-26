package nats

import (
	"testing"
)

func TestSystemUpdateConstants(t *testing.T) {
	tests := []struct {
		name     string
		got      string
		expected string
	}{
		{"UNKNOWN", UNKNOWN, "systemupdate.unknown"},
		{"NOT_CONFIGURED", NOT_CONFIGURED, "systemupdate.not_configured"},
		{"DISABLED", DISABLED, "systemupdate.disabled"},
		{"NOTIFY_BEFORE_DOWNLOAD", NOTIFY_BEFORE_DOWNLOAD, "systemupdate.notify_before_download"},
		{"NOTIFY_BEFORE_INSTALLATION", NOTIFY_BEFORE_INSTALLATION, "systemupdate.notify_before_installation"},
		{"NOTIFY_SCHEDULED_INSTALLATION", NOTIFY_SCHEDULED_INSTALLATION, "systemupdate.notify_scheduled_installation"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.expected {
				t.Errorf("got %q, want %q", tt.got, tt.expected)
			}
		})
	}
}

func TestSystemUpdatePossibleStatus(t *testing.T) {
	statuses := SystemUpdatePossibleStatus()

	if len(statuses) != 6 {
		t.Fatalf("expected 6 statuses, got %d", len(statuses))
	}

	expected := []string{UNKNOWN, NOT_CONFIGURED, DISABLED, NOTIFY_BEFORE_DOWNLOAD, NOTIFY_BEFORE_INSTALLATION, NOTIFY_SCHEDULED_INSTALLATION}
	for i, s := range statuses {
		if s != expected[i] {
			t.Errorf("status[%d] = %q, want %q", i, s, expected[i])
		}
	}
}
