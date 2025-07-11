################################################################################
########################## Docker Compose shortcuts ############################
################################################################################

.PHONY: docker-up
docker-up:
	docker compose up -d

.PHONY: docker-down
docker-down:
	docker compose down

################################################################################
################################ Go shortcuts ##################################
################################################################################

# go-sqlite3 requires cgo to work
.PHONY: build
build:
	go tool templ generate
	go fmt ./...
	CGO_ENABLED=1 go build -o ./tmp/make_build ./main.go

.PHONY: run
run:
	./tmp/make_build 

.PHONY: build-run
build-run: build run

################################################################################
################################## Go Tests ####################################
################################################################################

.PHONY: test
test:
	GTA_ROOT_PATH=$(CURDIR) IS_TEST=true CGO_ENABLED=1 go test ./... $(ARGS)

.PHONY: test-v
test-v:
	GTA_ROOT_PATH=$(CURDIR) IS_TEST=true CGO_ENABLED=1 go test -v ./... $(ARGS)

.PHONY: test-ci
test-ci:
	GTA_ROOT_PATH=$(CURDIR) IS_TEST=true CGO_ENABLED=1 go test -race ./...

.PHONY: test-coverage
test-coverage:
	GTA_ROOT_PATH=$(CURDIR) IS_TEST=true CGO_ENABLED=1 go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

################################################################################
################################ API Documentation #############################
################################################################################

# Generate OpenAPI documentation
.PHONY: docs
docs:
	go tool swag init --parseDependency --parseInternal --output docs
	go tool swag fmt

# Legacy target for backward compatibility
.PHONY: swag
swag: docs

################################################################################
######################## Atlas shortcuts (Migrations) ##########################
################################################################################

.PHONY: migrate-gen
migrate-gen:
	atlas migrate diff --env gorm $(name)

.PHONY: migrate-new
migrate-new:
	atlas migrate new --env gorm $(name)

.PHONY: migrate-up
migrate-up:
	atlas migrate apply --env local

# See: https://atlasgo.io/versioned/apply#down-migrations
.PHONY: migrate-down
migrate-down:
	atlas schema apply --env local --to "file://migrations?version=$(version)&format=golang-migrate" --exclude "atlas_schema_revisions"
	atlas migrate set --env local $(version)