BIN_NAME=nomad-invoc
STATIC_FLAGS= CGO_ENABLED=0
GO_BUILD = $(STATIC_FLAGS) go build

.PHONY bin/$(BIN_NAME): cmd
	$(GO_BUILD) -o $@ ./$<

clean:
	rm -fr bin
