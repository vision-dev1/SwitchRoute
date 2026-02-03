// Codes by vision
package rotator

import (
	"fmt"
	"sync"
)

type ProxyStatus struct {
	URL      string
	Active   bool
	Failures int
}

type Rotator struct {
	proxies []ProxyStatus
	current int
	mu      sync.RWMutex
}

func New(proxyURLs []string) *Rotator {
	proxies := make([]ProxyStatus, len(proxyURLs))
	for i, url := range proxyURLs {
		proxies[i] = ProxyStatus{
			URL:      url,
			Active:   true,
			Failures: 0,
		}
	}

	return &Rotator{
		proxies: proxies,
		current: 0,
	}
}

func (r *Rotator) GetNext() (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if len(r.proxies) == 0 {
		return "", fmt.Errorf("no proxies available")
	}

	activeCount := 0
	for _, p := range r.proxies {
		if p.Active {
			activeCount++
		}
	}

	if activeCount == 0 {
		return "", fmt.Errorf("no active proxies available")
	}

	startIdx := r.current
	for {
		if r.proxies[r.current].Active {
			selected := r.proxies[r.current].URL
			r.current = (r.current + 1) % len(r.proxies)
			return selected, nil
		}

		r.current = (r.current + 1) % len(r.proxies)
		
		if r.current == startIdx {
			return "", fmt.Errorf("no active proxies found")
		}
	}
}

func (r *Rotator) MarkFailed(proxyURL string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i := range r.proxies {
		if r.proxies[i].URL == proxyURL {
			r.proxies[i].Failures++
			if r.proxies[i].Failures >= 3 {
				r.proxies[i].Active = false
			}
			break
		}
	}
}

func (r *Rotator) MarkSuccess(proxyURL string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i := range r.proxies {
		if r.proxies[i].URL == proxyURL {
			r.proxies[i].Failures = 0
			r.proxies[i].Active = true
			break
		}
	}
}

func (r *Rotator) Add(proxyURL string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, p := range r.proxies {
		if p.URL == proxyURL {
			return fmt.Errorf("proxy already exists: %s", proxyURL)
		}
	}

	r.proxies = append(r.proxies, ProxyStatus{
		URL:      proxyURL,
		Active:   true,
		Failures: 0,
	})

	return nil
}

func (r *Rotator) Remove(proxyURL string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, p := range r.proxies {
		if p.URL == proxyURL {
			r.proxies = append(r.proxies[:i], r.proxies[i+1:]...)
			
			if r.current >= len(r.proxies) && len(r.proxies) > 0 {
				r.current = 0
			}
			return nil
		}
	}

	return fmt.Errorf("proxy not found: %s", proxyURL)
}

func (r *Rotator) List() []ProxyStatus {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]ProxyStatus, len(r.proxies))
	copy(result, r.proxies)
	return result
}

func (r *Rotator) Count() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.proxies)
}

func (r *Rotator) ActiveCount() int {
	r.mu.RLock()
	defer r.mu.RUnlock()

	count := 0
	for _, p := range r.proxies {
		if p.Active {
			count++
		}
	}
	return count
}
