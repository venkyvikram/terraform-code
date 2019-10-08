COMMIT := $(shell git log | head -n 1 | cut  -f 2 -d ' '))
DATE := $(shell date +"%d%m%Y%H%M%S")

clean:
	rm -rf bin/
build-mac:
	env GOARCH=amd64 GOOS=darwin go build -ldflags "-X synergy/cmd.commit=$(COMMIT) -X synergy/cmd.date=$(DATE)" -o bin/mac/synergy
build-linux:
	env GOOS=linux GOARCH=amd64 go build -ldflags "-X synergy/cmd.commit=$(COMMIT) -X synergy/cmd.date=$(DATE)" -o bin/linux/synergy
