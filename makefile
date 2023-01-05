.PHONY: run build clean test-app test-app-cover test-store test-store-cover test-http test-http-cover coverage

run:
	go run ./cmd/student-api-server/main.go

build: 
	CGO_ENABLED=0 GOOS=linux go build -v -installsuffix cgo ./cmd/student-api-server

clean:
	rm -rf student-api-server
	go clean -i .

test-app:
	go test ./pkg/app -v --cover

test-app-cover:
	go test ./pkg/app -v --coverprofile=cover.out && go tool cover -html=cover.out ; rm -f cover.outcover
	
test-store:
	go test ./pkg/store/memory -v --cover

test-store-cover:
	go test ./pkg/store/memory -v --coverprofile=cover.out && go tool cover -html=cover.out ; rm -f cover.outcover

test-http:
	go test ./pkg/http -v --cover

test-http-cover:
	go test ./pkg/http -v --coverprofile=cover.out && go tool cover -html=cover.out ; rm -f cover.outcover

coverage:
	go test -covermode=count -coverprofile=count.out fmt; rm -f count.out