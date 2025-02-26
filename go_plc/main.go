package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// 发送比特位控制命令到 PLC
func sendBitControl(digit int, value int) {
	// 设置 PLC 比特位的接口 URL
	url := fmt.Sprintf("http://deviceshifu-plc.deviceshifu.svc.cluster.local/sendsinglebit?rootaddress=Q&address=0&start=0&digit=%d&value=%d", digit, value)

	// 发送 HTTP GET 请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应并打印
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	fmt.Printf("%s\n", string(body))
}

func main() {
	// 初始化 Q0 内存，将所有位清空为 0
	for i := 0; i < 16; i++ {
		sendBitControl(i, 0) // 将每个比特位设置为 0
		time.Sleep(2 * time.Millisecond)
	}
	// 循环递增
	for {
		// 设置每一位的 digit 和 value
		for i := 0; i < 16; i++ {
			sendBitControl(i, 1)
			time.Sleep(200 * time.Millisecond) // 控制步长，可以根据需求调整
		}
		for i := 0; i < 16; i++ {
			sendBitControl(i, 0)
			time.Sleep(200 * time.Millisecond) // 控制步长，可以根据需求调整
		}
	}

}
