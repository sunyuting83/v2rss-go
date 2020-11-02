package getdata

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

// SyncGetData get data
func SyncGetData(list []*Config) (d []string) {
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
			fmt.Println(list[i0].Types)
			if list[i0].Types == "sub" {
				s, _ := ioutil.ReadAll(res.Body)
				fmt.Println(string(s))
				d = append(d, string(s))
			}
			wg.Done()
		}()
	}
	// 等待N个后台线程完成
	wg.Wait()
	return
}
