# golog

### golog is a Go package extended the official [log](https://golang.org/pkg/log/) for witting logs into files.

#### Features

* rotating log files in customized directory
* compress log file into gzip format
* preserve specified number of history files
* 4 levels(DEBUG/INFO/WARN/ERROR) of output

#### Examples

##### Basic

```
package main

import "github.com/koalaxiaot/golog"

func main(){

    golog.Debug("This is debug string: ", "debug mode")
    golog.Info("This is info string")
    golog.Warn("This is warn string")
    golog.Error("This is error string")

}
```

```
$ go run main.go
$ cat ./logs/golog.log
DEBUG 2017/01/05 18:10:23 main.go:7: This is debug string:  debug mode
INFO  2017/01/05 18:10:23 main.go:8: This is info string
WARN  2017/01/05 18:10:23 main.go:9: This is warn string
ERROR 2017/01/05 18:10:23 main.go:10: This is error string
```

##### Customsize

```
package main

import "github.com/koalaxiaot/golog"

func main(){

    golog.Setlevel(golog.INFO).Setpath("./mylogs")

    golog.Debug("This is debug string: ", "debug mode") // will be omitted
    golog.Info("This is info string")
    golog.Warn("This is warn string")
    golog.Error("This is error string")

}
```
```
$ go run main.go
$ cat mylogs/golog.log 
INFO  2017/01/05 18:16:13 main.go:10: This is info string
WARN  2017/01/05 18:16:13 main.go:11: This is warn string
ERROR 2017/01/05 18:16:13 main.go:12: This is error string
```

#### Main Funcs

##### func (l \*Logger) Setpath(path string)

`golog.Setpath("/mydir")` Default: `./logs`

##### func (l \*Logger) Setlevel(level Level)

`golog.Setlevel(golog.WARN)` Default: `golog.DEBUG`

##### func (l \*Logger) SetRotateSize(size int64)

`golog.SetRotateSize(10)` Default: `20`

`golog.log` will be rotated to `2006-01-02-15-04-05.000.gz` when file size greater than 10MB

##### func (l \*Logger) SetMaxFileNum(num int)

`golog.SetMaxFileNum(20)` Default: `10`

max number of gzip files preserved

#### Tips

* `./logs/` is the default directory which log files saved(can be changed by `Setpath`).
* `golog.log` is name of the log file.
* `2006-01-02-15-04-05.000.gz` is the gzip filename format.
