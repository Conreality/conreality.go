PACKAGE := conreality
VERSION := $(shell cat VERSION)

GO = go
PANDOC  ?= pandoc

SOURCES := conreality.go
TARGETS :=

%.html: %.rst
	$(PANDOC) -o $@ -t html5 -s $<

all: build

build: $(TARGETS)
	$(GO) build

check: src/conreality_test.go
	$(GO) test

dist:
	@echo "not implemented"; exit 2 # TODO

install:
	$(GO) install

uninstall:
	@echo "not implemented"; exit 2 # TODO

clean:
	@$(GO) clean
	@rm -f *~ $(TARGETS)

distclean: clean

mostlyclean: clean

.PHONY: check dist install clean distclean mostlyclean
.SECONDARY:
.SUFFIXES:
