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
make server
```

## Others

Docker replacement
```
brew install docker docker-compose
brew install colima
colima start
```

Install SQLC

https://docs.sqlc.dev/en/latest/overview/install.html

Install Go Mock
https://github.com/uber-go/mock

```
go install go.uber.org/mock/mockgen@latest
```

Add Path

```
nano ~/.zshrc
export PATH=$PATH:~/go/bin
```


#### Notes
1. make postgres-image (download postgres image)
2. make postgres (start container)
3. make createdb (make dropdb)
4. make sqlc (if sql files are not generated)
5. make migrateup (make migratedown)
6. make test