# gkd

-----------------------------
golang key-value database

## usage

-----------------------------
### build server
```
go build -o out/server server/server.go
```
### build client
```
go build -o out/client client/client.go
```
### run and try str
![gkd_1.png](https://github.com/Roderland/img/blob/master/gkd_1.png)
### run and try list
![gkd_2.png](https://github.com/Roderland/img/blob/master/gkd_2.png)

## support command

-----------------------------

+ save // save data to file "test.data"
### str command
+ set
+ get
+ setnx
+ getset
+ strlen
+ append
+ incrby
### list command
+ lpush
+ rpush
+ lrange
+ llen
+ lpop
+ lindex
+ rpop 
+ lrem 
+ lset 
