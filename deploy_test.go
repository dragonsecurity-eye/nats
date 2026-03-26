package nats

import (
	"encoding/json"
	"testing"
	"time"
)

func TestDeployActionMarshalJSON(t *testing.T) {
	when := time.Date(2025, 6, 15, 10, 30, 0, 0, time.UTC)
	action := DeployAction{
		AgentId:        "agent-001",
		Action:         "install",
		When:           when,
		PackageId:      "pkg-100",
		PackageName:    "Firefox",
		PackageVersion: "125.0",
		Repository:     "winget",
		Failed:         false,
	}

	data, err := json.Marshal(action)
	if err != nil {
		t.Fatalf("failed to marshal DeployAction: %v", err)
	}

	var decoded DeployAction
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("failed to unmarshal DeployAction: %v", err)
	}

	if decoded.AgentId != "agent-001" {
		t.Errorf("AgentId = %q, want %q", decoded.AgentId, "agent-001")
	}
	if decoded.Action != "install" {
		t.Errorf("Action = %q, want %q", decoded.Action, "install")
	}
	if decoded.PackageName != "Firefox" {
		t.Errorf("PackageName = %q, want %q", decoded.PackageName, "Firefox")
	}
	if !decoded.When.Equal(when) {
		t.Errorf("When = %v, want %v", decoded.When, when)
	}
}

func TestDeployActionFailedState(t *testing.T) {
	action := DeployAction{
		AgentId: "agent-002",
		Action:  "uninstall",
		Failed:  true,
		Info:    "package not found",
	}

	data, err := json.Marshal(action)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var decoded DeployAction
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if !decoded.Failed {
		t.Error("Failed should be true")
	}
	if decoded.Info != "package not found" {
		t.Errorf("Info = %q, want %q", decoded.Info, "package not found")
	}
}
