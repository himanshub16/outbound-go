.build:
	go build -ldflags="-s -w"

run:	.build
	./outbound-go

clean:
	rm ./outbound-go
