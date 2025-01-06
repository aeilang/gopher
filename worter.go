package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

// 工作者
type Worker struct {
	inCh   chan []string  // 输入通道
	goodCh chan []string  // 输出通道
	badCh  chan []string  // 输出通道
	client http.Client    // http客户端，设置超时时间
	wg     sync.WaitGroup // 等待所有工作者结束
	done   chan struct{}  // 等待文件写入协程结束
}

// 开启n个工作者协程
func (w *Worker) Start(n int) {
	for range n {
		w.wg.Add(1)
		go w.process()
	}
}

// 生产者
func (w *Worker) Produce(rows [][]string) {
	for _, row := range rows {
		w.inCh <- row
	}

	close(w.inCh)
}

// 从通道里接收值，写入文件，n 是总条数，用来打印进度
func (w *Worker) WriteToFile(badCsv, goodCsv *csv.Writer, n int) {
	var finished int = 1 // 初始值是1, 原因为表头没算入。

	// 关闭done通道，表明，WriteToFile协程结束
	defer close(w.done)

	for {
		select {
		case row, ok := <-w.goodCh:
			if !ok {
				return
			}

			goodCsv.Write(row)

			finished++
			fmt.Printf("已完成: %.3f\n", float64(finished)/float64(n))

		case row, ok := <-w.badCh:
			if !ok {
				return
			}

			badCsv.Write(row)

			finished++
			fmt.Printf("已完成: %.3f\n", float64(finished)/float64(n))
		}
	}
}

// 工作者的实际逻辑
func (w *Worker) process() {
	defer w.wg.Done()

	for row := range w.inCh {
		if len(row) < 5 {
			continue
		}

		if w.isValidUrl(row[4]) {
			w.goodCh <- row
		} else {
			w.badCh <- row
		}
	}
}

// 判断url是否正常
func (w *Worker) isValidUrl(url_ string) bool {
	url_ = strings.TrimSpace(url_)

	_, err := url.ParseRequestURI(url_)
	if err != nil {
		return false
	}

	resp, err := w.client.Get(url_)

	if err != nil {
		return false
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		return true
	}

	return false
}
