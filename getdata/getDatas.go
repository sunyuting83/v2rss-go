package getdata

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

// SyncGetData get data
func SyncGetData(list []*Config, n string, w bool, tow int) (d []string) {
	var wg sync.WaitGroup
	// 开N个后台打印线程
	for i := 0; i < len(list); i++ {
		i0 := i
		wg.Add(1)
		go func() {
			res, err := http.Get(list[i0].URL)
			if err != nil {
				fmt.Println(err)
			}
			defer res.Body.Close()
			if res.StatusCode != 200 {
				// log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
				fmt.Println(res.StatusCode)
			}
			if list[i0].Types == "sub" {
				s, _ := ioutil.ReadAll(res.Body)
				// codes, err := base64.RawStdEncoding.DecodeString(string(s))
				// if err != nil {
				// 	fmt.Println(err)
				// }
				codes := make([]byte, base64.StdEncoding.DecodedLen(len(s))) // 计算解码后的长度
				base64.StdEncoding.Decode(codes, s)
				// strings.TrimRight(string(codes),"\n")
				d = append(d, strings.TrimRight(string(codes), "\n"))
			} else {
				data, status := ExampleScrape(n, w, tow, res.Body)
				if status {
					d = append(d, MakeList(data))
				}
			}
			wg.Done()
		}()
	}
	// 等待N个后台线程完成
	wg.Wait()
	return
}
