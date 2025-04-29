package interfaces

import "github.com/bougou/go-ipmi"

// 风扇自适应
func (s *InterfacesDefault) DellJobSetFanAuto() error {
	_, err := s.Client.RawCommand(s.Ctx, ipmi.NetFn(0x30), 0x30, []byte{0x01, 0x01}, "Set Fan Adaptive")
	return err
}
