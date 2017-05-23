# Assets Tracking System

Internal Assets Tracking System for Maintain different Components.

[![Go Report Card](https://goreportcard.com/badge/github.com/improwised/cantaloupe)](https://goreportcard.com/report/github.com/improwised/cantaloupe)

# Development

## Dependencies
* Go 1.6
* [Glide](https://github.com/Masterminds/glide) - Package manager for Go
* PostgreSQL 9.4.4+
* npm

## Run postgres inside docker container from project directory

```
docker-compose up -d
```

## OR

```
docker run -d -it -v /Users/ashwin/cantaloupe/src/github.com/improwised/cantaloupe/postgres_data:/var/lib/postgresql/data -e POSTGRES_PASSWORD=123123 -e POSTGRES_USER=postgres -e POSTGRES_DB=cantaloupe -p 5432:5432 postgres:9.6-alpine
```

## Installation

### Back - end

* Setup go workspace by following this [Guide](https://golang.org/doc/code.html#Organization)
* Set Go vendor flag. Add the following to your `.bashrc` or `.zshrc` along with `$GOPATH` above
```
export GO15VENDOREXPERIMENT=1
```
* Install dependencies
```
$ glide install
```
* Create Database
```
$ create database cantaloupe
```
* Database Migration
```
$ go run ./db/migration.go
```
* Run project from cantaloupe
```
$ go run main.go
```
* access system on port 8000

# OR
### Using Makefile

To make necessary directories.
```
    $ make makedir
```

To download all dependencies.
```
    $ make get
```

To build the project and generate executable file.
```
    $ make build
```

To run the project.
```
    $ make run
```

To test the project.
```
    $ make test
```

To delete project dependencies and executable files.
```
    $ make clean
```

To make diractories, download dependencies, generating executable and then run the project.
```
    $ make all
```

# OR

### Docker-way
* Create database `cantaloupe`.
* Run migrations script from db directory `go run migrations.go`.
* Execute `docker-compose  up` from root.


### Front - end

* Switch to front-end directory.

* Run following commands to download dependencies.
```
    $ npm init

    $ npm install
```

* Start server by $ npm start

* Access system on port 8080

## Testing

TODO: How to run tests

Run following command from cantaloupe
```
GO_ENV2=testing go test
```
# Copyright & License

Copyright 2016 Improwised Technologies Pvt Ltd
www.improwised.com

This work is the sole property of Improwised Technologies Pvt Ltd brand
you may not use this work except in compliance with the License.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
