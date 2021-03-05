# pinger
A very simple tool to test network connectivity

## How to use
[GIN-debug] GET    /start/:ip                --> main.main.func1 (3 handlers)
[GIN-debug] GET    /summary/:ip              --> main.main.func2 (3 handlers)
[GIN-debug] GET    /stop/:ip                 --> main.main.func3 (3 handlers)

/start/:ip  to launch a worker to ping {{ip}} 
/stop/:ip to stop that worker
/summary/:ip check summary of that worker

e.g. response
```
startTime: 2021-03-05 16:15:48.709027 +0800 CST m=+10.842523994 , endTime: 2021-03-05 16:16:44.742428 +0800 CST m=+66.874365032 
57 packets transmitted, 56 packets received, 1.7543859649122806% packet loss
round-trip min/avg/max/stddev = 173.981ms/181.859017ms/213.544ms/7.212067ms
```
