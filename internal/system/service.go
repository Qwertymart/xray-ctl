package system

import "fmt"

func RestartXray() error {
	fmt.Println("Перезапуск сервиса Xray...")
	
	if err := RunCommand("systemctl", "daemon-reload"); err != nil {
		return fmt.Errorf("daemon-reload failed: %w", err)
	}

	if err := RunCommand("systemctl", "enable", "xray"); err != nil {
		return fmt.Errorf("enable xray failed: %w", err)
	}

	if err := RunCommand("systemctl", "restart", "xray"); err != nil {
		return fmt.Errorf("restart xray failed: %w", err)
	}

	return nil
}