.PHONY: build run open

all: build run

build:
	gopherjs build client.go

run:
	go run server.go

open:
	@-for cmd in xdg-open open start; \
	  do if which $$cmd 2> /dev/null > /dev/null; \
	        then $$cmd http://localhost:3000; \
	     fi; \
	  done
