package main

// 生产者
func producer(data chan int) {
	for i := 0; i < 10; i++ {
		// 传输数据
		data <- i
	}

	// 数据生产结束，关闭通道
	close(data)
}

// 消费者
func consumer(data chan int, done chan bool) {
	// 消费数据，直至通道关闭
	for x := range data {
		println("recv:", x)
	}

	// 结束消费
	done <- true
}

func main() {
	data := make(chan int)
	done := make(chan bool)

	go producer(data)
	go consumer(data, done)

	// 阻塞，直至消费者发回结束信号
	<-done
}
