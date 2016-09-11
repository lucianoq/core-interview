.PHONY: db server client clean

all: db server client

db:
	sqlite3 db.sqlite  ".read ./server/storage/db.sql"

server:
	go build -o yoti_server ./server/main

client:
	go build -o yoti ./client/main

clean:
	rm -f db.sqlite
	rm -f yoti_server
	rm -f yoti


