
___
### обновление и fish
```
sudo apt update

sudo apt install fish
```

```
fish
```
### install xray 
```
bash -c "$(curl -L https://github.com/XTLS/Xray-install/raw/main/install-release.sh)" @ install
```

```
# Генерация UUID для клиента
xray uuid
```

```
# Генерация ключей Reality
xray x25519
```

```
# Генерация Short ID (случайные 16 hex-символов)
openssl rand -hex 8
```

### Конфиг конечного сервера (не ру)

путь 
```
/usr/local/etc/xray/config.json
```

```

{
    "log": {
        "loglevel": "warning"
    },
    "routing": {
        "domainStrategy": "IPIfNonMatch",
        "rules": [
            {
                "type": "field",
                "domain": [
                    "geosite:category-ads-all"
                ],
                "outboundTag": "block"
            }
        ]
    },
    "inbounds": [
        {
            "listen": "0.0.0.0",
            "port": 443,
            "protocol": "vless",
            "settings": {
                "clients": [
                    {
                        "email": "main",
                        "id": "UUID_СЮДА",
                        "flow": "xtls-rprx-vision"
                    }
                ],
                "decryption": "none"
            },
            "streamSettings": {
                "network": "tcp",
                "security": "reality",
                "realitySettings": {
                    "show": false,
                    "dest": "amd.com:443",
                    "xver": 0,
                    "serverNames": [
                        "amd.com",
                        "www.amd.com"
                    ],
                    "privateKey": "PRIVATE_KEY",
                    "shortIds": [
                        "SHORT_ID"
                    ]
                }
            },
            "sniffing": {
                "enabled": true,
                "destOverride": [
                    "http",
                    "tls"
                ]
            }
        }
    ],
    "outbounds": [
        {
            "protocol": "freedom",
            "tag": "direct"
        },
        {
            "protocol": "blackhole",
            "tag": "block"
        }
    ],
    "policy": {
        "levels": {
            "0": {
                "handshake": 3,
                "connIdle": 180
            }
        }
    }
}
```

### Конфиг ру сервера

```
{
    "log": {
        "loglevel": "warning"
    },
    "inbounds": [
        {
            "listen": "0.0.0.0",
            "port": 443,
            "protocol": "vless",
            "settings": {
                "clients": [
                    {
                        "id": "UUID_КЛИЕНТА",
                        "flow": "xtls-rprx-vision"
                    }
                ],
                "decryption": "none"
            },
            "streamSettings": {
                "network": "tcp",
                "security": "reality",
                "realitySettings": {
                    "dest": "amd.com:443",
                    "xver": 0,
                    "serverNames": [
                        "amd.com",
                        "www.amd.com"
                    ],
                    "privateKey": "PRIVATE_KEY_ДЛЯ_ВХОДЯЩИХ",
                    "shortIds": [
                        "SHORT_ID_ДЛЯ_ВХОДЯЩИХ"
                    ]
                }
            },
            "sniffing": {
                "enabled": false
            }
        }
    ],
    "outbounds": [
        {
            "tag": "proxy",
            "protocol": "vless",
            "settings": {
                "vnext": [
                    {
                        "address": "IP_КОНЕЧНОГО",
                        "port": 443,
                        "users": [
                            {
                                "id": "UUID_С_КОНЕЧНОГО",
                                "flow": "xtls-rprx-vision",
                                "encryption": "none"
                            }
                        ]
                    }
                ]
            },
            "streamSettings": {
                "network": "tcp",
                "security": "reality",
                "realitySettings": {
                    "serverName": "amd.com",
                    "fingerprint": "chrome",
                    "shortId": "SHORT_ID_С_КОНЕЧНОГО",
                    "publicKey": "PUBLIC_KEY_С_КОНЕЧНОГО"
                }
            }
        }
    ]
}
```


### Подключение(каскад)
```
vless://UUID_РУССКОГО@IP_РУССКОГО:443?encryption=none&flow=xtls-rprx-vision&security=reality&sni=yandex.ru&fp=chrome&pbk=PUBLIC_KEY_РУССКОГО&sid=SHORT_ID_РУССКОГО&type=tcp#Name
```

### После смены конфигов рестартить
```
systemctl enable xray
systemctl restart xray
systemctl status xray
```

### геобазы

```
bash <(wget -qO- https://github.com/Davoyan/ipregion/raw/main/ipregion.sh | psub)
```


### warp

```

# Добавление GPG ключа
curl -fsSL https://pkg.cloudflareclient.com/pubkey.gpg | sudo gpg --yes --dearmor --output /usr/share/keyrings/cloudflare-warp-archive-keyring.gpg

# Добавление репозитория в список источников
echo "deb [signed-by=/usr/share/keyrings/cloudflare-warp-archive-keyring.gpg] https://pkg.cloudflareclient.com/ $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/cloudflare-client.list

# Обновление кэша пакетов
sudo apt update


sudo apt install cloudflare-warp


# Регистрация клиента
warp-cli registration new

# Установка режима локального прокси (важно для работы с Xray)
warp-cli mode proxy

# Установка порта (тот же, что мы указали в конфиге Xray)
warp-cli proxy port 4000

# Подключение к сети Cloudflare
warp-cli mode proxy
warp-cli connect

# Включение постоянного переподключения при перезагрузке
warp-cli registration set-strategy infinite
```
