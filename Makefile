.PHONY: regexp-lint

regexp-lint:
	rm -rf regexp-lint/
	mkdir -p regexp-lint/src/github.com/quasilyte
	cd regexp-lint && \
		git clone https://github.com/quasilyte/regexp-lint.git src/github.com/quasilyte/regexp-lint && \
		export GOPATH=`pwd` && \
		GOOS=js GOARCH=wasm go build -o main.wasm ./src/github.com/quasilyte/regexp-lint && \
		cp ./src/github.com/quasilyte/regexp-lint/www/wasm_exec.js . && \
		cp ./src/github.com/quasilyte/regexp-lint/www/index.html . && \
		rm -rf src
