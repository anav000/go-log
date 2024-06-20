# Definitions
NAME = go-log
BUILD = $(CURDIR)/build
MAIN = $(CURDIR)/cmd/$(NAME)/main.go
DIST = $(CURDIR)/dist
MKDIR = mkdir -p
TAR = tar -czvf

M = $(shell printf "\033[34;1m▶\033[0m")
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null)
DATE    ?= $(shell date +%FT%T)

LDFLAGS = -w -s
LDFLAGS := $(LDFLAGS) -X '$(NAME)/cmd.Version=$(VERSION)' -X '$(NAME)/cmd.BuildDate=$(DATE)'

SOURCES := $(wildcard $(CURDIR)/cli/*.go) \
	$(wildcard $(CURDIR)/cmd/*.go) \
	$(wildcard $(CURDIR)/tools/*.go)

# Define the objects for distribution
OBJS = $(BUILD)/$(NAME) $(BUILD)/$(NAME)-arm $(BUILD)/$(NAME)-arm64

# Build targets for different architectures
$(BUILD)/$(NAME): $(SOURCES)
	@$(MKDIR) $(BUILD)
	$(info $(M) creating binary $@..)
	@GOOS=linux go build -ldflags '$(LDFLAGS)' -o $@ $(MAIN)

$(BUILD)/$(NAME)-arm: $(SOURCES)
	@$(MKDIR) $(BUILD)
	$(info $(M) creating binary $@..)
	@GOOS=linux GOARCH=arm go build -ldflags '$(LDFLAGS)' -o $@ $(MAIN)

$(BUILD)/$(NAME)-arm64: $(SOURCES)
	@$(MKDIR) $(BUILD)
	$(info $(M) creating binary $@..)
	@GOOS=linux GOARCH=arm64 go build -ldflags '$(LDFLAGS)' -o $@ $(MAIN)

# Build rule
build:	$(OBJS)

# Create the distribution package
$(DIST)/$(NAME)-$(VERSION).tar.gz: build
	@$(MKDIR) $(DIST)
	$(info $(M) creating archive $@..)
	@$(TAR) $@ -C $(BUILD) $(notdir $(OBJS))

# Dist rule
dist: $(DIST)/$(NAME)-$(VERSION).tar.gz

# Changelog rule
changelog:
	@git-chglog -o $(CURDIR)/CHANGELOG.md

# Clean rule
clean: ; $(info $(M) cleaning ..)
	@$(RM) $(OBJS)

# Fclean rule
fclean:	clean ; $(info $(M) cleaning all ..)
	@$(RM) $(DIST)/$(NAME)-$(VERSION).tar.gz

# Version rule
version:
	@echo $(VERSION)

# Default target
all: build

# Re rule
re:	clean	all

.PHONY: build version clean dist
