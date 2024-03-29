# Keep test at the top so that it is default when `make` is called.
# This is used by Travis CI.
ci: clean coverage.txt
coverage.txt: install
	docker-compose stop integration-test-db
	docker-compose rm -f integration-test-db
	docker-compose up -d integration-test-db
	sleep 2
	go test -tags integration -race -coverprofile=coverage.txt -covermode=atomic -coverpkg=./pkg/...,./ ./...
	docker-compose stop integration-test-db
view-cover: clean coverage.txt
	go tool cover -html=coverage.txt
unit-test:
	go test ./test/...
build:
	go build ./...
install: build
	go install ./...
run:
	docker-compose stop db
	docker-compose rm -f db
	docker-compose up -d db
	sleep 2
	matchstick-video
inspect: build
	go vet
	golint ./...
update:
	go get -u ./...
	# TODO: Golint should be replaces by staticcheck and only be installed if binary is unavailable.
	go mod download golang.org/x/lint
	go get golang.org/x/lint@v0.0.0-20210508222113-6edffad5e616
	go install golang.org/x/lint/golint
pre-commit: update clean coverage.txt inspect
	go mod tidy
clean:
	rm -f ${GOPATH}/bin/matchstick-video
	rm -f coverage.txt