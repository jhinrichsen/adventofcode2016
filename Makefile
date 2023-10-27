GO = CGO_ENABLED=0 go

.PHONY: all
all: lint test

.PHONY: setup
setup:
	$(GO) install honnef.co/go/tools/cmd/staticcheck@latest
	$(GO) get github.com/boumenot/gocover-cobertura

.PHONY: bench
bench:
	$(GO) test -bench=. -run="" -benchmem

.PHONY: lint
lint:
	$(GO) vet
	staticcheck

.PHONY: test
test:
	$(GO) test -coverprofile=coverage.txt -covermode count gitlab.com/jhinrichsen/adventofcode2016
	$(GO) run github.com/boumenot/gocover-cobertura < coverage.txt > coverage.xml

prof:
	$(GO) -bench=. -benchmem -memprofile mprofile.out -cpuprofile cprofile.out
	$(GO) pprof cpu.profile

# some asciidoc targets
.PHONY: doc
doc: README.html README.pdf

README.html: README.adoc
	asciidoctor $<

README.pdf: README.adoc
	asciidoctor-pdf -a allow-uri-read $<

.PHONY: clean
clean:
	rm README.pdf README.html
