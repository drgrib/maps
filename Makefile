generate: clean
	go generate ./...

clean:
	rm -f map_*_*.go
