# go-rest
Simple REST API to learn go programming.

## Running
```
$ go run .
```

## Endpoints
* GET /user
* GET /user/{userID}
* POST /user

## Example usage:
* GET /user
```
$ curl "http://localhost:8080/api/v1/user" -i
```
```
HTTP/1.1 200 OK
Date: Mon, 10 Aug 2020 11:47:59 GMT
Content-Length: 125
Content-Type: text/plain; charset=utf-8

[{"UserName":"aaa","Email":"aaa@example.com","IsEnabled":true},{"UserName":"ccc","Email":"ccc@example.com","IsEnabled":true}]
```

* GET /user/{userID}
```
curl -X GET "http://localhost:8080/api/v1/user/aaa" -i
```
```
HTTP/1.1 200 OK
Date: Mon, 10 Aug 2020 12:08:33 GMT
Content-Length: 61
Content-Type: text/plain; charset=utf-8

{"UserName":"aaa","Email":"aaa@example.com","IsEnabled":true}
```

* POST /user
```
$ curl -X POST "http://localhost:8080/api/v1/user" -H  "Content-Type: application/json" -d "{\"userName\": \"abcd\", \"email\":\"abcd@example.com\"}" -i
```
```
HTTP/1.1 201 Created
Date: Mon, 10 Aug 2020 11:48:47 GMT
Content-Length: 64
Content-Type: text/plain; charset=utf-8

{"UserName":"abcd","Email":"abcd@example.com","IsEnabled":false}
```
