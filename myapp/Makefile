APP_NAME=myapp
BINARY_NAME=${APP_NAME}.exe

## build: builds all binaries
build: clean
	@echo ---
	@echo ################################################################################
	@echo Building ${APP_NAME}...
	@go build -o dist/${BINARY_NAME} .
	@echo ${APP_NAME} built!
	@echo ################################################################################

run: build
	@echo ---
	@echo ################################################################################
	@echo Starting ${APP_NAME}...
	@start /min cmd /c tmp\${BINARY_NAME} &
	@echo ${APP_NAME} started!
	@echo ################################################################################

run_fg:
	@echo ---
	@echo ################################################################################
	@echo Running ${APP_NAME} in foreground...
	@go run .\...
	@echo ################################################################################

clean: stop
	@echo ---
	@echo ################################################################################
	@echo Cleaning dist folder...
	@echo y | DEL /S dist
	@go clean
	@echo Cleaned!
	@echo ################################################################################

test:
	@echo ---
	@echo ################################################################################
	@echo Testing ${APP_NAME}...
	@go test ./...
	@echo Done!
	@echo ################################################################################

start: run

start_fg: run_fg

stop:
	@echo ---
	@echo ################################################################################
	@echo Stopping ${APP_NAME}...
	@taskkill /IM dist/${BINARY_NAME} /F  /FI "MemUsage gt 2"
	@echo Stopped ${APP_NAME}
	@echo ################################################################################

restart: stop start

restart_fg: stop start_fg
