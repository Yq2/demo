package main

import (
	"fmt"
	"github.com/sony/sonyflake"
	"os"
	"time"
)

func main() {
	t, _ := time.Parse("2006-01-02", "2018-01-01")
	settings := sonyflake.Settings{
		StartTime:      t,
		MachineID:      nil, // 用户自定义的函数，如果不指定的话，默认是本机IP地址的低16位
		CheckMachineID: nil, // 由用户指定的ID冲突检查函数,可以考虑使用Redis set来检查
	}

	sf := sonyflake.NewSonyflake(settings)

	for i := 0; i < 30; i++ {
		id, err := sf.NextID()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("id: ", id)
	}
}
