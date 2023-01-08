package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	fp := os.Args[0]
	fp = strings.Replace(fp, "shjc-", "", -1)
	if os.Args[0] == fp {
		fmt.Println("不能守护自己：", fp)
		return
	}
	fmt.Println("开始守护:", fp)
	loop := 1
	for {
		cmd := exec.Command(fp)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Start()
		if err != nil {
			//实际上程序会直接退出，不会执行这些程序。
			fmt.Fprintf(os.Stderr, "[-] Error: %s\n", err)
			break
		} else {
			fmt.Println(loop, time.Now())
			loop++
		}
		cmd.Wait()
	}
	fmt.Println("找不到程序：", fp)
	//os.Exit(0)
}
