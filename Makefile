.build:
	go build -ldflags="-s -w" -o outbound-go

run:	.build
	./outbound-go

clean:
	rm ./outbound-go

release:
	zip outbound-linux-amd64.zip outbound-go templates/* static/* .env.*.json Dockerfile docker.env docker-compose.yml ./LICENSE Procfile
