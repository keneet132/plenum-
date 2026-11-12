package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	"golang.zx2c4.com/wireguard/conn"
	"golang.zx2c4.com/wireguard/device"
	"golang.zx2c4.com/wireguard/tun"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)
func genkeys()(string , string , error ){
	pvtkey,err:=wgtypes.GeneratePrivateKey()
	if err!=nil{
		return "","",err
	}
	pvthex := hex.EncodeToString(pvtkey[:])
    pubkeybase64 := pvtkey.PublicKey().String()
	return pvthex,pubkeybase64,nil
}

func startwg(dev tun.Device,pvtkey string,listenport int)(*device.Device,error){
	logger:= device.NewLogger(device.LogLevelError, "[Plenum] ")
	wg:=device.NewDevice(dev,conn.NewDefaultBind(),logger)
	config:=fmt.Sprintf(
		"private_key=%s\nlisten_port=%d\n",
		pvtkey,
		listenport,
	)
	if err := wg.IpcSet(config); err != nil {
		wg.Close()
		return nil, fmt.Errorf("wireguard configuration failed: %w", err)
	}
	if err := wg.Up(); err != nil {
		wg.Close()
		return nil, fmt.Errorf("failed to start wireguard: %w", err)
	}
	return wg,nil 
}

func peers(wg *device.Device,peer Peerconfig) error{
	decodedPubKey, err := base64.StdEncoding.DecodeString(peer.PublicKey)
    if err != nil {
        return fmt.Errorf("failed to decode peer public key from base64: %w", err)
    }
	peerPubKeyHex := hex.EncodeToString(decodedPubKey)

	var config strings.Builder
	fmt.Fprintf(&config, "public_key=%s\n", peerPubKeyHex)
	if peer.Endpoint != "" {
		fmt.Fprintf(&config, "endpoint=%s\n", peer.Endpoint)
	}

	for _, ip := range peer.AllowedIPs {
		fmt.Fprintf(&config, "allowed_ip=%s\n", ip)
	}

	if err := wg.IpcSet(config.String()); err != nil {
		return fmt.Errorf("failed to add peer: %w", err)
	}
	return nil 
}