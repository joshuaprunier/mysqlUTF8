# mysqlUTF8
[![Build Status](https://travis-ci.org/joshuaprunier/mysqlUTF8.svg?branch=master)](https://travis-ci.org/joshuaprunier/mysqlUTF8) [![Coverage](http://gocover.io/_badge/github.com/joshuaprunier/mysqlUTF8)](http://gocover.io/github.com/joshuaprunier/mysqlUTF8) [![GoDoc](https://godoc.org/github.com/joshuaprunier/mysqlUTF8?status.svg)](https://godoc.org/github.com/joshuaprunier/mysqlUTF8) [![license](https://img.shields.io/badge/license-GPLv2-blue.svg)](https://raw.githubusercontent.com/joshuaprunier/mysqlUTF8/master/LICENSE)

mysqlUTF8 is a small library for handling MySQL UTF8 filename encoding. Database and table names with special characters are encoded on disk as detailed in the [MySQL documentation](http://dev.mysql.com/doc/en/identifier-mapping.html).

For example if you create a database named `Ω` the directory `@7P` will be created in the MySQL data directory. The NeedsEncoding() function determines if a string requires encoding and the EncodeFilename() function will return an encoded string.


```go
func NeedsEncoding(s string) bool

fmt.Println(mysqlUTF8.NeedsEncoding("test"))
fmt.Println(mysqlUTF8.NeedsEncoding("¿"))
fmt.Println(mysqlUTF8.NeedsEncoding("¢ent"))

// Output:
// false
// true
// true
```
```go
func EncodeFilename(s string) string

fmt.Println(mysqlUTF8.EncodeFilename("test"))
fmt.Println(mysqlUTF8.EncodeFilename("¿"))
fmt.Println(mysqlUTF8.EncodeFilename("(╯°□°)╯︵ ┻━┻"))

// Output:
// test
// @00bf
// @0028@256f@00b0@25a1@00b0@0029@256f@fe35@0020@253b@2501@253b
```
