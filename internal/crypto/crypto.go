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
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) < 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		if strings.Contains(key, "PrivateKey") {
			keys.Private = value
		} else if strings.Contains(key, "PublicKey") {
			keys.Public = value
		}
	}

	if keys.Private == "" || keys.Public == "" {
		return nil, fmt.Errorf("failed to parse xray keys. Output was:\n%s", string(out))
	}

	return keys, nil
}
