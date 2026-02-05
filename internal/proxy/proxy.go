// Codes by vision
package proxy

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"golang.org/x/net/proxy"
)

type Config struct {
	Timeout time.Duration
	Retries int
}

type Proxy struct {
	URL    string
	config Config
}

func New(proxyURL string, config Config) *Proxy {
	return &Proxy{
		URL:    proxyURL,
		config: config,
	}
}

func DefaultConfig() Config {
	timeout := 10 * time.Second
	retries := 3

	if timeoutEnv := os.Getenv("SWITCHROUTE_TIMEOUT"); timeoutEnv != "" {
		if t, err := time.ParseDuration(timeoutEnv); err == nil {
			timeout = t
		}
	}

	if retriesEnv := os.Getenv("SWITCHROUTE_RETRIES"); retriesEnv != "" {
		if r, err := strconv.Atoi(retriesEnv); err == nil {
			retries = r
		}
	}

	return Config{
		Timeout: timeout,
		Retries: retries,
	}
}

func (p *Proxy) SendRequest(targetURL string) (*http.Response, error) {
	var lastErr error

	for attempt := 0; attempt < p.config.Retries; attempt++ {
		resp, err := p.doRequest(targetURL)
		if err == nil {
			return resp, nil
		}
		lastErr = err
		time.Sleep(time.Second * time.Duration(attempt+1))
	}

	return nil, fmt.Errorf("all %d attempts failed: %w", p.config.Retries, lastErr)
}

func (p *Proxy) doRequest(targetURL string) (*http.Response, error) {
	client, err := p.createClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", "SwitchRoute/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	return resp, nil
}

func (p *Proxy) createClient() (*http.Client, error) {
	if p.URL == "" || p.URL == "direct" {
		return &http.Client{
			Timeout: p.config.Timeout,
		}, nil
	}

	proxyURL, err := url.Parse(p.URL)
	if err != nil {
		return nil, fmt.Errorf("invalid proxy URL: %w", err)
	}

	switch proxyURL.Scheme {
	case "http", "https":
		return &http.Client{
			Timeout: p.config.Timeout,
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
			},
		}, nil

	case "socks5":
		dialer, err := proxy.SOCKS5("tcp", proxyURL.Host, nil, proxy.Direct)
		if err != nil {
			return nil, fmt.Errorf("failed to create SOCKS5 dialer: %w", err)
		}

		return &http.Client{
			Timeout: p.config.Timeout,
			Transport: &http.Transport{
				Dial: dialer.Dial,
			},
		}, nil

	default:
		return nil, fmt.Errorf("unsupported proxy scheme: %s", proxyURL.Scheme)
	}
}

func GetResponseBody(resp *http.Response) (string, error) {
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}
	return string(body), nil
}
