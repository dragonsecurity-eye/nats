# nats

Shared Go library providing NATS messaging types and connection helpers for the [DragonEye](https://github.com/dragonsecurity-eye) platform.

## Overview

This package defines the data types exchanged between DragonEye components (agents, servers, console) over [NATS](https://nats.io/) messaging, along with connection utilities supporting mTLS and token-based authentication.

## Installation

```bash
go get eye.dragonsecurity.io/nats
```

## Usage

### Connect with mTLS (with optional WebSocket fallback)

```go
conn, err := nats.ConnectWithNATS(servers, clientCert, clientKey, caCert, websocketPort)
```

### Connect with token

```go
conn, err := nats.ConnectWithNATSTOKEN(servers, token)
```

## Package Contents

| File | Description |
|------|-------------|
| `connect.go` | NATS connection helpers (mTLS + WebSocket fallback, token auth) |
| `agent.go` | Agent report payload and all hardware/software inventory types |
| `profiles.go` | Configuration profile and task report types |
| `updater.go` | Agent and server update/release types |
| `deploy.go` | Software deployment action type |
| `winget.go` | Windows package manager types |
| `certificate_request.go` | PKI certificate request/response types |
| `notification.go` | Email notification type |
| `register_status.go` | User registration status constants |
| `system_update_settings.go` | Windows Update configuration status constants |

## License

Apache License 2.0 - see [LICENSE](LICENSE) for details.
