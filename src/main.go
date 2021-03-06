package main

import (
	"bytes"
	"log"
	"lyrics/pkg/lyricsMaker"
	"lyrics/pkg/player/cc"
	"lyrics/pkg/player/fmod"
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/bash", "-c", "pwd")

	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out

	//Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	err2 := cmd.Run()
	if nil != err2 {
		log.Println("err:", err2)
	}
	log.Println(out.String())
	player := fmod.GeneratePlayerFMOD()
	err := player.LoadMedia("/Users/mac/Documents/Projects/Go/lyrics/asserts/test.mp3")
	if nil != err {
		log.Println("Load err: ", err)
	}
	go player.Play()
	defer player.Close()
	cc.TestCTransfer()
	lyricsMaker.NewLyricsMakerClient().Start()
}
