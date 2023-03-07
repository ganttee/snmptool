
build.linux:
	GOARCH=amd64 GOOS=linux go build -o snmptool .

clean:
	rm -rf snmptool