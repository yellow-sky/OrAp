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

