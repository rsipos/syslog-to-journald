# Go parameters
GOCMD=go
INSTALL=install
RM=rm
GOBUILD=$(GOCMD) build
BINARY_NAME=syslog-to-journald

all: syslog-to-journald

syslog-to-journald:
	CGO_ENABLED=0 GOOS=linux $(GOBUILD) -a -ldflags '-extldflags "-static"' -o $(BINARY_NAME)

install: install-bin install-service

install-bin: syslog-to-journald
	$(INSTALL) -D -m0755 syslog-to-journald $(DESTDIR)/usr/local/bin/syslog-to-journald

install-service:
	$(INSTALL) -D -m0644 syslog-to-journald.service $(DESTDIR)/etc/systemd/system/syslog-to-journald.service

clean:
	$(RM) -f $(BINARY_NAME)

uninstall:
	$(RM) -f $(DESTDIR)/usr/local/bin/syslog-to-journald $(DESTDIR)/etc/systemd/system/syslog-to-journald.service
