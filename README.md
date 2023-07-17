# OrAp
Simple service for organize access point with OrangePi board

## Build for Armbian & OrangePi Zero LTS

```shell
env GOOS=linux GOARM=7 GOARCH=arm go build -o build/orap_armv7l
```

## Build for Armbian  & OrangePi Zero 2
```shell
env GOOS=linux GOARCH=arm64 go build -o build/orap_aarch64
```


## Update swagger docs  
```shell
#swag init --parseDependency --parseDepth 2 -g api/api_service.go
swag init -g api/api_service.go
```

## TODO
 - Add speedtest backend (https://github.com/showwin/speedtest-go)
 - Add speedtest frontend (???)
 - Add AT commands cli
 - Add imei changer for well known modems
 - Add ttl replacer (???)
 - Add ssl (???) (https://github.com/FiloSottile/mkcert)
 - Add system dependency checker (nm, mm, os, packages)
