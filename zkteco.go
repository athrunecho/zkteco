package zkteco

import (
        "fmt"
        "regexp"
        "strconv"

	"github.com/garyburd/redigo/redis"
)

type Kaoqin struct {
	redisAddr     string
	redisPassword string
	c             redis.Conn
}

type AbnormalRecord struct {
	Date               string
	ExceptedsBeginTime string
	ExceptedEndTime    string
	ActualBeginTime    string
	ActualEndTime      string
	Category           string
}

//Update Attendances Records To Redis field := fmt.Printf("%v-%02d", re.FindStringSubmatch(date[2]), k)

func (k *Kaoqin) UpdateAttendances(records [][]string) (err error) {
        var datepattern = `\(d{4}).(\d*)`
        var time = `\d{2}.\d{4}.\d{2}|\d{2}.\d{2}` 

        if records, err := zkteco.GetCSVRecords(p); err != nil {
         debugPrintf("GetCSVRecordserr:%v\n", err)
         }else{
             
             //year & month
             date := records[2]
             re := regexp.MustCompile(datepattern)
             d := re.FindString(date[2])
             
             //key
             l := len(records)
             for i := 4; i <= l-1; i += 2 {
             name := records[i]
             fmt.Printf("Name: %v\n", name[10])
             
                 for k := 1; k <= 31; k ++ {
                 columns := records[i+1]
                     if columns[k] != nil {
                        starconv
                        re := regexp.MustCompile(time)
                        v := re.FindString(columns[k])

                        //day = k
                        f := fmt.Sprintf("%v-%02d", d, k)
                        k.c.Do(HSET name[10], f, v)
                     } else {
                        continue
                  }                
                 }
             }
             //field = date + day 
        }	
	//k.c.Do(HSET name[10] re.FindStringSubmatch(date[2])+k  columns[k])
        //              key                field                    value 
	return nil
}





func (k *Kaoqin) UpdateArrangements(records [][]string) (err error) {

	//k.c.Do(HSET key field value)
	return nil
}

func (k *Kaoqin) GetAbnormalRecords(name, beginDate, endDate string) (records []AbnormalRecord, err error) {
	return []AbnormalRecord{}, nil

}
