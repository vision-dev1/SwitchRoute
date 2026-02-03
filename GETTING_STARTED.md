# 🚀 Getting Started with SwitchRoute

**A Complete Beginner's Guide to Running the IP Rotation Tool**

This guide will walk you through everything you need to know to run **SwitchRoute** on Linux, macOS, or Windows. No prior experience required!

---

## 📋 Table of Contents

1. [Installing Go](#1-installing-go)
2. [Cloning the Repository](#2-cloning-the-repository)
3. [Understanding the Project Structure](#3-understanding-the-project-structure)
4. [Configuring Your Proxies](#4-configuring-your-proxies)
5. [Running the Tool](#5-running-the-tool)
6. [Interactive CLI Commands](#6-interactive-cli-commands)
7. [Building a Standalone Binary](#7-building-a-standalone-binary-optional)
8. [Understanding Logs](#8-understanding-logs)
9. [Troubleshooting](#9-troubleshooting)

---

## 1. Installing Go

SwitchRoute is built with Go, so you'll need to install it first.

### 🐧 Linux

**Option A: Using Package Manager (Ubuntu/Debian)**
```bash
sudo apt update
sudo apt install golang-go
```

**Option B: Download Official Binary**
```bash
# Download Go (replace with latest version from https://go.dev/dl/)
wget https://go.dev/dl/go1.21.6.linux-amd64.tar.gz

# Remove old installation (if any)
sudo rm -rf /usr/local/go

# Extract to /usr/local
sudo tar -C /usr/local -xzf go1.21.6.linux-amd64.tar.gz

# Add to PATH (add this to ~/.bashrc or ~/.zshrc)
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

### 🍎 macOS

**Option A: Using Homebrew (Recommended)**
```bash
brew install go
```

**Option B: Download Official Installer**
1. Visit [https://go.dev/dl/](https://go.dev/dl/)
2. Download the `.pkg` file for macOS
3. Run the installer and follow the prompts

### 🪟 Windows

**Option A: Using Chocolatey**
```powershell
choco install golang
```

**Option B: Download Official Installer**
1. Visit [https://go.dev/dl/](https://go.dev/dl/)
2. Download the `.msi` file for Windows
3. Run the installer and follow the prompts
4. The installer will automatically add Go to your PATH

### ✅ Verify Installation

Open a new terminal/command prompt and run:
```bash
go version
```

You should see output like:
```
go version go1.21.6 linux/amd64
```

---

## 2. Cloning the Repository

Now let's download the SwitchRoute project from GitHub.

### Using Git

**If you have Git installed:**
```bash
# Clone the repository
git clone https://github.com/vision-dev1/SwitchRoute.git

# Navigate into the directory
cd SwitchRoute
```

### Without Git

**If you don't have Git:**
1. Visit [https://github.com/vision-dev1/SwitchRoute](https://github.com/vision-dev1/SwitchRoute)
2. Click the green **"Code"** button
3. Select **"Download ZIP"**
4. Extract the ZIP file
5. Open a terminal and navigate to the extracted folder:
   ```bash
   cd path/to/SwitchRoute
   ```

### 📦 Download Dependencies

Once inside the SwitchRoute directory, download the required Go modules:
```bash
go mod download
```

This will download all necessary dependencies for the project.

---

## 3. Understanding the Project Structure

Here's what each folder and file does:

```
SwitchRoute/
├── cmd/
│   └── main.go                 # Main program - this is what you run
├── internal/
│   ├── banner/                 # Handles colored terminal output
│   ├── rotator/                # Core IP rotation logic
│   ├── proxy/                  # Proxy connection handling
│   └── logger/                 # Request logging system
├── config/
│   └── proxies.txt            # YOUR PROXY LIST GOES HERE
├── logs/                       # Request logs are saved here (auto-created)
├── go.mod                      # Go module dependencies
└── README.md                   # Project documentation
```

**Key Points:**
- **`cmd/main.go`** - The entry point of the application
- **`config/proxies.txt`** - Where you add your proxy servers
- **`logs/`** - All request logs are automatically saved here

---

## 4. Configuring Your Proxies

Before running the tool, you need to add your proxy servers.

### 📝 Edit the Proxy Configuration File

Open `config/proxies.txt` in any text editor:

```bash
# Linux/macOS
nano config/proxies.txt

# Or use any text editor you prefer
code config/proxies.txt    # VS Code
vim config/proxies.txt     # Vim
```

### ✍️ Proxy Format

Add your proxies one per line in this format:
```
protocol://host:port
```

**Supported Protocols:**
- `http://` - HTTP proxies
- `https://` - HTTPS proxies
- `socks5://` - SOCKS5 proxies
- `direct` - Direct connection (no proxy)

### 📋 Example Configuration

```txt
# HTTP Proxies
http://proxy1.example.com:8080
http://192.168.1.100:3128

# HTTPS Proxies
https://secure-proxy.example.com:8443

# SOCKS5 Proxies
socks5://socks-proxy.example.com:1080

# Direct connection (useful for testing)
direct

# Lines starting with # are comments and will be ignored
# Empty lines are also ignored
```

> **⚠️ Important:** Replace the example proxies with your actual proxy servers. The tool comes with only `direct` by default, which means no proxy.

### 🔐 Authenticated Proxies

If your proxy requires authentication, use this format:
```
http://username:password@proxy.example.com:8080
```

---

## 5. Running the Tool

Now you're ready to run SwitchRoute!

### 🎯 Start the Application

From the SwitchRoute directory, run:

```bash
go run cmd/main.go
```

### 🎨 What You'll See

When the tool starts, you'll see:

1. **ASCII Banner** - A stylish SwitchRoute logo
2. **Initialization Messages** - Confirmation of loaded proxies
3. **Command Prompt** - `switchroute>` waiting for your input

**Example Output:**
```
╔═══════════════════════════════════════════════════════════╗
║                                                           ║
║   ███████╗██╗    ██╗██╗████████╗ ██████╗██╗  ██╗        ║
║   ██╔════╝██║    ██║██║╚══██╔══╝██╔════╝██║  ██║        ║
║   ███████╗██║ █╗ ██║██║   ██║   ██║     ███████║        ║
║   ╚════██║██║███╗██║██║   ██║   ██║     ██╔══██║        ║
║   ███████║╚███╔███╔╝██║   ██║   ╚██████╗██║  ██║        ║
║   ╚══════╝ ╚══╝╚══╝ ╚═╝   ╚═╝    ╚═════╝╚═╝  ╚═╝        ║
║                                                           ║
║   ██████╗  ██████╗ ██╗   ██╗████████╗███████╗            ║
║   ██╔══██╗██╔═══██╗██║   ██║╚══██╔══╝██╔════╝            ║
║   ██████╔╝██║   ██║██║   ██║   ██║   █████╗              ║
║   ██╔══██╗██║   ██║██║   ██║   ██║   ██╔══╝              ║
║   ██║  ██║╚██████╔╝╚██████╔╝   ██║   ███████╗            ║
║   ╚═╝  ╚═╝ ╚═════╝  ╚═════╝    ╚═╝   ╚══════╝            ║
║                                                           ║
║            Professional IP Rotation Tool                 ║
║                 CLI Edition v1.0                         ║
║                                                           ║
╚═══════════════════════════════════════════════════════════╝

✓ Loaded 1 proxies (1 active)
→ Type 'help' for available commands

switchroute>
```

---

## 6. Interactive CLI Commands

Once the tool is running, you can use these commands:

### 📋 `help` - Show All Commands

```bash
switchroute> help
```

**Output:**
```
Available Commands:
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  list                 Display current proxy pool with status
  add <proxy>          Add a new proxy to the pool
  remove <proxy>       Remove a proxy from the pool
  test <url>           Send GET request through rotated proxy
  help                 Show this help message
  exit/quit            Exit the program
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
```

---

### 📊 `list` - View All Proxies

Shows all proxies in your pool with their current status.

```bash
switchroute> list
```

**Example Output:**
```
Proxy Pool Status:
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  Proxy                                              Status      Failures
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  http://proxy1.example.com:8080                     ACTIVE      0
  http://proxy2.example.com:3128                     ACTIVE      0
  direct                                             ACTIVE      0
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Total: 3 | Active: 3 | Failed: 0
```

**Status Indicators:**
- 🟢 **ACTIVE** - Proxy is working and available
- 🔴 **FAILED** - Proxy has failed and is temporarily disabled

---

### ➕ `add <proxy>` - Add a New Proxy

Add a proxy to the pool while the tool is running.

**Syntax:**
```bash
add protocol://host:port
```

**Examples:**
```bash
switchroute> add http://proxy3.example.com:8080
✓ Added proxy: http://proxy3.example.com:8080

switchroute> add socks5://192.168.1.50:1080
✓ Added proxy: socks5://192.168.1.50:1080
```

---

### ➖ `remove <proxy>` - Remove a Proxy

Remove a proxy from the pool.

**Syntax:**
```bash
remove protocol://host:port
```

**Example:**
```bash
switchroute> remove http://proxy3.example.com:8080
✓ Removed proxy: http://proxy3.example.com:8080
```

---

### 🧪 `test <url>` - Test a Request

Send a GET request to a URL using the next proxy in rotation.

**Syntax:**
```bash
test <url>
```

**Examples:**

**Test with IP checker:**
```bash
switchroute> test http://httpbin.org/ip
🔄 Active IP: http://proxy1.example.com:8080
→ Sending request to: http://httpbin.org/ip
✓ Request successful! Status: 200

Response:
{
  "origin": "203.0.113.45"
}
```

**Test with any website:**
```bash
switchroute> test https://api.github.com
🔄 Active IP: http://proxy2.example.com:3128
→ Sending request to: https://api.github.com
✓ Request successful! Status: 200

Response (truncated):
{"current_user_url":"https://api.github.com/user"...
```

**What Happens:**
1. The tool selects the next proxy in rotation
2. Sends a GET request through that proxy
3. Shows the response status and body
4. Logs the request to the logs folder
5. If the request fails, marks the proxy as failed and tries the next one

---

### 🚪 `exit` or `quit` - Exit the Program

Stop the tool and return to your terminal.

```bash
switchroute> exit
✓ Goodbye!
```

---

## 7. Building a Standalone Binary (Optional)

Instead of running with `go run`, you can build a standalone executable.

### 🔨 Build the Binary

**Linux/macOS:**
```bash
go build -o switchroute cmd/main.go
```

**Windows:**
```bash
go build -o switchroute.exe cmd/main.go
```

### ▶️ Run the Binary

**Linux/macOS:**
```bash
./switchroute
```

**Windows:**
```bash
switchroute.exe
```

### 📦 Build for Different Platforms

You can build for other operating systems from any platform:

**Build for Linux (from any OS):**
```bash
GOOS=linux GOARCH=amd64 go build -o switchroute-linux cmd/main.go
```

**Build for macOS (from any OS):**
```bash
GOOS=darwin GOARCH=amd64 go build -o switchroute-macos cmd/main.go
```

**Build for Windows (from any OS):**
```bash
GOOS=windows GOARCH=amd64 go build -o switchroute.exe cmd/main.go
```

**Build for ARM (Raspberry Pi, etc.):**
```bash
GOOS=linux GOARCH=arm64 go build -o switchroute-arm cmd/main.go
```

### 📤 Distribute Your Binary

The compiled binary is standalone and can be distributed to others without requiring Go installation. Just share the binary file!

---

## 8. Understanding Logs

SwitchRoute automatically logs all requests for auditing and debugging.

### 📁 Log Location

Logs are stored in the `logs/` directory:
```
logs/
└── requests_2026-02-01.json
```

Each day gets its own log file named `requests_YYYY-MM-DD.json`.

### 📄 Log Format

Logs are in JSON format, one entry per line:

```json
{"timestamp":"2026-02-01T00:27:33+05:45","proxy":"http://proxy1.example.com:8080","url":"http://httpbin.org/ip","status":"SUCCESS","response_code":200}
{"timestamp":"2026-02-01T00:28:15+05:45","proxy":"http://proxy2.example.com:3128","url":"https://api.github.com","status":"SUCCESS","response_code":200}
{"timestamp":"2026-02-01T00:29:02+05:45","proxy":"http://bad-proxy.example.com:8080","url":"http://httpbin.org/ip","status":"FAILED","response_code":0,"error":"dial tcp: connection refused"}
```

### 🔍 Log Fields

| Field | Description |
|-------|-------------|
| `timestamp` | When the request was made (ISO 8601 format) |
| `proxy` | Which proxy was used |
| `url` | The target URL |
| `status` | `SUCCESS` or `FAILED` |
| `response_code` | HTTP status code (200, 404, etc.) |
| `error` | Error message (only present if failed) |

### 📊 Analyzing Logs

**View today's logs:**
```bash
cat logs/requests_$(date +%Y-%m-%d).json
```

**Count successful requests:**
```bash
grep "SUCCESS" logs/requests_*.json | wc -l
```

**Find failed requests:**
```bash
grep "FAILED" logs/requests_*.json
```

**Pretty-print JSON logs:**
```bash
cat logs/requests_*.json | jq '.'
```

---

## 9. Troubleshooting

### ❌ "go: command not found"

**Problem:** Go is not installed or not in your PATH.

**Solution:**
1. Install Go following [Step 1](#1-installing-go)
2. Verify with `go version`
3. If installed but not found, add Go to your PATH:
   ```bash
   export PATH=$PATH:/usr/local/go/bin
   ```

---

### ❌ "Failed to load proxies"

**Problem:** The `config/proxies.txt` file is missing or has errors.

**Solution:**
1. Check if the file exists:
   ```bash
   ls -l config/proxies.txt
   ```
2. Verify the format (one proxy per line)
3. Make sure there are no syntax errors
4. The tool will start with an empty pool if the file is missing

---

### ❌ "No active proxies available"

**Problem:** All proxies have failed.

**Solution:**
1. Check your proxy configuration
2. Verify proxies are reachable:
   ```bash
   curl -x http://proxy.example.com:8080 http://httpbin.org/ip
   ```
3. Add `direct` to your proxy list as a fallback
4. Check the logs for error details

---

### ❌ "Request failed: connection refused"

**Problem:** The proxy server is not responding.

**Solution:**
1. Verify the proxy is online
2. Check firewall settings
3. Ensure the proxy format is correct
4. Try a different proxy

---

### ❌ "Permission denied" when creating logs

**Problem:** The tool can't create the logs directory.

**Solution:**
```bash
# Create the logs directory manually
mkdir -p logs

# Or run with appropriate permissions
sudo go run cmd/main.go
```

---

### 🐛 Still Having Issues?

1. **Check the logs** in `logs/` for detailed error messages
2. **Run with verbose output** (if implemented)
3. **Open an issue** on GitHub: [https://github.com/vision-dev1/SwitchRoute/issues](https://github.com/vision-dev1/SwitchRoute/issues)

---

## 🎉 You're All Set!

You now know how to:
- ✅ Install Go on any platform
- ✅ Clone and set up SwitchRoute
- ✅ Configure your proxy list
- ✅ Run the tool and use all commands
- ✅ Build standalone binaries
- ✅ Understand and analyze logs
- ✅ Troubleshoot common issues

### 🚀 Next Steps

1. Add your real proxies to `config/proxies.txt`
2. Test with `http://httpbin.org/ip` to verify rotation
3. Integrate SwitchRoute into your projects
4. Build a binary for easy distribution

### 📚 Additional Resources

- **Main README:** [README.md](README.md)
- **GitHub Repository:** [https://github.com/vision-dev1/SwitchRoute](https://github.com/vision-dev1/SwitchRoute)
- **Go Documentation:** [https://go.dev/doc/](https://go.dev/doc/)

---

**Happy Routing! 🔄**

*Built with ❤️ by vision-dev1*
