package pvisor

import (
	"bytes"
	"encoding/json"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"net/http"
	"packet-visor/helper"
)


var seq int64
var  DebugURL string
func DebugPacket(p gopacket.Packet,mark string)  {

	wrapper := helper.Packet{}

	b := p.Data()
	tcpLayer := p.Layer(layers.LayerTypeTCP)
	udpLayer := p.Layer(layers.LayerTypeUDP)

	wrapper.SrcIP = helper.ParseSrcIP(b)
	wrapper.DstIP = helper.ParseDstIP(b)
	if tcpLayer != nil || udpLayer != nil {
		wrapper.SrcPort = helper.ParseSrcPort(b)
		wrapper.DstPort = helper.ParseDstPort(b)
	}

	wrapper.Time = p.Metadata().Timestamp.Unix()
	seq++
	wrapper.Sequence = seq
	if tcpLayer != nil{
		wrapper.Type = "TCP"
		wrapper.LayerData = tcpLayer.LayerPayload()
	}else if udpLayer != nil{
		wrapper.Type = "UDP"
		wrapper.LayerData = udpLayer.LayerPayload()
	}else{
		wrapper.Type = "IP"
	}
	wrapper.Mark = mark

	b,err := json.Marshal(wrapper)
	if err == nil{

		req, err := http.NewRequest("POST", DebugURL, bytes.NewBuffer(b))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

	}
}
