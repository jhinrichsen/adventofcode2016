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

.PHONY: bench12
bench12: bench12part1.diff.bench bench12part2.diff.bench
bench12part1.diff.bench: bench12part1.bench bench12V2part1.bench
	benchstat  $^ | tee $@
bench12part2.diff.bench: bench12part2.bench bench12V2part2.bench
	benchstat  $^ | tee $@
bench12part1.bench:
	$(GO) test -run="^$$" -bench=Day12Part1 -benchmem -count 10 | tee $@
bench12V2part1.bench:
	$(GO) test -run="^$$" -bench=Day12V2Part1 -benchmem -count 10 | sed -e 's/V2//' | tee $@ 
bench12part2.bench:
	$(GO) test -run="^$$" -bench=Day12Part2 -benchmem -count 10 | tee $@
bench12V2part2.bench:
	$(GO) test -run="^$$" -bench=Day12V2Part2 -benchmem -count 10 | sed -e 's/V2//' | tee $@