package nats

import (
	"encoding/json"
	"testing"
)

func TestWingetPackageMarshalJSON(t *testing.T) {
	pkg := WingetPackage{
		ID:   "Mozilla.Firefox",
		Name: "Firefox",
	}

	data, err := json.Marshal(pkg)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var decoded WingetPackage
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if decoded.ID != "Mozilla.Firefox" {
		t.Errorf("ID = %q, want %q", decoded.ID, "Mozilla.Firefox")
	}
	if decoded.Name != "Firefox" {
		t.Errorf("Name = %q, want %q", decoded.Name, "Firefox")
	}
}

func TestSoftwarePackageMarshalJSON(t *testing.T) {
	pkg := SoftwarePackage{
		ID:     "Mozilla.Firefox",
		Name:   "Firefox",
		Source: "winget",
	}

	data, err := json.Marshal(pkg)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var decoded SoftwarePackage
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if decoded.Source != "winget" {
		t.Errorf("Source = %q, want %q", decoded.Source, "winget")
	}
}

func TestWingetCfgDeployMarshalJSON(t *testing.T) {
	deploy := WingetCfgDeploy{
		PackageID: "Mozilla.Firefox",
		Installed: true,
	}

	data, err := json.Marshal(deploy)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var decoded WingetCfgDeploy
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if decoded.PackageID != "Mozilla.Firefox" {
		t.Errorf("PackageID = %q, want %q", decoded.PackageID, "Mozilla.Firefox")
	}
	if !decoded.Installed {
		t.Error("Installed should be true")
	}
}
