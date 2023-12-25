# HW solution

## Handlers:

### GET :/days_left_until
Returns how many days left until 1 Jan 2030 or query date.

Accept query parameter:
- `until` - format `2111-01-01`

Response body format: string

Examples
```
$ http ':1323/days_left_until?until=2111-01-01'                                                   1 ↵
HTTP/1.1 200 OK
Content-Length: 5
Content-Type: text/plain; charset=UTF-8
Date: Mon, 25 Dec 2023 10:18:54 GMT

31783


$ http ':1323/days_left_until'                 
HTTP/1.1 200 OK
Content-Length: 4
Content-Type: text/plain; charset=UTF-8
Date: Mon, 25 Dec 2023 10:19:45 GMT

2199


$ http ':1323/days_left_until?until=foo'       
HTTP/1.1 400 Bad Request
Content-Length: 42
Content-Type: text/plain; charset=UTF-8
Date: Mon, 25 Dec 2023 10:20:01 GMT

bad query. Accept like '?until=2030-01-01'
```


## Run and build

```
make run

make build
```
