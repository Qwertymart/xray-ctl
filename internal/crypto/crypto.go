package crypto

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
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

	path := "/usr/local/bin/xray"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		path = "xray"
	}

	out, err := exec.Command(path, "x25519").Output()
	if err != nil {
		return nil, fmt.Errorf("failed to run xray x25519: %w", err)
	}

	lines := strings.Split(string(out), "\n")
	keys := &KeyPair{}

	for _, line := range lines {
		if strings.Contains(line, "Private key:") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				keys.Private = strings.TrimSpace(parts[1])
			}
		}
		if strings.Contains(line, "Public key:") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				keys.Public = strings.TrimSpace(parts[1])
			}
		}
	}

	if keys.Private == "" || keys.Public == "" {
		return nil, fmt.Errorf("failed to parse xray keys. Output was: %s", string(out))
	}

	return keys, nil
}
