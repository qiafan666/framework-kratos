GOPATH:=$(shell go env GOPATH)

# 从命令行输入的路径变量
PROTO_PATH := $(shell read -p "Enter proto file path: " path; echo $$path)

# 提取目录路径
DIR_PATH := $(dir $(PROTO_PATH))

# 删除第一个路径和最后一个文件名
PATH_WITHOUT_PROJECT = $(subst $(word 1,$(subst /, ,$(DIR_PATH)))/,,$(DIR_PATH))

# 使用basename函数获取文件名
FILE_NAME := $(basename $(notdir $(PROTO_PATH)))
UPPER_FILE_NAME := $(shell echo $(FILE_NAME) | awk '{print toupper(substr($$0, 1, 1)) substr($$0, 2)}')

name:
	@echo $(PATH_WITHOUT_PROJECT)
	@echo $(DIR_PATH)
	@echo $(FILE_NAME)
	@echo $(UPPER_FILE_NAME)

APP_PATH := "app"

proto: add grpc http swagger errors protoc server

app: project

.PHONY: init
# init env
init:
	echo $(GOPATH)
	cd $(GOPATH)/bin; \
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest

.PHONY: add
# add proto
add:
	if [ -d $(PROTO_PATH) ]; then exit 0; fi; \
	kratos proto add $(PROTO_PATH)
	mkdir -p $(DIR_PATH)gen

.PHONY: project
# 假设您要在命令行中指定路径
project:
	#创建项目
	if [ ! -d $(APP_PATH) ]; then mkdir $(APP_PATH); fi; \
	cd $(APP_PATH); \
	if [ -d $(PATH_WITHOUT_PROJECT) ]; then exit 0; fi; \
	kratos new $(PATH_WITHOUT_PROJECT) & pid=$$!; \
	echo "Waiting for 10 seconds..."; \
	sleep 10; \
	wait $$pid; \
	cd $(PATH_WITHOUT_PROJECT); \
	rm -rf api third_party .github LiCENSE README.md go.mod go.sum \
	   internal/biz/greeter.go internal/data/greeter.go internal/service/greeter.go; \
	touch README.md

.PHONY: grpc
# generate grpc code
grpc:
	protoc --proto_path=$(DIR_PATH) \
		   --proto_path=./third_party \
		   --go_out=paths=source_relative:$(DIR_PATH)gen \
		   --go-grpc_out=paths=source_relative:$(DIR_PATH)gen \
		   $(DIR_PATH)$(FILE_NAME).proto

.PHONY: http
# generate http code
http:
	 protoc --proto_path=. \
           --proto_path=./third_party \
           --go_out=paths=source_relative:$(DIR_PATH)gen \
           --go-http_out=paths=source_relative:$(DIR_PATH)gen \
          $(DIR_PATH)$(FILE_NAME).proto

.PHONY: errors
# generate errors code
errors:
	 protoc --proto_path=. \
           --proto_path=./third_party \
           --go_out=paths=source_relative:$(DIR_PATH)gen \
           --go-errors_out=paths=source_relative:$(DIR_PATH)gen \
          $(DIR_PATH)$(FILE_NAME).proto

.PHONY: swagger
# generate swagger
swagger:
	 protoc --proto_path=. \
	        --proto_path=./third_party \
	        --openapiv2_out $(DIR_PATH)gen \
	        --openapiv2_opt logtostderr=true \
           $(DIR_PATH)$(FILE_NAME).proto

.PHONY: protoc
# generate internal proto struct
protoc:
	protoc --proto_path=. \
           --proto_path=./third_party \
           --go_out=paths=source_relative:$(DIR_PATH)gen \
          $(DIR_PATH)$(FILE_NAME).proto

.PHONY: server
server:
	kratos proto server $(PROTO_PATH) -t $(APP_PATH)/$(PATH_WITHOUT_PROJECT)"internal/service"

UNAME_S := $(shell uname -s)
# 替换文件内容的目标
replace:
	@echo "Processing proto file: $(PROTO_PATH)"
	@echo "Path without project: $(PATH_WITHOUT_PROJECT)"
	@cd $(APP_PATH) && \
	LAST_FOLDER=$$(basename $(PATH_WITHOUT_PROJECT)) && \
	cd $$(dirname $(PATH_WITHOUT_PROJECT)) && \
	echo "Running find in: $$(pwd)/$$LAST_FOLDER" && \
	if [ "$(UNAME_S)" = "Darwin" ]; then \
		find $$LAST_FOLDER -type f -exec sed -i '' 's/Greeter/$(UPPER_FILE_NAME)/g' {} + && \
		find $$LAST_FOLDER -type f -exec sed -i '' 's/greeter/$(FILE_NAME)/g' {} + ; \
	else \
		find $$LAST_FOLDER -type f -exec sh -c '\
			sed -i "" -e "s/Greeter/$(UPPER_FILE_NAME)/g" $$0 \
		' {} \; && \
		find $$LAST_FOLDER -type f -exec sh -c '\
			sed -i "" -e "s/greeter/$(FILE_NAME)/g" $$0 \
		' {} \; ; \
	fi

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=. \
		   --proto_path=./third_party \
		   --go_out=paths=source_relative:. \
		   $(PROTO_PATH)