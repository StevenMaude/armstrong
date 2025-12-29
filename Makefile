build-linux:
	docker build . -t armstrong
	docker run --rm --entrypoint cat armstrong /go/bin/armstrong > armstrong_linux_amd64
	chmod u+x armstrong_linux_amd64

build-windows:
	docker build --build-arg GOOS=windows . -t armstrong
	docker run --rm --entrypoint cat armstrong /go/bin/windows_amd64/armstrong.exe > armstrong.exe

build-linux-local:
	go build armstrong.go
	mv armstrong armstrong_linux_amd64

clean-linux-local:
	rm armstrong_linux_amd64

.PHONY: build-linux build-windows
