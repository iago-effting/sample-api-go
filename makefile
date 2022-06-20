MAKEFLAGS += -s # log silence
MOD ?= mod
VERSION ?= 0.0.1 # we can change for git tags here
DSN ?=

OS := $(if $(GOOS),$(GOOS),$(shell go env GOOS))
ARCH := $(if $(GOARCH),$(GOARCH),$(shell go env GOARCH))

BUILD_DIRS := bin/$(OS)_$(ARCH)

$(BUILD_DIRS):
	mkdir -p $@

version:
	echo $(VERSION)
	
build: $(BUILD_DIRS)
	echo "# Building for $(OS)/$(ARCH)"	
	/bin/sh -c "								\
		VERSION=$(VERSION)						\
		PATH_BIN=$(BUILD_DIRS)					\
		./scripts/build.sh ./cmd/api/main.go	\
	"

run: $(BUILD_DIRS)
	echo "# Running for $(OS)/$(ARCH)"	
	make build
	/bin/sh -c "							\
		PATH_BIN=$(BUILD_DIRS)				\
		./scripts/run.sh 					\
	"

test:
	/bin/sh -c "				\
		MOD=$(MOD)				\
		./scripts/test.sh ./...	\
	"

lint:
	/bin/sh -c "				\
		./scripts/lint.sh ./...	\
	"
deps: ./bin/migrate
	/bin/sh -c "				\
		./scripts/deps.sh		\
	"

./bin/migrate:
	/bin/sh -c "					\
		./scripts/tool-deps.sh		\
	"