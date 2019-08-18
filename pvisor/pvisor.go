package pvisor

import (
	"bytes"
	"encoding/json"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/iesreza/gutil/log"
	"net/http"
)

var seq int64
var DebugURL string

func DebugPacket(p gopacket.Packet, mark string) {
	go func() {
		wrapper := Packet{}

		b := p.Data()
		tcpLayer := p.Layer(layers.LayerTypeTCP)
		udpLayer := p.Layer(layers.LayerTypeUDP)

		wrapper.SrcIP = ParseSrcIP(b)
		wrapper.DstIP = ParseDstIP(b)
		if tcpLayer != nil || udpLayer != nil {
			wrapper.SrcPort = ParseSrcPort(b)
			wrapper.DstPort = ParseDstPort(b)
		}

		wrapper.Time = p.Metadata().Timestamp.Unix()
		seq++
		wrapper.Sequence = seq
		if tcpLayer != nil {
			wrapper.Type = "TCP"
			wrapper.LayerData = tcpLayer.LayerPayload()
		} else if udpLayer != nil {
			wrapper.Type = "UDP"
			wrapper.LayerData = udpLayer.LayerPayload()
		} else {
			wrapper.Type = "IP"
		}
		wrapper.Mark = mark

		b, err := json.Marshal(wrapper)
		if err == nil {

			log.Warning("Send packet to %s", DebugURL)
			req, err := http.NewRequest("POST", DebugURL, bytes.NewBuffer(b))
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				log.Error(err)
			} else {
				defer resp.Body.Close()
			}
		}
	}()
}
