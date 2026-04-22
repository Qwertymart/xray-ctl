package crypto

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os/exec"
	"strings"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	return uuid.New().String()
}

func GenerateShortID() string {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return "0000000000000000"
	}
	return hex.EncodeToString(b)
}

type KeyPair struct {
	Private string
	Public  string
}

func GenerateX25519() (*KeyPair, error) {
	out, err := exec.Command("xray", "x25519").Output()
	if err != nil {
		return nil, fmt.Errorf("failed to run xray x25519: %w (make sure xray is installed)", err)
	}

	lines := strings.Split(string(out), "\n")
	keys := &KeyPair{}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Private key:") {
			keys.Private = strings.TrimSpace(strings.TrimPrefix(line, "Private key:"))
		}
		if strings.HasPrefix(line, "Public key:") {
			keys.Public = strings.TrimSpace(strings.TrimPrefix(line, "Public key:"))
		}
	}

	if keys.Private == "" || keys.Public == "" {
		return nil, fmt.Errorf("failed to parse xray keys from output")
	}

	return keys, nil
}