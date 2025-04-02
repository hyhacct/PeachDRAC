package control

import "github.com/bougou/go-ipmi"

// 其他选项：
// - ipmi.ControlPowerDown: 关机
// - ipmi.ControlPowerCycle: 电源循环（重启）
// - ipmi.ControlPowerHardReset: 硬重启
// - ipmi.ControlPowerPulseDiag: 诊断脉冲
// - ipmi.ControlPowerAcpiSoft: ACPI 软关机

// 电源控制开机
func (s *ControlService) PowerOn() error {
	_, err := s.Client.ChassisControl(s.Ctx, ipmi.ChassisControlPowerUp)
	return err
}

// 电源控制关机
func (s *ControlService) PowerOff() error {
	_, err := s.Client.ChassisControl(s.Ctx, ipmi.ChassisControlPowerDown)
	return err
}

// 电源控制软重启
func (s *ControlService) PowerRestart() error {
	_, err := s.Client.ChassisControl(s.Ctx, ipmi.ChassisControlPowerCycle)
	return err
}

// 电源控制硬重启
func (s *ControlService) PowerHardRestart() error {
	_, err := s.Client.ChassisControl(s.Ctx, ipmi.ChassisControlHardReset)
	return err
}

// 电源控制软关机
func (s *ControlService) PowerSoftShutdown() error {
	_, err := s.Client.ChassisControl(s.Ctx, ipmi.ChassisControlSoftShutdown)
	return err
}
