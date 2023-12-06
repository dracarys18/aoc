DAY?=day1
BINARY_NAME?=advent_of_code_$(DAY)

run:
	go run ./$(DAY)

init:
	mkdir $(DAY) && cd $(DAY)
	go mod init $(DAY)

build:
	go build -o $(BINARY_NAME) ./$(DAY)

tidy:
	go mod tidy

input:
	curl 
clean:
	go clean
	rm $(BINARY_NAME)
