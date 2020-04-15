.PHONY: regexp-lint

regexp-lint:
	rm -rf regexp-lint/
	mkdir -p regexp-lint/src/github.com/quasilyte
	git clone https://github.com/quasilyte/regexp-lint.git regexp-lint/src/github.com/quasilyte/regexp-lint
	export GOPATH=`pwd`/regexp-lint
	GOOS=js GOARCH=wasm go build -o regexp-lint/main.wasm ./regexp-lint/src/github.com/quasilyte/regexp-lint
	cp ./regexp-lint/src/github.com/quasilyte/regexp-lint/www/wasm_exec.js regexp-lint/
	cp ./regexp-lint/src/github.com/quasilyte/regexp-lint/www/index.html regexp-lint/
	rm -rf regexp-lint/src
