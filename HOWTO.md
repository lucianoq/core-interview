### Dependencies

This project needs:
 - `sqlite3` command to create and use the DB
 - `make` to build all
 - `go` to compile (I used `go version go1.7 darwin/amd64`)


### Build
To build the client and the server, please type

```
make
```


### Run

##### Server
The default port is 8080. If you want to change it, use the `-port` parameter.
To run the server:

```
./yoti_server
./yoti_server -port 8888
```

---

##### Store
The client should know the correct port of the server. Default is 8080 but you can change with `-port` parameter.

```
./yoti store -id example -data "This is a message"
./yoti store -id example -data "This is a message" -port 8888
```

The output is the AES key. Use that to retrieve.

---

##### Retrieve

```
./yoti retrieve -id example -key "TheKeyObtainedAbove"
./yoti retrieve -id example -key "TheKeyObtainedAbove" -port 8888
```

The output is the decrypted message stored.