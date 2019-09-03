package main

import (
	"fmt"
	"time"
)

func main() {
	a := 1567496760
	format := time.Unix(int64(a), 0)
	fmt.Println(format)
	//a := "http://mp.weixin.qq.com/s?__biz=MzU3ODE2NTMxNQ==&MID=2247485961&idx=1&sn=431af867d04efd973fd16df359365dd6&chksm=fd78c525ca0f4c334da2c677c1622f32058b7d3b89d255d5bb6e21a11a7f32407b67b13245bd&scene=27#wechat_redirect"
	//
	////index := strings.Index(a, "__biz=")
	////fmt.Println(index)
	////print(a[index:])
	//bizIndex := strings.Index(a, "__biz=")
	//if bizIndex == -1 {
	//	return
	//}
	//bizEnd := strings.Index(a[bizIndex:], "&")
	//biz := a[bizIndex+6 : bizEnd+bizIndex]
	//fmt.Println(biz)
	//
	////mid
	//midIndex := strings.Index(a, "mid=")
	//if midIndex == -1 {
	//	midIndex = strings.Index(a, "MID=")
	//}
	//midEnd := strings.Index(a[midIndex:], "&")
	//mid := a[midIndex+4 : midIndex+midEnd]
	//fmt.Println(mid)
	//
	//idxIndex := strings.Index(a, "&idx=")
	//if midIndex == -1 {
	//	idxIndex = strings.Index(a, "&idx=")
	//}
	//idxEnd := strings.Index(a[idxIndex+5:], "&")
	//idx := a[idxIndex+5 : idxEnd+idxIndex+5]
	//fmt.Println(idx)

	//compile := regexp.MustCompile(`biz=(\w*).*?mid=(\w*)\w+&idx=(\d+)`)
	//ids := compile.FindAllString(a, -1)
	//for k, v := range ids {
	//	fmt.Println(k, v)
	//}
}

func read(c chan bool, i int) {
	fmt.Printf(" go func : %d \n", i)
	<-c
}
