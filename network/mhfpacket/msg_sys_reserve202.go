package mhfpacket

import (
	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgSysReserve202 represents the MSG_SYS_reserve202
type MsgSysReserve202 struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysReserve202) Opcode() network.PacketID {
	return network.MSG_SYS_reserve202
}

// Parse parses the packet from binary
func (m *MsgSysReserve202) Parse(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgSysReserve202) Build(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}
