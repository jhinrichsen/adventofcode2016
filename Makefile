.phony: flame
flame:
	go test -run=xxx -bench=Day18Part2V3 -benchmem -cpuprofile=cpu.out -memprofile=mem.out
	go tool pprof -http :8081 cpu.out
	curl -o flame.html http://localhost:8081/ui/flamegraph
