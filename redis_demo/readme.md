## 1、使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。

~~~
测试环境
CPU：2.3 GHz 双核Intel Core i5  
内存：8G
~~~

redis-benchmark -c 50 -n 1000000  -t set,get -q -d 10  
SET: 105351.88 requests per second, p50=0.223 msec                    
GET: 98706.95 requests per second, p50=0.231 msec

redis-benchmark -c 50 -n 1000000  -t set,get -q -d 20  
SET: 99750.62 requests per second, p50=0.231 msec                     
GET: 105887.34 requests per second, p50=0.223 msec

redis-benchmark -c 50 -n 1000000  -t set,get -q -d 50  
SET: 103820.60 requests per second, p50=0.231 msec                    
GET: 103766.73 requests per second, p50=0.231 msec

redis-benchmark -c 50 -n 1000000  -t set,get -q -d 100  
SET: 101978.38 requests per second, p50=0.231 msec                    
GET: 103018.45 requests per second, p50=0.231 msec

redis-benchmark -c 50 -n 1000000  -t set,get -q -d 200  
SET: 92097.99 requests per second, p50=0.239 msec                    
GET: 97713.50 requests per second, p50=0.231 msec

redis-benchmark -c 50 -n 1000000  -t set,get -q -d 1000  
SET: 94939.71 requests per second, p50=0.239 msec                    
GET: 98658.25 requests per second, p50=0.231 msec

redis-benchmark -c 50 -n 1000000  -t set,get -q -d 5000  
SET: 50367.68 requests per second, p50=0.463 msec                   
GET: 79051.38 requests per second, p50=0.287 msec

redis-benchmark -c 50 -n 1000000  -t set,get -q -d 10000  
SET: 45504.19 requests per second, p50=0.503 msec                   
GET: 50895.76 requests per second, p50=0.471 msec

redis-benchmark -c 50 -n 1000000  -t set,get -q -d 100000  
SET: 11164.70 requests per second, p50=2.223 msec                   
GET: 9525.26 requests per second, p50=1.015 msec


可以看到value大小在100字节内时无明显区别，平均每秒处理的读写请求数在10万左右。  
value达到1k时，性能略微下降，可以接受。  
value达到5k时，写性能大幅下降，只有原来50%，读性能下降到80%。


## 2、写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息  , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

写入前内存使用状态
~~~
used_memory:871840
used_memory_human:851.41K
used_memory_rss:5709824
used_memory_rss_human:5.45M
used_memory_peak:931760
used_memory_peak_human:909.92K
used_memory_peak_perc:93.57%
used_memory_overhead:830352
used_memory_startup:809856
used_memory_dataset:41488
used_memory_dataset_perc:66.93%
~~~


写入10万key-value， 每个key10个字节，value 10个字节   
20 * 100,000 字节， 大约1.9M， 使用内存实际增长 11M左右
~~~
used_memory:12569088
used_memory_human:11.99M
used_memory_rss:17895424
used_memory_rss_human:17.07M
used_memory_peak:12590096
used_memory_peak_human:12.01M
used_memory_peak_perc:99.83%
used_memory_overhead:9327576
used_memory_startup:809856
used_memory_dataset:3241512
used_memory_dataset_perc:27.57%
~~~

写入10万key-value， 每个key10个字节，value 1000个字节  
1010 * 100,000 字节, 大约96M， 使用内存实际增长 107M左右
~~~
used_memory:113370096
used_memory_human:108.12M
used_memory_rss:122101760
used_memory_rss_human:116.45M
used_memory_peak:113391104
used_memory_peak_human:108.14M
used_memory_peak_perc:99.98%
used_memory_overhead:9327576
used_memory_startup:809856
used_memory_dataset:104042520
used_memory_dataset_perc:92.43%
~~~
