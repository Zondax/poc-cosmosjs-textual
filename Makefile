build:
	mkdir -p output
	GOOS=js GOARCH=wasm go build -o output/main.wasm