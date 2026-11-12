package main

import (
	"flag"
	"fmt"
	"log"
)

func main(){
	mode := flag.String("mode","","client or host")
	Pendpoint:= flag.String("peer","","peer's publicIP:PORT if client")
	flag.Parse()
	if *mode!="host" && *mode!="client"{
		log.Fatal("user did not select a mode. exiting program")
		return 
	}

	_, pubport, err := stunlookup()
	if err != nil {
		log.Fatalf("STUN lookup failed: %v", err)
	}
	fmt.Println("Mode:", *mode)
	dev,err:=createdev("plenum0")
	if err!=nil{
		fmt.Println(err)
		return 
	}
	defer dev.Close()
	pvtkey,pubkey,err:=genkeys()
	if err!=nil{
		fmt.Println("failed to generate wg keys",err)
	}
	fmt.Println("generated wg keys.")
	fmt.Printf("[+] My WireGuard Public Key: %s\n", pubkey)
	listenport:=pubport 
	wg,err:=startwg(dev,pvtkey,listenport)
	if err!=nil{
		log.Fatal("failed to start wg",err)
	}
	defer wg.Close()
	if *mode=="client"{
		if *Pendpoint==""{
			log.Fatal("no peer flag")
		}
		var hostpubkey string 
		fmt.Print("Enter Host's Public Key: ")
		fmt.Scanln(&hostpubkey)
		peer := Peerconfig{
			PublicKey:  hostpubkey,
			Endpoint:   *Pendpoint,
			AllowedIPs: []string{"0.0.0.0/0"},
		}
		if err := peers(wg, peer); err != nil {
			log.Fatal("Failed to add peer:", err)
		}
		fmt.Println("connection successful, sending handshake")
	}
	fmt.Println("connection succesful, host mode.")

	select{}
}