package utils

import (
	"fmt"
	"log"
	"net"
	"strconv"
)

// GetInitialPort parses the string port and falls back to 3000 if invalid
func GetInitialPort(portStr string) int {
	port, err := strconv.Atoi(portStr)
	if err != nil || port == 0 {
		log.Printf("Invalid or empty port in config. Defaulting to 3000.")
		return 3000
	}
	return port
}

// FindAvailablePort tries ports from start and increments until an available one is found
func FindAvailablePort(start int) int {
	port := start
	for {
		ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			log.Printf("Port %d is in use. Trying %d...", port, port+1)
			port++
			continue
		}
		_ = ln.Close()
		return port
	}
}
