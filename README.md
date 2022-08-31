# cloudflare-pgbouncer

simple application for test cloudflare pgbouncer and benchmarking

pgbouncer : <p/>
```
Concurrency Level:      100
Time taken for tests:   48.828 seconds
Complete requests:      50000
Failed requests:        3507
   (Connect: 0, Receive: 0, Length: 3507, Exceptions: 0)
Non-2xx responses:      3507
Keep-Alive requests:    50000
Total transferred:      11688476 bytes
HTML transferred:       4268336 bytes
Requests per second:    1045.41 [#/sec] (mean)
Time per request:       95.656 [ms] (mean)
Time per request:       0.957 [ms] (mean, across all concurrent requests)
Transfer rate:          238.66 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       2
Processing:     2   96 623.3     13   13822
Waiting:        2   96 623.3     13   13822
Total:          2   96 623.3     13   13822

Percentage of the requests served within a certain time (ms)
  50%     13
  66%     79
  75%     94
  80%    105
  90%    130
  95%    301
  98%    374
  99%    428
 100%  14822 (longest request)
```
<p/>
cloudflare pgbouncer 

```
Concurrency Level:      100
Time taken for tests:   45.637 seconds
Complete requests:      50000
Failed requests:        4009
   (Connect: 0, Receive: 0, Length: 4009, Exceptions: 0)
Non-2xx responses:      4009
Keep-Alive requests:    50000
Total transferred:      11722612 bytes
HTML transferred:       4292432 bytes
Requests per second:    1095.61 [#/sec] (mean)
Time per request:       91.273 [ms] (mean)
Time per request:       0.913 [ms] (mean, across all concurrent requests)
Transfer rate:          250.85 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       2
Processing:     3   91 553.3     13   12363
Waiting:        1   91 553.3     13   12363
Total:          3   91 553.3     13   12363

Percentage of the requests served within a certain time (ms)
  50%     13
  66%     76
  75%     93
  80%    102
  90%    132
  95%    298
  98%    317
  99%    467
 100%  12363 (longest request)
```
