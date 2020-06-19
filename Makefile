build:
	docker build . -t armstrong
	docker run --rm --entrypoint cat armstrong /go/bin/armstrong > armstrong
	chmod u+x armstrong

.PHONY: build
