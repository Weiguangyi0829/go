package process
import (
	"sever-client/profect/method"
	"fmt"
	"net"
	"os"
)

func ShowMenu()  {
	fmt.Println("------------welcome to this chatroom,sir!------------------")
	fmt.Println("---------1、Display a list of online users---------------------")
	fmt.Println("---------2、Send the message---------------------")
	fmt.Println("---------3、A list of message---------------------")
	fmt.Println("---------4、Exit the System---------------------")
	var key int
	fmt.Scanf("%d\n",&key)
	switch key {
		case 1:
			fmt.Println("---------1、Display a list of online users---------------------")
		case 2:
			fmt.Println("---------2、Send the message---------------------")
		case 3:
			fmt.Println("---------3、A list of message---------------------")
		case 4:
			fmt.Println("---------4、Exit the System---------------------")
			os.Exit(0)
		default:
			fmt.Println("Your input is incorrect")
	}
}

func severProcessResMes(conn net.Conn)  {
	tf := &method.Transfer{
		Conn:conn,
	}
	for{
		fmt.Println("client is waiting the sever send the message")
		mes ,err := tf.Readpkg()
		if err != nil{
			fmt.Println("tf Readpkg err = ",err)
			return
		}
		fmt.Printf("mes = %v\n",mes)
	}
}