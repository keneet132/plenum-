package main

import (
	"fmt"

	"golang.zx2c4.com/wireguard/tun"
)


func createdev(name string)(tun.Device,error){
	const DefaultMTU=1420 
	tundev,err:= tun.CreateTUN(name,DefaultMTU)
	if err!=nil{
		return nil, fmt.Errorf("error %s %w", name, err)
	}
	
		println("created TUNdevice")
	return tundev,nil
}