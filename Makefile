# Makefile
NAME=terraform-provider-pingfederate
OS_NAME := $(shell uname -s | tr A-Z a-z)
CURDATE := ${shell date +'%y%m%d'}
TEST_COUNT          ?= 1
ACCTEST_TIMEOUT     ?= 180m
ACCTEST_PARALLELISM ?= 20

ifneq ($(origin TESTS), undefined)
	RUNARGS = -run='$(TESTS)'
endif

pf-init:
	@docker-compose up -d

checks:
	@go fmt ./...
	@go vet ./...
	@staticcheck ./...
	@gosec ./...
	@goimports -w ./internal
	@govulncheck ./...

fmt:
	@find ./internal/sdkv2provider -type f -name '*_test.go' | sort -u | xargs -I {} terrafmt fmt --fmtcompat {}
	@find ./internal/framework -type f -name '*_test.go' | sort -u | xargs -I {} terrafmt fmt --fmtcompat {}
	@find ./templates -type f | sort -u | xargs -I {} terrafmt fmt --fmtcompat {}

unit-test:
	@go test -mod=vendor ./... -v

sweep:
	@echo "WARNING: This will destroy infrastructure. Use only in development accounts."
	@go test ./internal/framework -v -sweep=all -timeout 60m
	@go test ./internal/sdkv2provider -v -sweep=all -timeout 60m

test:
	@rm -f pingfederate/terraform.log && TF_LOG=TRACE TF_LOG_PATH=./terraform.log TF_ACC=1 go test -mod=vendor ./... -v

test-framework:
	@TF_ACC=1 go test -mod=vendor ./internal/framework/. -v -trimpath -coverprofile=coverage.out && go tool cover -func=coverage.out

test-sdkv2:
	@TF_ACC=1 go test -mod=vendor ./internal/sdkv2provider/. -v -trimpath -coverprofile=coverage.out && go tool cover -func=coverage.out

test-and-report:
	@rm -f pingfederate/terraform.log coverage.out report.json
	@TF_LOG=TRACE TF_LOG_PATH=./terraform.log TF_ACC=1 go test -mod=vendor ./... -v -coverprofile=coverage.out -json > report.json && go tool cover -func=coverage.out

testacc:
	TF_ACC=1 go test ./internal/... -v -count $(TEST_COUNT) -parallel $(ACCTEST_PARALLELISM) $(RUNARGS) -timeout $(ACCTEST_TIMEOUT)

build:
	@go build -mod=vendor -o ${NAME} -gcflags "all=-trimpath=$GOPATH" .

install:
	@go install -gcflags "all=-trimpath=$GOPATH" -mod=vendor .

func-init:
	@rm -rf func-tests/.terraform.lock.hcl
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
