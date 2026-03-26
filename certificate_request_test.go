package nats

import (
	"encoding/json"
	"testing"
)

func TestCertificateRequestMarshalJSON(t *testing.T) {
	req := CertificateRequest{
		FullName:     "John Doe",
		Email:        "john@example.com",
		Username:     "jdoe",
		Organization: "DragonSecurity",
		Country:      "US",
		YearsValid:   2,
		AgentId:      "agent-001",
		DNSName:      "workstation.local",
	}

	data, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("failed to marshal CertificateRequest: %v", err)
	}

	var decoded CertificateRequest
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("failed to unmarshal CertificateRequest: %v", err)
	}

	if decoded.FullName != req.FullName {
		t.Errorf("FullName = %q, want %q", decoded.FullName, req.FullName)
	}
	if decoded.Email != req.Email {
		t.Errorf("Email = %q, want %q", decoded.Email, req.Email)
	}
	if decoded.YearsValid != 2 {
		t.Errorf("YearsValid = %d, want 2", decoded.YearsValid)
	}
	if decoded.DNSName != req.DNSName {
		t.Errorf("DNSName = %q, want %q", decoded.DNSName, req.DNSName)
	}
}

func TestAgentCertificateDataMarshalJSON(t *testing.T) {
	certData := AgentCertificateData{
		CertBytes:       []byte("-----BEGIN CERTIFICATE-----"),
		PrivateKeyBytes: []byte("-----BEGIN PRIVATE KEY-----"),
	}

	data, err := json.Marshal(certData)
	if err != nil {
		t.Fatalf("failed to marshal AgentCertificateData: %v", err)
	}

	var decoded AgentCertificateData
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("failed to unmarshal AgentCertificateData: %v", err)
	}

	if string(decoded.CertBytes) != string(certData.CertBytes) {
		t.Errorf("CertBytes mismatch")
	}
	if string(decoded.PrivateKeyBytes) != string(certData.PrivateKeyBytes) {
		t.Errorf("PrivateKeyBytes mismatch")
	}
}
