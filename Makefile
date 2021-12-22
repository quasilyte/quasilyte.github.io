.PHONY: regexp-lint gophers-and-dragons gocorpus

#		wasm-opt main.wasm -Oz -o main.wasm && \

gocorpus:
	rm -rf gocorpus/
	mkdir -p gocorpus && \
		git clone --depth 1 https://github.com/quasilyte/gocorpus.git && \
		cd gocorpus && \
		make all

gophers-and-dragons:
	rm -rf gophers-and-dragons/
	mkdir -p gophers-and-dragons/src/github.com/quasilyte
	cd gophers-and-dragons && \
		git clone https://github.com/quasilyte/gophers-and-dragons.git src/github.com/quasilyte/gophers-and-dragons && \
		cd src/github.com/quasilyte/gophers-and-dragons && \
		GO111MODULE=on GOOS=js GOARCH=wasm go build -o www/go.wasm ./wasm && \
		tsc --target es6 ./www/game.ts
	cd gophers-and-dragons && \
		cp ./src/github.com/quasilyte/gophers-and-dragons/www/styles.css . && \
		cp ./src/github.com/quasilyte/gophers-and-dragons/www/index.html . && \
		cp ./src/github.com/quasilyte/gophers-and-dragons/www/game.html . && \
		cp ./src/github.com/quasilyte/gophers-and-dragons/www/wasm_exec.js . && \
		cp ./src/github.com/quasilyte/gophers-and-dragons/www/lz-string.js . && \
		cp ./src/github.com/quasilyte/gophers-and-dragons/www/game.js . && \
		cp -a ./src/github.com/quasilyte/gophers-and-dragons/www/img . && \
		cp ./src/github.com/quasilyte/gophers-and-dragons/www/go.wasm . && \
		wasm-opt go.wasm -Oz -o go.wasm && \
		rm -rf src

regexp-lint:
	rm -rf regexp-lint/
	mkdir -p regexp-lint/src/github.com/quasilyte
	cd regexp-lint && \
		git clone https://github.com/quasilyte/regexp-lint.git src/github.com/quasilyte/regexp-lint && \
		export GOPATH=`pwd` && \
		GOOS=js GOARCH=wasm go build -o main.wasm ./src/github.com/quasilyte/regexp-lint && \
		wasm-opt main.wasm -Oz -o main.wasm && \
		cp ./src/github.com/quasilyte/regexp-lint/www/wasm_exec.js . && \
		cp ./src/github.com/quasilyte/regexp-lint/www/index.html . && \
		rm -rf src
