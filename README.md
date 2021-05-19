# hello-go
go 'hello world' test task

## Prerequisites
docker
go 

## Development
* branch off latest main
* develop unit tests, make sure they fail
* develop new funtionality
* run unit tests, make sure they pass
* develop API tests, make sure they pass
* commit
* push
* create PR to main

## Build and run
#### localhost
```
go mod download
go build -o main .
./main
```

#### docker
```
docker build -t hello-go .
docker run -d -p 8080:8080
```

## Deployment
* feature branch -> local(docker) env -> unit testing + API testing + static analysis/linting ( -> bugfixes ) -> PR to main(develop)
* main(develop) branch -> dev env -> performance testing + security testing ( -> bugfixes ) -> branch release
* release branch -> staging env -> UAT testing ( -> bugfixes ) -> release candidate
* release candidate -> canary env -> canary testing ( -> rollback ) -> PR's to master and main(develop)
* master -> prod env
