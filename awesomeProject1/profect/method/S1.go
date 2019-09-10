package method

import (
	"awesomeProject1/profect/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

//读取数据包
func Readpkg(conn net.Conn) (mes message.Message, err error)  {
	//将从客户端接受到是已经序列化好的数据
	//读取client 数据
	buf := make([]byte,8096)
	fmt.Println("读取客户端发送的数据.....")
	_, err = conn.Read(buf[:4])
	if  err !=nil{
		//err = errors.New("read pgk header err")
		return
	}
	//根据buf[:4](是从客户端接收到的数据)转成一个uint32类型
	var pkglen uint32  //包的长度
	pkglen = binary.BigEndian.Uint32(buf[0:4])
	//此时pkglen还是个链接中的uint32类型
	//根据pkglen 读取消息长度  判断是否掉包
	n , err := conn.Read(buf[:pkglen])
	if n != int(pkglen) || err != nil{
		//err = errors.New("read pgk header err")
		return
	}
	//读取消息
	//将连接中的pgklen反序列化  ->message.Message
	err = json.Unmarshal(buf[:pkglen],&mes)
	if err != nil{
		fmt.Println("json.Unmarshal err =" ,err)
		return
	}
	return
}

//发送函数
func Writepkg(conn net.Conn , data []byte) (err error)  {
	//先获取到data长度 ->转成一个表示长度的byte切片(字节类型的切片)
	var pkglen uint32
	pkglen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4],pkglen)
	//发送消息长度
	n ,err := conn.Write(buf[0:4])
	if n != 4 || err != nil{
		fmt.Println("conn write fail ", err)
		return
	}
	//发送消息本身
	n, err = conn.Write(data)
	if n != int(pkglen) || err != nil{
		fmt.Println("conn write fail ", err)
		return
	}
	return
}