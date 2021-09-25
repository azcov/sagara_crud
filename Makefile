# This version-strategy uses git tags to set the version string
APP_NAME = article
VERSION := $(shell git describe --tags --always --dirty)
GIT_COMMIT := $(shell git rev-list -1 HEAD)

version: ## Show version
	@echo $(APP_NAME) $(VERSION) \(git commit: $(GIT_COMMIT)\)

local:
	air -c .air.toml
