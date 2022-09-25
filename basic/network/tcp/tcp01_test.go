package tcp

import "testing"

/*
网络编程(tcp)
golang的主要设计目标之一，就是面向大规模后端服务程序，网络通信这块是服务端程序必不可少也是至关重要的一部分。
一，网络编程有两种：
1,TCP socket编程
是网络编程的主流，之所以叫TCP socket编程，是因为底层是基于tcp/ip协议的。如qq聊天

TCP/IPv4 |  TCP/IPv6

2,b/s结构的http编程
浏览器访问服务器时，使用的就是http协议，而http底层依旧是tcp socket实现的。如：新浪.(这属于go web开发范畴)

书籍:

二，网络编程基础:
1,协议(TCP/IP): Transmission Control Protocol/Internet Protocol
传输控制协议/因特网协议，又叫网络通讯协议，这个协议是Internet最基本的协议、Internet国际互联网络的基础。
简单的说，就是由网络层的IP协议和传输层的TCP协议组成。

			语言(协议)
			中文(具体协议,http/https/smtp(SimpleMailTransferProtocol简单邮件传输协议)/.../自定义协议)
			英文
中国人    --------------->  外国人
	     <---------------

2,OSI与tcp/ip参考模型
---[OSI模型(理论)]----------[tcp/ip模型(现实)]-----
应用层(application)       应用层:(application)
表示层(presention) 		 smtp,ftp,telnet,http
------------------------------------------------
会话层(session)			 传输层:(transport)
传输层(transport)		 解释数据
------------------------------------------------
网络层(ip)				 网络层:(ip)定位ip地址和确定连接路径
------------------------------------------------
数据链路层(link)			 链路层:(link)与硬件驱动对话
物理层(physical)
------------------------------------------------

QQ相互通信案例:

				皮皮凤 QQ 中国														 皮皮双 QQ 美国
					   	0:send msg is "hello"
						     |
			应用层		1:App(标识是哪个app发过来的)|hello								 应用层 11.hello
							 |
			传输层(tcp)  2:TCP头(需要准确的将数据包发送到指定接收地址)|APP|hello				 传输层(tcp) 10.APP｜hello
皮皮凤tcp/ip					 |											 皮皮双tcp/ip
			IP层		3:IP头(需要记住我的IP和对方的IP)|TCP头｜APP｜hello				 IP层  9.TCP头｜APP｜hello
							 |
			链路层		4:桢头｜IP头|TCP头｜APP｜hello｜桢尾 (此时数据已变成二进制)		     链路层 8.IP头|TCP头｜APP｜hello
						     |																      |
						5:网卡------------6.网关----->路由---->路由...------------->7.网关------------
																					桢头｜IP头|TCP头｜APP｜hello｜桢尾
3,ip地址
每个internet上的主机和路由器都有一个ip地址，它包括网络号和主机号，ip地址有ipv4(32位)或ipv6(128位)。可通过ipconfig来查看.

4,端口
4.1,tcp/ip协议中的端口。
1),只要是做服务的程序，都必须要监听一个端口.
2),该端口就是其它程序和该服务通讯的通道.
3),一台电脑有65535个端口（1-65535，有些端口为有名端口，不能使用）.
4),一旦一个端口被某个程序监听(占用)，那么其他的程序就不能在该端口上监听.

4.2,端口分类：
1),0号是保留端口
2),1-1024是固定端口，又叫有名端口，即被某些程度固定使用，一般程序员不能用。
   22: ssh远程登录协议  23：telnet使用  21：FTP使用   25:smtp服务使用   80：iis使用   7:echo服务...
3),1025-65535是动态端口
这些端口，程序员可以使用

4.3,端口使用注意事项:
1)在计算机(尤其是做服务器)要尽可能的少开端口
2)一个端口只能被一个程序监听
3)如果使用netstat -an可以查看本机有哪些端口在监听
4)可以使用netstat -anb为查看监听端口的pid，在结合任务管理器关闭不安全的端口

5,tcp socket编程的客户端和服务端
服务端处理流程
1)监听端口 8888
2)接收client的tcp连接，建立client和server端的链接
3)创建goroutine,处理该链接的请求(通常client端会通过链接发送请求包)
客户端处理流程
1)建立与服务端的连接
2)发送请求数据，接收server返回的结果数据
3)关闭链接
*/
func TestTcp01(t *testing.T) {
	tcp11()
}

func tcp11() {

}
