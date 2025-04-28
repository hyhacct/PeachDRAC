package interfaces

import (
	"PeachDRAC/backend/model"
	"context"
	"net/http"

	"github.com/bougou/go-ipmi"
)

type InterfacesDefault struct {
	Address  string
	Username string
	Password string
	Port     int
	Ctx      context.Context
	Client   *ipmi.Client
	Actions  model.ModelActions

	Cookies []*http.Cookie // DELL-登录后的Cookies
	St1     string         // DELL-登录后的ST1
	St2     string         // DELL-登录后的ST2
}

// 连接到服务器IPMI, 返回错误, 错误为空即为成功
func (s *InterfacesDefault) Connect(Address string, Username string, Password string, Port int) error {
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
func (s *InterfacesDefault) Close() error {
	return s.Client.Close(s.Ctx)
}

// 获取服务器型号和SN以及厂商
func (s *InterfacesDefault) GetModelAndSN() (string, string, string, error) {
	system, err := s.Client.GetFRU(s.Ctx, 0x00, "Get Device ID")
	if err != nil {
		return "", "", "", err
	}
	model := string(system.ProductInfoArea.Name)                // 型号
	sn := string(system.ProductInfoArea.SerialNumber)           // 序列号
	manufacturer := string(system.ProductInfoArea.Manufacturer) // 厂商
	return model, sn, manufacturer, nil
}
