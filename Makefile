all:
	go build src/main.go
	mv main wiki

install:
	install -Dt /usr/local/bin -m 755 wiki
