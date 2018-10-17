# Notes

<!-- *Database setup, structure change, tests setup, etc.* -->

# Set up Postgres
```console
$ # First fire up postgres server
$ docker-compose run -d db
...
$ # Next, create database for test and development
$ docker-compose run db createdb -hdb -Upostgres go_assignment_dev
$ docker-compose run db createdb -hdb -Upostgres go_assignment_test
$ # Lastly, we are good go for test and development
$ docker-compose up
...
```

# Run test
```console
$ # Make sure you created database for test and postgres is running background.
$ # Run this as following.
$ # If you are running Docker on Linux, tmp dir is owned by root. So check/change the ownership first.
$ docker-compose run web \
  /usr/local/go/bin/go test -p 1 \
  /go/src/github.com/ablce9/go-assignment/adapters/http \
  /go/src/github.com/ablce9/go-assignment/providers/database \
  /go/src/github.com/ablce9/go-assignment/domain
...
```
