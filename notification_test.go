package nats

import (
	"encoding/json"
	"testing"
)

func TestNotificationMarshalJSON(t *testing.T) {
	notif := Notification{
		From:             "admin@example.com",
		To:               "user@example.com",
		Subject:          "Certificate Ready",
		MessageTitle:     "Your certificate is ready",
		MessageGreeting:  "Hello",
		MessageText:      "Your client certificate has been generated.",
		MessageAction:    "Download",
		MessageActionURL: "https://example.com/download",
	}

	data, err := json.Marshal(notif)
	if err != nil {
		t.Fatalf("failed to marshal Notification: %v", err)
	}

	var decoded Notification
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("failed to unmarshal Notification: %v", err)
	}

	if decoded.From != notif.From {
		t.Errorf("From = %q, want %q", decoded.From, notif.From)
	}
	if decoded.To != notif.To {
		t.Errorf("To = %q, want %q", decoded.To, notif.To)
	}
	if decoded.Subject != notif.Subject {
		t.Errorf("Subject = %q, want %q", decoded.Subject, notif.Subject)
	}
	if decoded.MessageActionURL != notif.MessageActionURL {
		t.Errorf("MessageActionURL = %q, want %q", decoded.MessageActionURL, notif.MessageActionURL)
	}
}

func TestNotificationWithAttachments(t *testing.T) {
	notif := Notification{
		From:                   "admin@example.com",
		To:                     "user@example.com",
		Subject:                "Certificates",
		MessageAttachFileName:  "cert.pem",
		MessageAttachFile:      "base64encodeddata",
		MessageAttachFileName2: "key.pem",
		MessageAttachFile2:     "base64encodedkey",
	}

	data, err := json.Marshal(notif)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var decoded Notification
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if decoded.MessageAttachFileName != "cert.pem" {
		t.Errorf("MessageAttachFileName = %q, want %q", decoded.MessageAttachFileName, "cert.pem")
	}
	if decoded.MessageAttachFileName2 != "key.pem" {
		t.Errorf("MessageAttachFileName2 = %q, want %q", decoded.MessageAttachFileName2, "key.pem")
	}
}
