all:
	godep go build

save:
	godep save ./...

clean:
	rm ./lemonbar
