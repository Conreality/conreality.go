GO = go

PACKAGE = conreality
VERSION = `cat VERSION`

SOURCES = conreality.go

BINARIES =

all: build

build: $(BINARIES)
	$(GO) build

check: conreality_test.go
	$(GO) test

dist:
	@echo "not implemented" # TODO

install:
	$(GO) install

clean:
	@$(GO) clean
	@rm -f *~

distclean: clean

mostlyclean: clean

.PHONY: build check dist install clean distclean mostlyclean
