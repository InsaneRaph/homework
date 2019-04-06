#homeworkprojet

Rest full api for card scheme storage using gorm and mysql

**Instalation:**

Unsure you have go installed

run ``go get ./...`` to update your go path

run ``go build`` to build the app

**Options**

| Option        | default         | Usage  |
| ------------- |-------------| ----------------------|
| -ds  | root:root@tcp(127.0.0.1:3306)/homework?charset=utf8&parseTime=True | mysql datasource|
| -loglevel  | info   | log level [trace,debug,info,warn,error,fatal,panic]|
| -port  | 8888   | listen port|
| -secret  | changeme   | Jwt api secret|