package nats

import (
	"testing"
)

func TestRegisterConstants(t *testing.T) {
	tests := []struct {
		name     string
		got      string
		expected string
	}{
		{"REGISTER_CERTIFICATE_SENT", REGISTER_CERTIFICATE_SENT, "users.certificate_sent"},
		{"REGISTER_COMPLETE", REGISTER_COMPLETE, "users.completed"},
		{"REGISTER_IN_REVIEW", REGISTER_IN_REVIEW, "users.review_request"},
		{"REGISTER_REVOKED", REGISTER_REVOKED, "users.certificate_revoked"},
		{"REGISTER_APPROVED", REGISTER_APPROVED, "users.approved"},
		{"REGISTER_PASSWORD_LINK_SENT", REGISTER_PASSWORD_LINK_SENT, "users.password_link_sent"},
		{"REGISTER_OIDC_FIRST_LOGIN", REGISTER_OIDC_FIRST_LOGIN, "users.oidc_first_login"},
		{"REGISTER_SEND_CERTIFICATE", REGISTER_SEND_CERTIFICATE, "users.send_certificate"},
		{"REGISTER_FORCE_PASSWORD_CHANGE", REGISTER_FORCE_PASSWORD_CHANGE, "users.force_change_password"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.expected {
				t.Errorf("got %q, want %q", tt.got, tt.expected)
			}
		})
	}
}

func TestRegisterPossibleStatus(t *testing.T) {
	statuses := RegisterPossibleStatus()

	if len(statuses) != 3 {
		t.Fatalf("expected 3 statuses, got %d", len(statuses))
	}

	expected := []string{REGISTER_CERTIFICATE_SENT, REGISTER_COMPLETE, REGISTER_IN_REVIEW}
	for i, s := range statuses {
		if s != expected[i] {
			t.Errorf("status[%d] = %q, want %q", i, s, expected[i])
		}
	}
}
