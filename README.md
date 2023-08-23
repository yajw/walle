# walle

![Test](https://github.com/yajw/walle/actions/workflows/test.yml/badge.svg?branch=main)

# 使用

## 1. LPR利率
```go
package main

import(
	"time"
	
	"github.com/yajw/walle/thirdparty/lpr"
)

func main() {
	// 取 5 年期 LPR 利率
    println(lpr.Get5Y(time.Date(2018, time.Month(8), 2, 0, 0, 0, 0, time.Local)).String()) // 0.0485
	
	// 取 1 年期 LPR 利率
	println(lpr.Get1Y(time.Date(2018, time.Month(8), 2, 0, 0, 0, 0, time.Local)).String()) // 0.0485
}
```