package interfaces

import "github.com/bougou/go-ipmi"

// 切换手动模式
func (s *InterfacesDefault) DellJobSetFan() error {
	_, err := s.Client.RawCommand(s.Ctx, ipmi.NetFn(0x30), 0x30, []byte{0x01, 0x00}, "Set Fan Adjust")
	if err != nil {
		return err
	}
	_, err = s.Client.RawCommand(s.Ctx, ipmi.NetFn(0x30), 0x30, []byte{0x02, 0xFF, byte(s.Actions.Fan)}, "Set Fan Speed")
	return err
}
