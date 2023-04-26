
# API with GraphQL in Golang

Hello, This is a simple project creating an API with GraphQL





## STEPS TO TUN THE APP

### CREATE THE IMAGE AND RUN THE CONTAINER
> docker build [IMAGE NAME]
> docker run [IMAGE NAME]


### INSTALL SQLITE ON LINUX AND CREATE THE TABLE FOR DATABASE
> sudo apt-get install sqlite3

> sudo sqlite3 database.db

> CREATE TABLE categories(id string, name string, description string);

>CREATE TABLE course(id string, name string, description string, categoryId string);

### RUN THE APP

> go run server.go
## FEATURES

- INSERT COURSE
- INSERT CATEGORY
- LIST COURSE
- LIST CATEGORY


## API DOCUMENTATION




#### Playground GraphQL

```http
  GET /
```





## Author

- [@bentojulio](https://www.github.com/octokatherine)

