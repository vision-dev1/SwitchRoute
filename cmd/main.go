// Codes by vision
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/vision-dev1/SwitchRoute/internal/banner"
	"github.com/vision-dev1/SwitchRoute/internal/logger"
	"github.com/vision-dev1/SwitchRoute/internal/proxy"
	"github.com/vision-dev1/SwitchRoute/internal/rotator"
)

func main() {
	banner.Display()

	log, err := logger.New("logs")
	if err != nil {
		banner.PrintError(fmt.Sprintf("Failed to initialize logger: %v", err))
		os.Exit(1)
	}

	proxies, err := loadProxies("config/proxies.txt")
	if err != nil {
		banner.PrintWarning(fmt.Sprintf("Failed to load proxies: %v", err))
		banner.PrintInfo("Starting with empty proxy pool")
		proxies = []string{}
	}

	rot := rotator.New(proxies)
	banner.PrintSuccess(fmt.Sprintf("Loaded %d proxies (%d active)", rot.Count(), rot.ActiveCount()))

	runCLI(rot, log)
}

func loadProxies(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var proxies []string
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		proxies = append(proxies, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return proxies, nil
}

func runCLI(rot *rotator.Rotator, log *logger.Logger) {
	reader := bufio.NewReader(os.Stdin)

	banner.PrintInfo("Type 'help' for available commands")
	fmt.Println()

	for {
		fmt.Print(banner.ColorCyan + "switchroute> " + banner.ColorReset)
		
		input, err := reader.ReadString('\n')
		if err != nil {
			banner.PrintError(fmt.Sprintf("Error reading input: %v", err))
			continue
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		parts := strings.Fields(input)
		command := parts[0]

		switch command {
		case "help":
			showHelp()

		case "list":
			listProxies(rot)

		case "add":
			if len(parts) < 2 {
				banner.PrintError("Usage: add <proxy>")
				continue
			}
			addProxy(rot, parts[1])

		case "remove":
			if len(parts) < 2 {
				banner.PrintError("Usage: remove <proxy>")
				continue
			}
			removeProxy(rot, parts[1])

		case "test":
			if len(parts) < 2 {
				banner.PrintError("Usage: test <url>")
				continue
			}
			testRequest(rot, log, parts[1])

		case "exit", "quit":
			banner.PrintSuccess("Goodbye!")
			return

		default:
			banner.PrintError(fmt.Sprintf("Unknown command: %s", command))
			banner.PrintInfo("Type 'help' for available commands")
		}

		fmt.Println()
	}
}

func showHelp() {
	fmt.Println("\n" + banner.ColorGreen + "Available Commands:" + banner.ColorReset)
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Printf("  %-20s %s\n", "list", "Display current proxy pool with status")
	fmt.Printf("  %-20s %s\n", "add <proxy>", "Add a new proxy to the pool")
	fmt.Printf("  %-20s %s\n", "remove <proxy>", "Remove a proxy from the pool")
	fmt.Printf("  %-20s %s\n", "test <url>", "Send GET request through rotated proxy")
	fmt.Printf("  %-20s %s\n", "help", "Show this help message")
	fmt.Printf("  %-20s %s\n", "exit/quit", "Exit the program")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
}

func listProxies(rot *rotator.Rotator) {
	proxies := rot.List()
	
	if len(proxies) == 0 {
		banner.PrintWarning("No proxies in pool")
		return
	}

	fmt.Println("\n" + banner.ColorGreen + "Proxy Pool Status:" + banner.ColorReset)
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Printf("  %-50s %-10s %s\n", "Proxy", "Status", "Failures")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	for _, p := range proxies {
		status := banner.ColorGreen + "ACTIVE" + banner.ColorReset
		if !p.Active {
			status = banner.ColorRed + "FAILED" + banner.ColorReset
		}
		fmt.Printf("  %-50s %-20s %d\n", p.URL, status, p.Failures)
	}

	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Printf("Total: %d | Active: %d | Failed: %d\n", 
		rot.Count(), rot.ActiveCount(), rot.Count()-rot.ActiveCount())
}

func addProxy(rot *rotator.Rotator, proxyURL string) {
	if err := rot.Add(proxyURL); err != nil {
		banner.PrintError(fmt.Sprintf("Failed to add proxy: %v", err))
		return
	}
	banner.PrintSuccess(fmt.Sprintf("Added proxy: %s", proxyURL))
}

func removeProxy(rot *rotator.Rotator, proxyURL string) {
	if err := rot.Remove(proxyURL); err != nil {
		banner.PrintError(fmt.Sprintf("Failed to remove proxy: %v", err))
		return
	}
	banner.PrintSuccess(fmt.Sprintf("Removed proxy: %s", proxyURL))
}

func testRequest(rot *rotator.Rotator, log *logger.Logger, targetURL string) {
	proxyURL, err := rot.GetNext()
	if err != nil {
		banner.PrintError(fmt.Sprintf("Failed to get proxy: %v", err))
		log.Log("none", targetURL, "FAILED", 0, err)
		return
	}

	banner.PrintActiveIP(proxyURL)
	banner.PrintInfo(fmt.Sprintf("Sending request to: %s", targetURL))

	config := proxy.DefaultConfig()
	p := proxy.New(proxyURL, config)

	resp, err := p.SendRequest(targetURL)
	if err != nil {
		banner.PrintError(fmt.Sprintf("Request failed: %v", err))
		rot.MarkFailed(proxyURL)
		log.Log(proxyURL, targetURL, "FAILED", 0, err)
		return
	}
	defer resp.Body.Close()

	rot.MarkSuccess(proxyURL)
	
	body, err := proxy.GetResponseBody(resp)
	if err != nil {
		banner.PrintWarning(fmt.Sprintf("Failed to read response body: %v", err))
	}

	banner.PrintSuccess(fmt.Sprintf("Request successful! Status: %d", resp.StatusCode))
	
	if len(body) > 500 {
		fmt.Printf("\nResponse (truncated):\n%s\n...\n", body[:500])
	} else {
		fmt.Printf("\nResponse:\n%s\n", body)
	}

	log.Log(proxyURL, targetURL, "SUCCESS", resp.StatusCode, nil)
}
