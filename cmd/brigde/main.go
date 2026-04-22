package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Qwertymart/xray-ctl/internal/config"
	"github.com/Qwertymart/xray-ctl/internal/crypto"
	"github.com/Qwertymart/xray-ctl/internal/generator"
	"github.com/Qwertymart/xray-ctl/internal/system"
)

func main() {
	if os.Geteuid() != 0 {
		log.Fatal("Run as root")
	}

	fmt.Println("Установка Xray Bridge Node")

	cfg, err := config.LoadConfig("config_bridge.yaml")
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	system.InstallBasePackages()
	system.InstallXray()

	localUUID := crypto.GenerateUUID()
	localSID := crypto.GenerateShortID()
	keys, _ := crypto.GenerateX25519()

	xrayCfg := generator.NewBridgeConfig(cfg, localUUID, keys.Private, localSID)

	system.WriteXrayConfig("/usr/local/etc/xray/config.json", xrayCfg)
	system.RestartXray()

	ip, _ := system.GetPublicIP()
	printSummary(cfg, ip, localUUID, keys.Public, localSID)
}

func printSummary(cfg *config.AppConfig, ip, uuid, pub, sid string) {
	fmt.Println("\n" + strings.Repeat("-", 60))
	fmt.Printf("Внешний IP: %s\n", ip)
	fmt.Println(strings.Repeat("-", 60))
	
	link := generator.GenerateVlessLink(cfg, ip, uuid, pub, sid, "sisuliki_bridge")
	fmt.Printf("Ссылка:\n\n%s\n\n", link)
	fmt.Println(strings.Repeat("-", 60))
}