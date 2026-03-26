# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go shared library (`eye.dragonsecurity.io/nats`) that provides NATS messaging types and connection helpers for the DragonEye/OpenUEM platform. It is imported by other services (agents, servers, consoles) — it is not a standalone application.

## Build Commands

```bash
go build ./...     # Compile the package
go vet ./...       # Static analysis
```

There are no tests or linter configurations in this repo currently.

## Architecture

This is a flat, single-package Go library (`package nats`) with no internal subpackages.

**Key files by concern:**

- `connect.go` — NATS connection helpers (`ConnectWithNATS` for mTLS with WebSocket fallback, `ConnectWithNATSTOKEN` for token auth)
- `agent.go` — Core data model: `AgentReport` (the main telemetry payload agents send) and all its nested types (Computer, OperatingSystem, NetworkAdapter, LogicalDisk, etc.), plus `Config`, `Release`, `AgentSetting`, `VNCConnection`, `RustDesk` types
- `profiles.go` — Configuration profile types (`ProfileConfig`, `ProfileReport`, `TaskReport`); imports `eye.dragonsecurity.io/wingetcfg`
- `updater.go` — Update/release types for both agent and server updates (`OpenUEMUpdateRequest`, `OpenUEMRelease`, `OpenUEMServerRelease`)
- `deploy.go` — Software deployment action type
- `winget.go` — Windows package manager (winget) types
- `certificate_request.go` — PKI certificate request/response types
- `notification.go` — Email notification type
- `register_status.go` — User registration status constants
- `system_update_settings.go` — Windows Update configuration status constants

## Key Patterns

- All struct fields use `json:"..."` tags with `omitempty`; some in `profiles.go` use `yaml:"..."` tags instead
- Constants follow a `DOMAIN.status_name` pattern (e.g., `users.completed`, `systemupdate.disabled`)
- Status enums provide a `*PossibleStatus() []string` helper function listing valid values
- The NATS connection logic tries direct TLS first, then falls back to WebSocket (`wss://`) if a WebSocket port is provided