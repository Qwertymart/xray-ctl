package system

import "fmt"

func SetupWarp(port int) error {
	fmt.Println("Настройка Cloudflare WARP...")
	keyCmd := "curl -fsSL https://pkg.cloudflareclient.com/pubkey.gpg | gpg --yes --dearmor -o /usr/share/keyrings/cloudflare-warp-archive-keyring.gpg"
	if err := RunCommand("bash", "-c", keyCmd); err != nil {
		return err
	}

	repoCmd := `echo "deb [signed-by=/usr/share/keyrings/cloudflare-warp-archive-keyring.gpg] https://pkg.cloudflareclient.com/ $(lsb_release -cs) main" | tee /etc/apt/sources.list.d/cloudflare-client.list`
	if err := RunCommand("bash", "-c", repoCmd); err != nil {
		return err
	}

	if err := RunCommand("apt", "update"); err != nil {
		return err
	}
	if err := RunCommand("apt", "install", "-y", "cloudflare-warp"); err != nil {
		return err
	}

	execs := [][]string{
        {"bash", "-c", "yes | warp-cli registration new"},
        {"warp-cli", "mode", "proxy"},
        {"warp-cli", "proxy", "port", fmt.Sprintf("%d", port)},
        {"warp-cli", "connect"},
    }

    for _, e := range execs {
        var err error
        if e[0] == "bash" {
            err = RunCommand("bash", e[1:]...)
        } else {
            err = RunCommand(e[0], e[1:]...)
        }
        if err != nil {
            fmt.Printf("Предупреждение при выполнении %v: %v\n", e, err)
        }
    }
    return nil
}