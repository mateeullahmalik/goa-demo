# goa-demo

`goa-demo` is a demo app to explore [goa](#https://goa.design/learn/getting-started/) it provides a REST API.
To view full flow, please follow `/loans/{id}` endpoint in code. Only this has been entirely built. However, it will return a hard-coded loan object.


## Quick Start

 `goa-demo` can be run with

``` shell
go run main.go
```

it will run server on localhost:8080 and exposes following endpoints
1. list loans
``` shell
GET: /loans
```
2. get loan [id = 1 will return a dummy obj]
``` shell
GET: /loans/{id}
```

3. get users

``` shell
GET: /users
```

4. get user
``` shell
GET: /users/{id}
```

5. create user
``` shell
POST: /users
```

6. delete user
``` 
DELETE: /users/{id}
```
## Structure 

1. Design is defined in `/api/design`. It has 3 services right now; user, loan & swagger service.
2. Code generated by goa is in `/api/gen`
3. `/api/services` contains api services (handler funcs)
4. `/domain` contains domain entities & related functions on entities 
5. `/internal` defines common tools, wrappers above common assets like log, err etc.
6. `/services` contains business logic services
