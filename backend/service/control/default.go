package control

import (
	"context"

	"github.com/bougou/go-ipmi"
)

type ControlService struct {
	Address  string
	Username string
	Password string
	Port     int
	Ctx      context.Context
	Client   *ipmi.Client
}

func NewService(address, username, password string, port int) *ControlService {
	return &ControlService{
		Address:  address,
		Username: username,
		Password: password,
		Port:     port,
		Ctx:      context.Background(),
	}
}

// 连接到服务器IPMI端口
func (s *ControlService) ConnectServer() error {
	client, err := ipmi.NewClient(s.Address, s.Port, s.Username, s.Password)
	if err != nil {
		return err
	}
	s.Client = client
	return client.Connect(s.Ctx)
}

// 断开服务器IPMI连接
func (s *ControlService) Close() error {
	return s.Client.Close(s.Ctx)
}
