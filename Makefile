OUT = clickterm.exe

run: make
	@./$(OUT)

help: make
	@./$(OUT) --help

make: 
	@go build