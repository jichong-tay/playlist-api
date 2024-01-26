# Playlist Backend Server
-------------------------
Capstone Project of a Backend Go Server.

-------------------------
## Steps to run

1. Ensure Golang is install.

https://go.dev

2.  Ensure golang-migrate is install.

```
brew install golang-migrate
```

3. Run the following cmd for setup

```
make start-and-migrate
```

4. Start Backend Server
```
go run main.go
```

#### Notes
1. make postgres-image (download postgres image)
2. make postgres (start container)
3. make createdb (make dropdb)
4. make sqlc (if sql files are not generated)
5. make migrateup (make migratedown)
6. make test