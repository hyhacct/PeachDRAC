package encapsulation

import (
	"context"

	"github.com/bougou/go-ipmi"
)

type IPMI struct {
	Address  string
	Username string
	Password string
	Port     int
	Ctx      context.Context
	Client   *ipmi.Client
}

// 连接到服务器IPMI, 返回错误, 错误为空即为成功
func (s *IPMI) Connect(Address string, Username string, Password string, Port int) error {
	s.Address = Address
	s.Username = Username
	s.Password = Password
	s.Port = Port
	s.Ctx = context.Background()

	client, err := ipmi.NewClient(s.Address, s.Port, s.Username, s.Password)
	if err != nil {
		return err
	}

	s.Client = client // 赋值给 Client
	err = s.Client.Connect(s.Ctx)
	if err != nil {
		return err
	}
	return nil
}

// 断开服务器IPMI连接
func (s *IPMI) Close() error {
	return s.Client.Close(s.Ctx)
}

// 获取服务器型号和SN以及厂商
func (s *IPMI) GetModelAndSN() (string, string, string, error) {
	system, err := s.Client.GetFRU(s.Ctx, 0x00, "Get Device ID")
	if err != nil {
		return "", "", "", err
	}
	model := string(system.ProductInfoArea.Name)                // 型号
	sn := string(system.ProductInfoArea.SerialNumber)           // 序列号
	manufacturer := string(system.ProductInfoArea.Manufacturer) // 厂商
	return model, sn, manufacturer, nil
}
