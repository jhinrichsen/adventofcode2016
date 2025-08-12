GO ?= CGO_ENABLED=0 go

.PHONY: all
all: lint test

.PHONY: clean
clean:
	rm README.pdf README.html

.PHONY: bench
bench:
	$(GO) test -bench=. -run="" -benchmem

.PHONY: lint
lint:
	gofmt -e -s -w *.go
	$(GO) vet
	$(GO) run honnef.co/go/tools/cmd/staticcheck@latest

.PHONY: test
test:
	$(GO) test -coverprofile=coverage.txt -covermode count gitlab.com/jhinrichsen/adventofcode2016
	$(GO) run github.com/boumenot/gocover-cobertura@latest < coverage.txt > coverage.xml

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

