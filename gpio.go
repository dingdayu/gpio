package gpio

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"syscall"
	"time"
)

const (
	export   = "/sys/class/gpio/export"
	unexport = "/sys/class/gpio/unexport"
	label    = "/sys/class/gpio/gpiochip0/label"
	path     = "/sys/class/gpio/gpio%d"

	direction = "/sys/class/gpio/gpio%d/direction"
	value     = "/sys/class/gpio/gpio%d/value"
)

const (
	OUT  = "out"
	IN   = "in"
	HIGH = "high"
	LOW  = "low"
)

func New() error {
	l, err := ioutil.ReadFile(label)
	if nil != err {
		return errors.New("not support gpio")
	}
	fmt.Println(string(l))
	return nil
}

type Pin struct {
	Port      int
	Direction string
	path      string
}

// 创建一个引脚: 端口,方向
func NewPin(port int, d string) *Pin {
	return &Pin{Port: port, Direction: d}
}

// 导出一个引脚
func (pin *Pin) Export() error {
	// export
	err := ioutil.WriteFile(export, []byte(strconv.Itoa(pin.Port)), syscall.O_WRONLY)
	if nil != err {
		_ = pin.UnExport()
		err = ioutil.WriteFile(export, []byte(strconv.Itoa(pin.Port)), syscall.O_WRONLY)
		if nil != err {
			return err
		}
	}
	time.Sleep(1 * time.Second)
	pin.path = fmt.Sprintf(path, pin.Port)

	// check dir
	//info, err := os.Stat(pin.path)
	//if err != nil || !os.IsExist(err) || !info.IsDir() {
	//	return errors.New("pin export failure")
	//}
	return ioutil.WriteFile(fmt.Sprintf(direction, pin.Port), []byte(pin.Direction), syscall.O_WRONLY)
}

func (pin *Pin) UnExport() error {
	return ioutil.WriteFile(unexport, []byte(strconv.Itoa(pin.Port)), syscall.O_WRONLY)
}

// 写入一个引脚
func (pin *Pin) Write(v int) error {
	return ioutil.WriteFile(fmt.Sprintf(value, pin.Port), []byte(strconv.Itoa(v)), syscall.O_RDONLY)
}
