# zkteco

package zkteco parses the output files from zkteco device KQ803 and update to redis.It's based on Redigo.

#### Update attendances records(*File) to redis..

        if err = k.UpdateAttendances(f); err != nil {
	fmt.Printf("GetCSVRecordserr:%v\n", err)
            return
        }

#### Installation

Install Redigo using the "go get" command:

    go get github.com/athrunecho/zkteco

#### Documentation
* [API Reference](http://godoc.org/github.com/athrunecho/zkteco)

#### License
* [MIT License](./LICENSE) 
