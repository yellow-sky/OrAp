all: build-amd64 build-arm64 build-armv7l

build-swagger:
	 swag init -g api/api_service.go

build-web-client:
	export API_ROOT=""; \
	cd ./web/orap-web-client; \
	quasar build

build-armv7l: build-web-client build-swagger
	env GOOS=linux GOARM=7 GOARCH=arm go build -o build/orap_armv7l

build-arm64: build-web-client build-swagger
	env GOOS=linux GOARCH=arm64 go build -o build/orap_aarch64

build-amd64: build-web-client build-swagger
	env GOOS=linux GOARCH=amd64 go build -o build/orap_amd64

