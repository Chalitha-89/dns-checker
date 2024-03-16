package main

import (
//    "fmt"
    "log"
    "net"
    "time"
)

func resolveDomain(domain string) {
    attempts := 3
    delay := 2 * time.Second // Delay between attempts

    for i := 0; i < attempts; i++ {
        ips, err := net.LookupIP(domain)
        if err != nil {
            log.Printf("Attempt %d: Failed to resolve %s: %v\n", i+1, domain, err)
            time.Sleep(delay)
            continue
        }
        log.Printf("IP addresses for %s: %v\n", domain, ips)
        return // Success, exit the loop
    }
    log.Printf("Failed to resolve %s after %d attempts\n", domain, attempts)
}

func main() {
    log.SetFlags(log.LstdFlags | log.Lmicroseconds) // Enhanced logging format

    domains := []string{"www.googleapis.com", "api.twilio.com"}
    for _, domain := range domains {
        resolveDomain(domain)
    }
}

