package control

import "github.com/bougou/go-ipmi"

// 手动设置风扇转速, 0-100
func (s *ControlService) DellFanAdjust(value uint8) error {
	// 切换手动模式
	_, err := s.Client.RawCommand(s.Ctx, ipmi.NetFn(0x30), 0x30, []byte{0x01, 0x00}, "Set Fan Adjust")
	if err != nil {
		return err
	}
	_, err = s.Client.RawCommand(s.Ctx, ipmi.NetFn(0x30), 0x30, []byte{0x02, 0xFF, value}, "Set Fan Speed")
	return err
}

// 风扇自适应
func (s *ControlService) DellFanAdaptive() error {
	_, err := s.Client.RawCommand(s.Ctx, ipmi.NetFn(0x30), 0x30, []byte{0x01, 0x01}, "Set Fan Adaptive")
	return err
}
