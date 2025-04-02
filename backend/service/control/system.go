package control

import "PeachDRAC/backend/model"

// 获取设备型号
func (s *ControlService) GetSystem() (model.DeviceSurvey, error) {

	system, err := s.Client.GetFRU(s.Ctx, 0x00, "Get Device ID")
	if err != nil {
		return model.DeviceSurvey{}, err
	}

	return model.DeviceSurvey{
		IP:           s.Address,
		ProductName:  string(system.ProductInfoArea.Name),
		SerialNumber: string(system.ProductInfoArea.SerialNumber),
		Manufacturer: string(system.ProductInfoArea.Manufacturer),
		Status:       true,
	}, nil
}
