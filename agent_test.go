package nats

import (
	"encoding/json"
	"testing"
	"time"
)

func TestAgentReportMarshalJSON(t *testing.T) {
	now := time.Now().Truncate(time.Second)
	report := AgentReport{
		AgentID:  "agent-001",
		OS:       "windows",
		Hostname: "workstation-1",
		IP:       "192.168.1.10",
		Computer: Computer{
			Manufacturer:   "Dell",
			Model:          "OptiPlex 7090",
			Serial:         "ABC123",
			Processor:      "Intel i7",
			ProcessorArch:  "x86_64",
			ProcessorCores: 8,
			Memory:         16384,
		},
		ExecutionTime: now,
	}

	data, err := json.Marshal(report)
	if err != nil {
		t.Fatalf("failed to marshal AgentReport: %v", err)
	}

	var decoded AgentReport
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("failed to unmarshal AgentReport: %v", err)
	}

	if decoded.AgentID != report.AgentID {
		t.Errorf("AgentID = %q, want %q", decoded.AgentID, report.AgentID)
	}
	if decoded.Hostname != report.Hostname {
		t.Errorf("Hostname = %q, want %q", decoded.Hostname, report.Hostname)
	}
	if decoded.Computer.Manufacturer != report.Computer.Manufacturer {
		t.Errorf("Computer.Manufacturer = %q, want %q", decoded.Computer.Manufacturer, report.Computer.Manufacturer)
	}
	if decoded.Computer.ProcessorCores != report.Computer.ProcessorCores {
		t.Errorf("Computer.ProcessorCores = %d, want %d", decoded.Computer.ProcessorCores, report.Computer.ProcessorCores)
	}
	if decoded.Computer.Memory != report.Computer.Memory {
		t.Errorf("Computer.Memory = %d, want %d", decoded.Computer.Memory, report.Computer.Memory)
	}
}

func TestAgentReportOmitEmpty(t *testing.T) {
	report := AgentReport{
		AgentID: "agent-002",
	}

	data, err := json.Marshal(report)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatalf("failed to unmarshal to map: %v", err)
	}

	if _, ok := raw["hostname"]; ok {
		t.Error("empty hostname should be omitted")
	}
	if _, ok := raw["apps"]; ok {
		t.Error("nil apps slice should be omitted")
	}
	if _, ok := raw["id"]; !ok {
		t.Error("non-empty AgentID should be present")
	}
}

func TestAgentReportWithNestedTypes(t *testing.T) {
	report := AgentReport{
		AgentID: "agent-003",
		LogicalDisks: []LogicalDisk{
			{Label: "C:", Usage: 75, Filesystem: "NTFS", SizeInUnits: "500GB"},
			{Label: "D:", Usage: 30, Filesystem: "NTFS", SizeInUnits: "1TB"},
		},
		NetworkAdapters: []NetworkAdapter{
			{Name: "Ethernet", MACAddress: "AA:BB:CC:DD:EE:FF", Addresses: "192.168.1.10", DHCPEnabled: true},
		},
		Applications: []Application{
			{Name: "Firefox", Version: "125.0", Publisher: "Mozilla"},
		},
	}

	data, err := json.Marshal(report)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var decoded AgentReport
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if len(decoded.LogicalDisks) != 2 {
		t.Fatalf("LogicalDisks len = %d, want 2", len(decoded.LogicalDisks))
	}
	if decoded.LogicalDisks[0].Label != "C:" {
		t.Errorf("LogicalDisks[0].Label = %q, want %q", decoded.LogicalDisks[0].Label, "C:")
	}
	if decoded.LogicalDisks[0].Usage != 75 {
		t.Errorf("LogicalDisks[0].Usage = %d, want 75", decoded.LogicalDisks[0].Usage)
	}
	if len(decoded.NetworkAdapters) != 1 {
		t.Fatalf("NetworkAdapters len = %d, want 1", len(decoded.NetworkAdapters))
	}
	if !decoded.NetworkAdapters[0].DHCPEnabled {
		t.Error("NetworkAdapters[0].DHCPEnabled should be true")
	}
	if len(decoded.Applications) != 1 {
		t.Fatalf("Applications len = %d, want 1", len(decoded.Applications))
	}
	if decoded.Applications[0].Name != "Firefox" {
		t.Errorf("Applications[0].Name = %q, want %q", decoded.Applications[0].Name, "Firefox")
	}
}

func TestConfigMarshalJSON(t *testing.T) {
	cfg := Config{
		Ok:                       true,
		AgentFrequency:           300,
		WinGetFrequency:          3600,
		SFTPDisabled:             false,
		RemoteAssistanceDisabled: true,
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		t.Fatalf("failed to marshal Config: %v", err)
	}

	var decoded Config
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("failed to unmarshal Config: %v", err)
	}

	if decoded.AgentFrequency != 300 {
		t.Errorf("AgentFrequency = %d, want 300", decoded.AgentFrequency)
	}
	if !decoded.RemoteAssistanceDisabled {
		t.Error("RemoteAssistanceDisabled should be true")
	}
}

func TestReleaseMarshalJSON(t *testing.T) {
	releaseDate := time.Date(2025, 6, 15, 0, 0, 0, 0, time.UTC)
	release := Release{
		Version:     "1.2.0",
		Channel:     "stable",
		Summary:     "Bug fixes",
		IsCritical:  true,
		ReleaseDate: releaseDate,
		Arch:        "amd64",
		Os:          "windows",
		FileURL:     "https://example.com/agent-1.2.0.exe",
		Checksum:    "abc123",
	}

	data, err := json.Marshal(release)
	if err != nil {
		t.Fatalf("failed to marshal Release: %v", err)
	}

	var decoded Release
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("failed to unmarshal Release: %v", err)
	}

	if decoded.Version != "1.2.0" {
		t.Errorf("Version = %q, want %q", decoded.Version, "1.2.0")
	}
	if !decoded.IsCritical {
		t.Error("IsCritical should be true")
	}
	if !decoded.ReleaseDate.Equal(releaseDate) {
		t.Errorf("ReleaseDate = %v, want %v", decoded.ReleaseDate, releaseDate)
	}
}

func TestVNCConnectionMarshalJSON(t *testing.T) {
	vnc := VNCConnection{
		PIN:        "123456",
		NotifyUser: true,
	}

	data, err := json.Marshal(vnc)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var decoded VNCConnection
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if decoded.PIN != "123456" {
		t.Errorf("PIN = %q, want %q", decoded.PIN, "123456")
	}
	if !decoded.NotifyUser {
		t.Error("NotifyUser should be true")
	}
}
