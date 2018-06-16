.build:
	go build -ldflags="-s -w"

run:	.build
	./outbound-go

clean:
	rm ./outbound-go

release:
	zip outbound-linux-amd64.zip outbound-go templates/* .env.example.json -R
