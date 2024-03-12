
PROJECT_NAME := "tempocerto-api"


.PHONY: all dep build clean lint

setup: ## Setup project
	@echo "Get the dependencies..."
	@make dep --silent 
	@echo "Install staticcheck to lint..."
	@go install honnef.co/go/tools/cmd/staticcheck@2022.1.2
	@echo "Install gosec to lint..."
	@go install github.com/securego/gosec/v2/cmd/gosec@v2.13.1
	@echo "Configuring hooks..."
	@git config core.hooksPath hooks/
	@chmod +x ./hooks/pre-commit
	@echo "Done."

create-temp:
	@mkdir -p tmp

init-db-test:
	sqlite3 ./database-test.db < ./init-db-test.sql

drop-db:
	sqlite3 ./database-test.db 'DELETE FROM roles; DELETE FROM permissions;'

coverage-merge:
	@rm -rf tmp/coverage-report-unit.cov
	@find tmp/ -name *.cov | xargs 3rd/gocovmerge/bin/gocovmerge > tmp/coverage-report-unit.cov

coverage-percent: ## Print coverage percent. run test-unit before
	@make coverage-cover --silent | grep total | awk '{print substr($$3, 1, length($$3)-1)}' # $$ only works in makefile

coverage-cover: 
	@go tool cover -func tmp/coverage-report-unit.cov

coverage-show: ## Open coverage report. run test-unit before
	@go tool cover -html=tmp/coverage-report-unit.cov

coverage-to-html: ## Open coverage report. run test-unit before
	@go tool cover -html=tmp/coverage-report-unit.cov -o tmp/coverage-report.html

test-verbose: create-temp ## Run verbose all tests
	go test -v -count=1 -race -coverprofile=tmp/coverage-report-unit.cov -covermode=atomic ./... 

test-unit: create-temp ## run tests	
	go test -count=1 -race -coverprofile=tmp/coverage-report-unit.cov -covermode=atomic --tags=unit ./...

test-unit-verbose: create-temp ## Run unit tests
	make init-db-test
	go test -v -count=1 -race -coverprofile=tmp/coverage-report-unit.cov -covermode=atomic --tags=unit ./...
##make drop-db
	
test: create-temp ## create temp folder 
	go test -count=1 -race -coverprofile=tmp/coverage-report.cov -covermode=atomic  ./...

test-clear-cache: ## clear test cache 
	go clean -testcache

test-e2e: ## run e2e tests
	@echo APPLICATION_URL=${APPLICATION_URL}
	@echo NEW_VERSION=${NEW_VERSION}
	@echo TEST_TIMEOUT=${TEST_TIMEOUT}
	go test --tags=e2e -v ./...

test-e2e-clear-cache: ## clear test cache 
	go clean -testcache

test-e2e-local: test-e2e-clear-cache ## run e2e tests local
	export
	APPLICATION_URL=http://localhost:8080 \
	APPLICATION_VERSION= \
	TEST_TIMEOUT=0 \
	go test --tags=e2e -v ./...


test-race:
	echo START > tmp/test-race.txt
	@for i in `seq 1 100`; \
		do echo [ $$i ]============================================ >> tmp/test-race.txt && \
			make test --silent >> tmp/test-race.txt; \
	done;
	echo "------------------------------" >> tmp/test-race.txt

go-to-uml:
	goplantuml -recursive .  > docs/role-api.puml