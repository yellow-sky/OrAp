GO_EXEC=~/go/go1.20.6/bin/go
#BUILD_TAGS=-tags swagger_enabled

all: build-amd64 build-arm64 build-armv7l

# Build components
build-swagger:
	swag init -g api/api_service.go

build-web-client:
	export API_ROOT=""; \
	cd ./web/orap-web-client; \
	quasar build

# Build executables
build-armv7l: build-web-client build-swagger
	env GOOS=linux GOARM=7 GOARCH=arm ${GO_EXEC} build -o build/orap_armv7l ${BUILD_TAGS}

build-arm64: build-web-client build-swagger
	env GOOS=linux GOARCH=arm64 ${GO_EXEC} build -o build/orap_aarch64 ${BUILD_TAGS}

build-amd64: build-web-client build-swagger
	env GOOS=linux GOARCH=amd64 ${GO_EXEC} build -o build/orap_amd64 ${BUILD_TAGS}

# Dev tools
dev-server: build-amd64
	export API_PORT=9292; \
	cd ./build; \
	./orap_amd64 configuration generate; \
	./orap_amd64 serve

dev-web-client:
	export API_ROOT="http://127.0.0.1:9292"; \
	cd ./web/orap-web-client; \
	quasar dev

dev-weight:
	# TODO: not worked in makefile
	${GO_EXEC} install github.com/jondot/goweight@latest; \
	~/go/bin/goweight
