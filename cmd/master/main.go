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
		log.Fatal("Этот скрипт должен быть запущен от имени root")
	}

	fmt.Println("Запуск установки Xray Master Node")

	appConf, err := config.LoadConfig("config_master.yaml")
	if err != nil {
		log.Fatalf("Ошибка загрузки config.yaml: %v", err)
	}

	if err := system.InstallBasePackages(); err != nil {
		log.Fatalf("Ошибка установки базовых пакетов: %v", err)
	}

	if err := system.InstallXray(); err != nil {
		log.Fatalf("Ошибка установки Xray: %v", err)
	}

	if appConf.Warp.Enabled {
		if err := system.SetupWarp(appConf.Warp.Port); err != nil {
			log.Printf("Предупреждение при настройке WARP: %v", err)
		}
	}

	clientUUID := crypto.GenerateUUID()
	shortID := crypto.GenerateShortID()

	keys, err := crypto.GenerateX25519()
	if err != nil {
		log.Fatalf("Ошибка генерации ключей Reality: %v", err)
	}

	xrayJsonConfig := generator.NewMasterConfig(appConf, clientUUID, keys.Private, shortID)

	configPath := "/usr/local/etc/xray/config.json"
	if err := system.WriteXrayConfig(configPath, xrayJsonConfig); err != nil {
		log.Fatalf("Ошибка записи конфига Xray: %v", err)
	}

	if err := system.RestartXray(); err != nil {
		log.Fatalf("Ошибка запуска сервиса: %v", err)
	}

	serverIP, err := system.GetPublicIP()
	if err != nil {
		serverIP = "YOUR_SERVER_IP"
	}

	printSummary(appConf, serverIP, clientUUID, keys.Public, shortID)
}

func printSummary(appConf *config.AppConfig, ip, uuid, pubKey, sid string) {
	fmt.Println("\n" + strings.Repeat("-", 60))
	fmt.Println("Установка завершена")
	fmt.Println(strings.Repeat("-", 60))

	fmt.Printf("Внешний IP:  %s\n", ip)
	fmt.Printf("UUID:        %s\n", uuid)
	fmt.Printf("Public Key:  %s\n", pubKey)
	fmt.Printf("Short ID:    %s\n", sid)

	fmt.Println(strings.Repeat("-", 60))
	fmt.Println("Ссылка для клиента:")

	vlessLink := generator.GenerateVlessLink(appConf, ip, uuid, pubKey, sid, "sisuliki_master")
	fmt.Printf("\n%s\n\n", vlessLink)
}