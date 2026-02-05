# 🔄 SwitchRoute

**Professional IP Rotation Tool - CLI Edition**

A powerful, modular IP rotation tool built in Go for managing proxy pools, rotating IPs, and handling HTTP/HTTPS/SOCKS5 requests with automatic failover.

![GitHub](https://img.shields.io/badge/GitHub-vision--dev1-blue?style=flat-square&logo=github)
![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go)
![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)

---

## ✨ Features

- 🎯 **IP Rotation Engine** - Round-robin rotation with automatic failover
- 🔌 **Multi-Protocol Support** - HTTP, HTTPS, and SOCKS5 proxies
- 🛡️ **Automatic Failover** - Marks and skips failed proxies automatically
- 📊 **JSON Logging** - Detailed request logs with timestamps and status
- 🎨 **Colored CLI** - Beautiful terminal output with ANSI colors
- ⚡ **Concurrent Requests** - Efficient goroutine-based request handling
- 🔧 **Dynamic Management** - Add/remove proxies while running
- 🔄 **Retry Mechanism** - Configurable retry attempts and timeouts
- 📝 **Request Tracking** - Complete audit trail of all requests

---

## 🚀 Installation

### Prerequisites
- Go 1.21 or higher

### Clone and Build
```bash
git clone https://github.com/vision-dev1/SwitchRoute.git
cd SwitchRoute
go mod download
```

---

## 📖 Usage

### Start the Tool
```bash
go run cmd/main.go
```

### Available Commands

| Command | Description | Example |
|---------|-------------|---------|
| `list` | Display current proxy pool with status | `list` |
| `add <proxy>` | Add a new proxy to the pool | `add http://proxy.example.com:8080` |
| `remove <proxy>` | Remove a proxy from the pool | `remove http://proxy.example.com:8080` |
| `test <url>` | Send GET request through rotated proxy | `test http://httpbin.org/ip` |
| `help` | Show available commands | `help` |
| `exit` | Exit the program | `exit` |

### Example Session
```
switchroute> list
Proxy Pool Status:
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  Proxy                                              Status      Failures
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  direct                                             ACTIVE      0
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Total: 1 | Active: 1 | Failed: 0

switchroute> add http://proxy.example.com:8080
✓ Added proxy: http://proxy.example.com:8080

switchroute> test http://httpbin.org/ip
🔄 Active IP: http://proxy.example.com:8080
→ Sending request to: http://httpbin.org/ip
✓ Request successful! Status: 200
```

---

## ⚙️ Configuration

### Proxy Configuration File
Edit `config/proxies.txt` to add your proxies:

```txt
# HTTP Proxies
http://proxy1.example.com:8080
http://proxy2.example.com:3128

# HTTPS Proxies
https://proxy3.example.com:8443

# SOCKS5 Proxies
socks5://proxy4.example.com:1080

# Direct connection (no proxy)
direct
```

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `SWITCHROUTE_TIMEOUT` | Request timeout duration | `10s` |
| `SWITCHROUTE_RETRIES` | Number of retry attempts | `3` |
| `SWITCHROUTE_STRATEGY` | Rotation strategy | `round-robin` |

**Example:**
```bash
export SWITCHROUTE_TIMEOUT=15s
export SWITCHROUTE_RETRIES=5
go run cmd/main.go
```

---

## 📁 Project Structure

```
SwitchRoute/
├── cmd/
│   └── main.go                 # Entry point with CLI loop
├── internal/
│   ├── banner/
│   │   └── banner.go          # ASCII banner and colored output
│   ├── rotator/
│   │   └── rotator.go         # IP rotation logic
│   ├── proxy/
│   │   └── proxy.go           # Proxy handling and requests
│   └── logger/
│       └── logger.go          # JSON logging module
├── config/
│   └── proxies.txt            # Proxy configuration
├── logs/                       # Request logs (auto-created)
├── go.mod                      # Go module file
└── README.md                   # Documentation
```

---

## 📊 Logging

All requests are logged to `logs/requests_YYYY-MM-DD.json` in JSON format:

```json
{
  "timestamp": "2026-02-01T00:00:00+05:45",
  "proxy": "http://proxy.example.com:8080",
  "url": "http://httpbin.org/ip",
  "status": "SUCCESS",
  "response_code": 200
}
```

---

## 🛠️ Development

### Run Tests
```bash
go test ./...
```

### Build Binary
```bash
go build -o switchroute cmd/main.go
./switchroute
```

---

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

---

## 📄 License

This project is licensed under the MIT License.

---

## 👨‍💻 Author

**vision-dev1**
- GitHub: [@vision-dev1](https://github.com/vision-dev1)

---

## ⭐ Show Your Support

Give a ⭐️ if this project helped you!

---

**Built with ❤️ using Go**
