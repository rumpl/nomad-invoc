BIN_NAME=nomad-invoc
STATIC_FLAGS= CGO_ENABLED=0
GO_BUILD = $(STATIC_FLAGS) go build

BASE_INVOCATION_IMAGE_NAME=nomad-invoc


.PHONY bin/$(BIN_NAME): cmd
	$(GO_BUILD) -o $@ ./$<

clean:
	rm -fr bin

invocation-image:
	docker build -f Dockerfile.invocation-image --target=invocation -t $(BASE_INVOCATION_IMAGE_NAME) .
