all:
	mkdir -vp build
	make binary
	make deb

binary:
	go build -o build/wiki ./main.go

deb:
	mkdir -vp build/wiki-deb/usr/local/bin/
	go build -o build/wiki-deb/usr/local/bin/
	mkdir -vp build/wiki-deb/DEBIAN/
	cp -v wiki/deb/control build/wiki-deb/DEBIAN/
	cp -v wiki/deb/postinst build/wiki-deb/DEBIAN/
	chmod 775 build/wiki-deb/DEBIAN/postinst
	dpkg-deb --build build/wiki-deb

clean:
	rm -rv build

install:
	install -vDt /usr/local/bin -m 755 build/wiki

uninstall:
	rm -v /usr/local/bin/wiki
