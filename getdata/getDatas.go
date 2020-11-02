package getdata

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"sync"
)

// SyncGetData get data
func SyncGetData(list []*Config, n string, w bool, tow int) (d []string) {
	var wg sync.WaitGroup
	// 开N个后台打印线程
	d = []string{}
	for i := 0; i < len(list); i++ {
		i0 := i
		wg.Add(1)
		go func() {
			res, errers := http.Get(list[i0].URL)
			if errers != nil {
				d = append(d, "")
			}
			if res != nil {
				defer res.Body.Close()
				if res.StatusCode != 200 {
					// log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
					d = append(d, "")
				}
				if list[i0].Types == "sub" {
					s, _ := ioutil.ReadAll(res.Body)
					codes := make([]byte, base64.StdEncoding.DecodedLen(len(s))) // 计算解码后的长度
					base64.StdEncoding.Decode(codes, s)
					// codes, err := base64.RawStdEncoding.DecodeString(string(s))
					// if err != nil {
					// 	fmt.Println(err)
					// }
					index := bytes.IndexByte(codes, 0)
					if index != -1 {
						codes = codes[:index]
					}
					d = append(d, string(codes))
				} else {
					data, status := ExampleScrape(n, w, tow, res.Body)
					if status {
						d = append(d, MakeList(data))
					}
				}
			}
			wg.Done()
		}()
	}
	// 等待N个后台线程完成
	wg.Wait()
	return
}
