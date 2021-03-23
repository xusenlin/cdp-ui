package main

import (
	"cdpui/js"
	"fmt"
	"github.com/zserge/lorca"
	"runtime"
)



func main() {
	var args []string
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}

	ui, err := lorca.New("", "", 1376, 900, args...)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ui.Close()

	if err := js.Bind(&ui); err != nil {
		fmt.Println(err)
		return
	}

	//ln, err := net.Listen("tcp", "127.0.0.1:0")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer ln.Close()
	//
	//go http.Serve(ln, http.FileServer(http.Dir("public/dist")))
	//err = ui.Load(fmt.Sprintf("http://%s", ln.Addr()))
	err = ui.Load("http://192.168.50.171:3000/")
	if err != nil {
		fmt.Println(err)
	}

	<-ui.Done()
}
