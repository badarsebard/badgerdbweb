# badgerdbweb
A simple web based badgerdb GUI Admin panel. Inspired by and forked from the [boltdbweb](https://github.com/evnix/boltdbweb) project by [evnix](https://github.com/evnix).


##### Installation
```
go get github.com/badarsebard/badgerdbweb
```

##### Usage
```
badgerdbweb --db-name=<DBfilename>[required] --port=<port>[optional] --static-path=<static-path>[optional]
```
- `--db-name:` The file name of the DB.
  - NOTE: If 'file.db' does not exist. it will be created as a BadgerDB file.
- `--port:` Port for listening on... (Default: 8080)
- `--static-path:` If you moved the binary to different folder you can determin the path of the `web` folder. (Default: Same folder where the binary is located.)


##### Example
```
badgerdbweb --db-name=test.db --port=8089 --static-path=/home/user/github/badgerdbweb
```
Goto: http://localhost:8089

##### Screenshots:

![](https://github.com/badgerdb/badgerdbweb/blob/main/screenshots/1.png?raw=true)

![](https://github.com/badgerdb/badgerdbweb/blob/main/screenshots/2.png?raw=true)

![](https://github.com/badgerdb/badgerdbweb/blob/main/screenshots/3.png?raw=true)
