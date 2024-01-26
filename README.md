Steps to run

Ensure Golang is install.

Run the following cmd

```
make start-and-migrate
```

Start WebServer
```
go run main.go
```

1. make postgres-image (download postgres image)
2. make postgres (start container)
3. make createdb (make dropdb)
4. make sqlc (if sql files are not generated)
5. make migrateup (make migratedown)
6. make test


PENDING
create test case for
  user
  user_playlist

DONE
test case for playlist