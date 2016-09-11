.PHONY: db server client clean run-server

all: db server client

db:
	sqlite3 db.sqlite  ".read ./server/storage/db.sql"

server:
	go build -o yoti_server ./server/main

client:
	go build -o yoti ./client/main

run-server:
	./yoti_server >/dev/null 2>&1 &

clean:
	rm -f db.sqlite
	rm -f yoti_server
	rm -f yoti


