package main

import (
	"fmt"
	"log"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

var records map[string]string

func main() {
	records = map[string]string{
		"google.com": "142.251.42.238",
		"amazon.com": "176.32.103.205",
	}

	addr := net.UDPAddr{
		Port: 8090,
		IP:   net.ParseIP("127.0.0.1"),
	}

	u, _ := net.ListenUDP("udp", &addr)

	for {
		tmp := make([]byte, 1024)
		_, addr, _ := u.ReadFrom(tmp)
		clientAddr := addr

		packet := gopacket.NewPacket(tmp, layers.LayerTypeDNS, gopacket.Default)
		dnsPacket := packet.Layer(layers.LayerTypeDNS)
		tcp, _ := dnsPacket.(*layers.DNS)
		serverDNS(u, clientAddr, tcp)
	}
}

func serverDNS(u *net.UDPConn, clientAddr net.Addr, request *layers.DNS) {
	replayMess := request
	var dnsAnswer layers.DNSResourceRecord
	dnsAnswer.Type = layers.DNSTypeA
	var ip string
	var err error
	var ok bool

	requestName := string(request.Questions[0].Name)
	ip, ok = records[requestName]
	if !ok {
		log.Printf("name %s not in records\n", requestName)
	}

	fmt.Printf("requestName: %s\n", requestName)
	a, _, _ := net.ParseCIDR(ip + "/24")
	dnsAnswer.IP = a
	dnsAnswer.Name = request.Questions[0].Name
	dnsAnswer.Class = layers.DNSClassIN

	replayMess.QR = true
	replayMess.ANCount = 1
	// replayMess.OpCode = layers.DNSOpCodeQuery
	replayMess.AA = true
	replayMess.Answers = append(replayMess.Answers, dnsAnswer)
	replayMess.ResponseCode = layers.DNSResponseCodeNoErr
	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{}

	err = replayMess.SerializeTo(buf, opts)
	if err != nil {
		panic(err)
	}

	_, _ = u.WriteTo(buf.Bytes(), clientAddr)
}
