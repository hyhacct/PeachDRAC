package actions

import "PeachDRAC/backend/model"

func (s *ServiceActions) Stop() model.WailsCommunicate {
	s.r.ForceQuit() // 取消批量动作
	s.r.Wait()      // 等待批量动作全部退出
	return model.WailsSuccess("任务被手动结束", "")
}
