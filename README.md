# killSeconds


###
This is ApacheBench, Version 2.3 <$Revision: 1843412 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 5000 requests
Completed 10000 requests
Completed 15000 requests
Completed 20000 requests
Completed 25000 requests
Completed 30000 requests
Completed 35000 requests
Completed 40000 requests
Completed 45000 requests
Completed 50000 requests
Finished 50000 requests


Server Software:
Server Hostname:        127.0.0.1
Server Port:            8080

Document Path:          /ks
Document Length:        12 bytes

Concurrency Level:      50
Time taken for tests:   163.770 seconds
Complete requests:      50000
Failed requests:        17817
   (Connect: 0, Receive: 0, Length: 17817, Exceptions: 0)
Non-2xx responses:      17817
Total transferred:      5826405 bytes
HTML transferred:       386196 bytes
Requests per second:    305.31 [#/sec] (mean)
Time per request:       163.770 [ms] (mean)
Time per request:       3.275 [ms] (mean, across all concurrent requests)
Transfer rate:          34.74 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    2   2.7      1      37
Processing:    10  161 152.0     66     664
Waiting:        3  114 116.7     49     644
Total:         16  163 154.4     67     674

Percentage of the requests served within a certain time (ms)
  50%     67
  66%    304
  75%    338
  80%    352
  90%    382
  95%    409
  98%    444
  99%    478
 100%    674 (longest request)
