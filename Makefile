# Makefile
VERSION ?= 0.0.1-BETA
NAME=terraform-provider-pingfederate_v${VERSION}

pf-init:
	@docker run --rm -d --hostname pingfederate -v `pwd`/pingfederate/pingfederate-data.zip:/opt/in/instance/server/default/data/drop-in-deployer/data.zip --name pingfederate -e PING_IDENTITY_DEVOPS_KEY=$(PING_IDENTITY_DEVOPS_KEY) -e PING_IDENTITY_DEVOPS_USER=$(PING_IDENTITY_DEVOPS_USER) -e PING_IDENTITY_ACCEPT_EULA=YES --publish 9999:9999 --publish 9031:9031 pingidentity/pingfederate:10.0.2-edge

checks:
	@go fmt ./...
	@staticcheck ./...
	@gosec ./...
	@goimports -w pingfederate

unit-test:
	@go test -mod=vendor ./... -v 

test:
	@rm -f pingfederate/terraform.log && TF_LOG=TRACE TF_LOG_PATH=./terraform.log TF_ACC=1 go test -mod=vendor ./... -v

test-and-report:
	@rm -f pingfederate/terraform.log coverage.out report.json
	@TF_LOG=TRACE TF_LOG_PATH=./terraform.log TF_ACC=1 go test -mod=vendor ./... -v -coverprofile=coverage.out -json > report.json && go tool cover -func=coverage.out

build:
	@go build -mod=vendor -o ${NAME} -gcflags "all=-trimpath=$GOPATH" .

release:
	@rm -rf build/*
	GOOS=darwin GOARCH=amd64 go build -o -mod=vendor build/darwin_amd64/${NAME} -gcflags "all=-trimpath=$GOPATH" . && zip -j build/darwin_amd64.zip build/darwin_amd64/${NAME}
	# GOOS=freebsd GOARCH=386 go build -o -mod=vendor build/freebsd_386/${NAME} -gcflags "all=-trimpath=$GOPATH" .
	# GOOS=freebsd GOARCH=amd64 go build -o -mod=vendor build/freebsd_amd64/${NAME} -gcflags "all=-trimpath=$GOPATH" .
	# GOOS=freebsd GOARCH=arm go build -o -mod=vendor build/freebsd_arm/${NAME} -gcflags "all=-trimpath=$GOPATH" .
	# GOOS=linux GOARCH=386 go build -o -mod=vendor build/linux_386/${NAME} -gcflags "all=-trimpath=$GOPATH" .
	GOOS=linux GOARCH=amd64 go build -o -mod=vendor build/linux_amd64/${NAME} -gcflags "all=-trimpath=$GOPATH" . && zip -j build/linux_amd64.zip build/linux_amd64/${NAME}
	# GOOS=linux GOARCH=arm go build -o -mod=vendor build/linux_arm/${NAME} -gcflags "all=-trimpath=$GOPATH" .
	# GOOS=openbsd GOARCH=386 go build -o -mod=vendor build/openbsd_386/${NAME} -gcflags "all=-trimpath=$GOPATH" .
	# GOOS=openbsd GOARCH=amd64 go build -o -mod=vendor build/openbsd_amd64/${NAME} -gcflags "all=-trimpath=$GOPATH" .
	# GOOS=solaris GOARCH=amd64 go build -o -mod=vendor build/solaris_amd64/${NAME} -gcflags "all=-trimpath=$GOPATH" .
	# GOOS=windows GOARCH=386 go build -o -mod=vendor build/windows_386/${NAME} -gcflags "all=-trimpath=$GOPATH" .
	GOOS=windows GOARCH=amd64 go build -o -mod=vendor build/windows_amd64/${NAME} -gcflags "all=-trimpath=$GOPATH" . && zip -j build/windows_amd64.zip build/windows_amd64/${NAME}
	
deploy-local:
	@mkdir -p ~/.terraform.d/plugins
	@cp ${NAME} ~/.terraform.d/plugins/

func-init:
	@rm -rf func-tests/.terraform
	@rm -rf func-tests/crash.log
	@rm -rf func-tests/run.log
	@cd func-tests && terraform init

func-plan:
	@cd func-tests && TF_LOG=TRACE TF_LOG_PATH=./terraform.log terraform plan

func-apply:
	@cd func-tests && TF_LOG=TRACE TF_LOG_PATH=./terraform.log terraform apply -auto-approve

func-destroy:
	@cd func-tests && terraform destroy -auto-approve
	
func-cli-gen:
	@cd ../pingfederate-sdk-go-gen-cli/ && make generate

sonar:
	@sonar-scanner \
		-Dsonar.projectKey=github.com.iwarapter.terraform-provider-pingfederate \
		-Dsonar.sources=. \
		-Dsonar.host.url=http://localhost:9001 \
		-Dsonar.login=8f46eca6fcc26e1cbb653f634458a89df72190c8 \
		-Dsonar.go.coverage.reportPaths=coverage.out \
		-Dsonar.go.tests.reportPaths=report.json \
		-Dsonar.exclusions=vendor/**/* \
		-Dsonar.tests="." \
		-Dsonar.test.inclusions="**/*_test.go"
		
