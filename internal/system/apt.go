package system

import (
	"fmt"
	"os"
	"os/exec"
)

func RunCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func InstallBasePackages() error {
	fmt.Println("Обновление кэша пакетов...")
	if err := RunCommand("apt", "update"); err != nil {
		return err
	}

	fmt.Println("Установка fish, gnupg, lsb-release...")
	return RunCommand("apt", "install", "-y", "fish", "gnupg", "lsb-release", "curl")
}

func InstallXray() error {
	fmt.Println("Установка Xray...")
	script := "curl -L https://github.com/XTLS/Xray-install/raw/main/install-release.sh | bash -s -- install"
	return RunCommand("bash", "-c", script)
}