package pvisor

import (
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
)

type HexPair struct {
	Bytes  string
	Data   string
	Start  int
	Length int
}

type Packet struct {
	Time      int64
	Sequence  int64
	Payload   string
	LayerData string
	Mark      string
	Type      string
	SrcIP     HexPair
	DstIP     HexPair
	SrcPort   HexPair
	DstPort   HexPair
}

func ParseSrcIP(b []byte) HexPair {
	pair := HexPair{
		Start:  26,
		Length: 4,
	}
	pair.Bytes = s(b[pair.Start : pair.Start+pair.Length])
	pair.Data = net.IP(pair.Bytes).String()
	return pair
}
func ParseDstIP(b []byte) HexPair {
	pair := HexPair{
		Start:  30,
		Length: 4,
	}
	pair.Bytes = s(b[pair.Start : pair.Start+pair.Length])
	pair.Data = net.IP(pair.Bytes).String()
	return pair
}

func ParseSrcPort(b []byte) HexPair {
	pair := HexPair{
		Start:  34,
		Length: 2,
	}
	pair.Bytes = s(b[pair.Start : pair.Start+pair.Length])
	pair.Data = strconv.Itoa(int(binary.BigEndian.Uint16(b[pair.Start : pair.Start+pair.Length])))
	return pair
}
func ParseDstPort(b []byte) HexPair {
	pair := HexPair{
		Start:  36,
		Length: 2,
	}
	pair.Bytes = s(b[pair.Start : pair.Start+pair.Length])
	pair.Data = strconv.Itoa(int(binary.BigEndian.Uint16(b[pair.Start : pair.Start+pair.Length])))
	return pair
}

func s(b []byte) string {
	return fmt.Sprint(b)
}
