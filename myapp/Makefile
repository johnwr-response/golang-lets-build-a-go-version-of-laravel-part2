APP_NAME=myapp
BINARY_NAME=${APP_NAME}.exe

## clean: stops binary if running and cleans dist folder
clean: stop
	@echo ---
	@echo ################################################################################
	@echo Cleaning dist folder...
	@echo y | DEL /S dist
	@go clean
	@echo Cleaned!
	@echo ################################################################################

## build: stops binary if running, cleans dist folder and builds binary
build: clean
	@echo ---
	@echo ################################################################################
	@echo Building ${APP_NAME}...
	@go build -o dist/${BINARY_NAME} .
	@echo ${APP_NAME} built!
	@echo ################################################################################

## stop: stops binary if running
stop:
	@echo ---
	@echo ################################################################################
	@echo Stopping ${APP_NAME}...
	@taskkill /IM ${BINARY_NAME} /F  /FI "MemUsage gt 2"
	@echo Stopped ${APP_NAME}
	@echo ################################################################################

## run: starts binary in background
run:
	@echo ---
	@echo ################################################################################
	@echo Starting ${APP_NAME}...
	@start /min cmd /c .\dist\${BINARY_NAME} &
	@echo ${APP_NAME} started!
	@echo ################################################################################

## run: starts binary in foreground
run_fg:
	@echo ---
	@echo ################################################################################
	@echo Running ${APP_NAME} in foreground...
	@.\dist\${BINARY_NAME}
	@echo ################################################################################

## test: run all tests
test:
	@echo ---
	@echo ################################################################################
	@echo Testing ${APP_NAME}...
	@go test ./...
	@echo Done!
	@echo ################################################################################

## test_v: run all tests verbose
test_v:
	@echo ---
	@echo ################################################################################
	@echo Testing ${APP_NAME}...
	@go test -v ./...
	@echo Done!
	@echo ################################################################################

## test_integration: run all integration tests verbose
test_integration:
	@echo ---
	@echo ################################################################################
	@echo Testing ${APP_NAME}...
	@go test -v ./... --tags integration --count=1
	@echo Done!
	@echo ################################################################################

## test_cover: run all tests and displaying coverage in browser
test_cover:
	@echo ---
	@echo ################################################################################
	@echo Testing ${APP_NAME}...
	@go test '-coverprofile=coverage.out' ./...
	@echo Done!
	@echo ################################################################################

## test_cover: run all integration tests and displaying coverage in browser
test_integration_cover:
	@echo ---
	@echo ################################################################################
	@echo Testing ${APP_NAME}...
	@go test '-coverprofile=coverage.out' ./... --tags integration --count=1
	@echo Done!
	@echo ################################################################################

## test_coverage: run all tests and print coverage percent
test_coverage:
	@echo ---
	@echo ################################################################################
	@echo Testing ${APP_NAME}...
	@go test -cover ./...
	@echo Done!
	@echo ################################################################################

## test_integration_coverage: run all integration tests and print coverage percent
test_integration_coverage:
	@echo ---
	@echo ################################################################################
	@echo Testing ${APP_NAME}...
	@go test -cover ./... --tags integration --count=1
	@echo Done!
	@echo ################################################################################

## convenience methods
## start: starts binary in background
start: run

## start_fg: starts binary in foreground
start_fg: run_fg

## restart: restarts binary in background
restart: build start

## restart: restarts binary in foreground
restart_fg: build start_fg
