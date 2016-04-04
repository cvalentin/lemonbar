all:
	godep go build

install:
	go get github.com/tools/godep

save:
	godep save ./...

clean:
	rm ./lemonbar
