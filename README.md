# zkteco

[![Build Status](https://travis-ci.org/athrunecho/zkteco.svg?branch=master)](https://travis-ci.org/athrunecho/zkteco)


package zkteco parses the output files from zkteco device KQ803 and update to redis.It's based on Redigo.

#### How to use zkteco

* Open opens the named file for reading. If successful, methods on the returned file can be used for reading

        if f, err = os.Open("/home/xx/1.csv"); err != nil {
		    fmt.Printf("os.Open err: %v\n", err)
	        return
        } 

* NewKaoqin returns an pointer to Kaoqin by given Redis address, Redis password. 

        if k, err = zkteco.NewKaoqin(":6379", ""); err != nil {
		    fmt.Printf("zkteco.NewKaoqin err: %v\n", err)
	        return
        }


* Update attendances records(*File) to redis..

        if err = k.UpdateAttendances(f); err != nil {
          	fmt.Printf("UpdateAttendances err:%v\n", err)
            return
        }

#### Installation

Install Redigo using the "go get" command:

    go get github.com/athrunecho/zkteco

#### Documentation
* [API Reference](http://godoc.org/github.com/athrunecho/zkteco)

#### License
* [MIT License](./LICENSE) 
