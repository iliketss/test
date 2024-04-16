# test
one of Interview`s demo

##### 有一组机器集群，集群内有若干台机器，现在想要知道这些机器之间的链路状态。


    假设所有机器使用的同一组mysql数据库集群
    基于gin+gorm+viper+mysql
    
    ...
##### 思路1:
    1.机器部署程序，自动获取本机IP，实现注册。
    2.拉取所有机器IP，未知各个IP之间的链路情况。
    3.使用第三方包ping来模拟机器之间的相互探测
    func PingByIp(ipAddr string, count int) *ping.Pinger 
    返回一个ping请求指针，可以计算探测结果，比如：round-trip min/avg/max/stddev、丢包率等。
    4.探测频率可以在PingByIp的count参数设置，0是无限。周期，可以自己设置。目前没有写在config.yaml文件。只在cron修改，测试我用的5s执行一次，直观一点。
    然后就是入库出库操作
    5.webapi就随便写了，gin框架的路由，/ips拉取所有机器，/rp（relationship）获取机器之间的情况
##### 思路2:
    1.使用 Go 语言中的 net 包来实现基本的网络通信，并使用时间戳来计算延迟
    我只做个大概，不在这个项目，也没有完善
    // 目标地址
    destAddr := "127.0.0.1:8888"
    // 建立UDP连接
	conn, err := net.Dial("udp", destAddr)
	if err != nil {
		fmt.Println("Failed to dial:", err)
		return
	}
    defer conn.Close()
    // 循环发送数据包
	for i := 0; ; i++ {
		// 构造消息
		message := fmt.Sprintf("Ping %d", i)
		// 发送消息
		start := time.Now()
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Failed to send message:", err)
			continue
		}
		// 接收响应
		response := make([]byte, 1024)
		_, err = conn.Read(response)
		if err != nil {
			fmt.Println("Failed to read response:", err)
			continue
		}
		// 计算延迟
		elapsed := time.Since(start)
		fmt.Printf("Received response: %s, Delay: %v\n", string(response), elapsed)
		// 等待一段时间再发送下一个消息
		time.Sleep(1 * time.Second)
	}
下面是模拟收数据，也是在同一个程序里面执行
    


    // 监听地址
    listenAddr := ":8888"
	// 建立UDP连接
	conn, err := net.ListenPacket("udp", listenAddr)
	if err != nil {
		fmt.Println("Failed to listen:", err)
		return
	}
	defer conn.Close()

	// 循环接收数据包
	for {
		// 接收消息
		message := make([]byte, 1024)
		n, addr, err := conn.ReadFrom(message)
		if err != nil {
			fmt.Println("Failed to read message:", err)
			continue
		}
		fmt.Printf("Received message from %s: %s\n", addr.String(), string(message[:n]))
		// 响应消息
		response := []byte("Pong")
		_, err = conn.WriteTo(response, addr)
		if err != nil {
			fmt.Println("Failed to send response:", err)
			continue
		}
	}
这一个我没有具体实现，这个方式比较原生，其实展开来写还是能写。