package pvisor

import (
	"encoding/binary"
	"net"
	"strconv"
)

type HexPair struct {
	Bytes  []byte
	Data   string
	Start  int
	Length int
}

type Packet struct {
	Time      int64
	Sequence  int64
	Payload   []byte
	LayerData []byte
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
	pair.Bytes = b[pair.Start : pair.Start+pair.Length]
	pair.Data = net.IP(pair.Bytes).String()
	return pair
}
func ParseDstIP(b []byte) HexPair {
	pair := HexPair{
		Start:  30,
		Length: 4,
	}
	pair.Bytes = b[pair.Start : pair.Start+pair.Length]
	pair.Data = net.IP(pair.Bytes).String()
	return pair
}

func ParseSrcPort(b []byte) HexPair {
	pair := HexPair{
		Start:  34,
		Length: 2,
	}
	pair.Bytes = b[pair.Start : pair.Start+pair.Length]
	pair.Data = strconv.Itoa(int(binary.BigEndian.Uint16(pair.Bytes)))
	return pair
}
func ParseDstPort(b []byte) HexPair {
	pair := HexPair{
		Start:  36,
		Length: 2,
	}
	pair.Bytes = b[pair.Start : pair.Start+pair.Length]
	pair.Data = strconv.Itoa(int(binary.BigEndian.Uint16(pair.Bytes)))
	return pair
}
