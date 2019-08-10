wasm:
	@cd module && GOARCH=wasm GOOS=js go build -o lib.wasm main.go
	@mv ./module/lib.wasm ./server

server: print
	@cd ./server && go run ./server.go

print:
	@echo "Running server..."