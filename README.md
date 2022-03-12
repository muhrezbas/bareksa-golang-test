# How to run

### 0. Config

a. first thing first .env.example change to .env

### 1. Run with docker

a. run  following command in root folder to build the app:

```
docker build --tag test-api .
```

b. run command below to run the app:

```
docker-compose up -d
```

c. remove docker image:

```
docker-compose down
```

### 2. Run with go command

a. run command below:

```
go get -d -v
```

b. run command below:

```
go run main.go
```

b. to stop, run command below:

```
ctrl + c
```
