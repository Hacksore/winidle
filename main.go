package main

import (
	"time"
  "unsafe"
  "strings"
	"syscall"
  "os/exec"

  "github.com/alexflint/go-arg"
)

type (
  GlobalArgumentsBody struct {
    Time           int      `arg:"-t,--time,required" help:"How long to wait until executing a command"`
    Command        string    `arg:"-c,--command,required" help:"Path to server binary/script"`
  }
)

var (
	user32            = syscall.MustLoadDLL("user32.dll")
	kernel32          = syscall.MustLoadDLL("kernel32.dll")
	getLastInputInfo  = user32.MustFindProc("GetLastInputInfo")
	getTickCount      = kernel32.MustFindProc("GetTickCount")
	lastInputInfo struct {
		cbSize uint32
		dwTime uint32
  }

  GlobalArguments GlobalArgumentsBody
)

func IdleTime() uint32 {
	lastInputInfo.cbSize = uint32(unsafe.Sizeof(lastInputInfo))
	currentTickCount, _, _ := getTickCount.Call()
	r1, _, err := getLastInputInfo.Call(uintptr(unsafe.Pointer(&lastInputInfo)))
	if r1 == 0 {
		panic("error getting last input info: " + err.Error())
	}
	return (uint32(currentTickCount) - lastInputInfo.dwTime) / 1000
}

func main() {
  arg.MustParse(&GlobalArguments)

  commandString := strings.Split(GlobalArguments.Command, " ")
  args := commandString[1:]

	t := time.NewTicker(1 * time.Second)
	for range t.C {

    if int(IdleTime()) > GlobalArguments.Time {

      cmd := exec.Command(commandString[0], args...)
      _, err := cmd.Output()
      if err != nil {
        println("spwaned")
      }
    }
	}
}