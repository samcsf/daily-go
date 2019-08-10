.PHONY: build
build: clean
	# 设置GOOS js + GOARCH wasm 就能生成wasm包，注意没goto，fmt跟gc很占空间
	GOOS=js GOARCH=wasm go build -o ./html/test.wasm ./hello/main.go
	# 从go预设的模板中调用js接口去读取go的wasm包
	cp $$(go env GOROOT)/misc/wasm/wasm_exec.html ./html/index.html
	cp $$(go env GOROOT)/misc/wasm/wasm_exec.js ./html/wasm_exec.js
		
clean:
	rm -f ./html/*

