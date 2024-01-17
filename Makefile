all: run clean 

build:
	@go build

run: build
	@./new-cpp-proj

clean:
	@rm new-cpp-proj
