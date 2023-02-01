APP=goback
APP_EXE="./build/$(APP)"

install:
	go get -u ./... && go mod tidy

build:
	mkdir -p ./build && CGO_ENABLED=0 GOOS=linux go build -o ${APP_EXE}

test:
	go test -cover -v ./...