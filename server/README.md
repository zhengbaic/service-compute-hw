# Go-server
  very simple one...use _beego_ web framework and implement some simple handlers.

#### How To Use
```
	http://localhost:3001/
	http://localhost:3001/temp
	http://localhost:3001/temp?age=AGE
```

#### Example

* http://localhost:3001/
```
hello
```
* http://localhost:3001/temp
```
{
  	"age": "no"
}
```
* http://localhost:3001/temp?age=11
```
{
      "age": 11
}
```

#### Test
```
$ curl -v http://localhost:3001/
* timeout on name lookup is not supported
*   Trying ::1...
* TCP_NODELAY set
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0* Connected to localhost (::1) port 3001 (#0)
> GET / HTTP/1.1
> Host: localhost:3001
> User-Agent: curl/7.55.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Server: beegoServer:1.9.0
< Date: Sat, 11 Nov 2017 03:01:51 GMT
< Content-Length: 5
< Content-Type: text/plain; charset=utf-8
<
{ [5 bytes data]
100     5  100     5    0     0      5      0  0:00:01 --:--:--  0:00:01   333hello
* Connection #0 to host localhost left intact
```

```
$ ./ab -n 1000 -c 100 http://localhost:3001/
This is ApacheBench, Version 2.3 <$Revision: 1807734 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests


Server Software:        beegoServer:1.9.0
Server Hostname:        localhost
Server Port:            3001

Document Path:          /
Document Length:        5 bytes

Concurrency Level:      100
Time taken for tests:   0.489 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      148000 bytes
HTML transferred:       5000 bytes
Requests per second:    2045.34 [#/sec] (mean)
Time per request:       48.892 [ms] (mean)
Time per request:       0.489 [ms] (mean, across all concurrent requests)
Transfer rate:          295.62 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.5      0       9
Processing:     1   44   6.9     47      98
Waiting:        1   44   6.9     47      98
Total:          1   45   6.9     47      99

Percentage of the requests served within a certain time (ms)
  50%     47
  66%     49
  75%     50
  80%     50
  90%     51
  95%     52
  98%     55
  99%     56
 100%     99 (longest request)
Finished 1000 requests
```

#### if I use net/http
	Recently I check out many offical doc in GO. And I find that it's easy to implement a simple server with net/http.
	The general process is like:
```
	Mux.HandleFunc("/", mainHandler)

	server := http.Server{
		Addr: (host, port)
		Handler: Mux
	}
	server.ListenAndServe()
```
#### Thanks 4 watching!

