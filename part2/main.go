package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// 打印程序耗时
	now := time.Now()
	defer func() {
		fmt.Printf("程序耗时： %v\n", time.Since(now))
	}()

	// 打开和创建相关文件
	data := mustFile(os.Open("data.csv"))
	defer data.Close()

	// os.Create 默认会覆盖原文件；如果不存在就会新建
	goodFile := mustFile(os.Create("good.csv"))
	defer goodFile.Close()

	badFile := mustFile(os.Create("bad.csv"))
	defer badFile.Close()

	dataCsv := csv.NewReader(data)

	badCsv := csv.NewWriter(badFile)
	defer badCsv.Flush() // 从缓存中刷新到badFile。

	goodCsv := csv.NewWriter(goodFile)
	defer goodCsv.Flush()

	// 读取所有结果
	rows, err := dataCsv.ReadAll()
	if err != nil {
		panic(err)
	}

	length := len(rows)
	if length < 2 {
		log.Fatal("只有表头，不需要处理")
	}

	// 写入表头
	badCsv.Write(rows[0])
	goodCsv.Write(rows[0])

	// worker pool
	w := Worker{
		inCh:   make(chan []string),
		goodCh: make(chan []string),
		badCh:  make(chan []string),
		client: http.Client{
			Timeout: 8 * time.Second,
		},
		done: make(chan struct{}),
	}

	// 开启一个生产者协程
	go w.Produce(rows[1:])

	// 开启100个worker 工作者协程
	w.Start(100)

	// 开启一个携程，收集工作者的结果
	go w.WriteToFile(badCsv, goodCsv, length)

	// 等待worker协程工作结束
	w.wg.Wait()
	// 结束后关闭结果通道。
	close(w.goodCh)
	close(w.badCh)

	// 等待WriteToFile运行结束
	<-w.done
}

// 工具函数，简便写法。must设计模式。
func mustFile(f *os.File, err error) *os.File {
	if err != nil {
		panic(err)
	}

	return f
}
