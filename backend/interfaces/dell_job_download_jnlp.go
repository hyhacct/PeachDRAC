package interfaces

import (
	"PeachDRAC/backend/utils"
	"fmt"
	"os/exec"
	"path"
)

// 下载JNLP文件并且启动
func (s *InterfacesDefault) DellJobDownloadJnlp() error {
	url := fmt.Sprintf("https://%s/viewer.jnlp(%s@0@%s,abcd,@1234567890@ST1=%s)", s.Address, s.Address, s.Address, s.St1)
	resp, err := utils.HttpPostSendAuth(url, s.Username, s.Password, "", s.Cookies, nil)
	if err != nil {
		return fmt.Errorf("下载JNLP文件失败: %v", err)
	}
	// 取JNLP文件目录
	dir_path, err := utils.MkdirJnlp()
	if err != nil {
		return fmt.Errorf("创建JNLP文件目录失败: %v", err)
	}
	// 将响应内容写入到目录下
	fileName := path.Join(dir_path, fmt.Sprintf("%s.jnlp", utils.RandomSuid()))
	err = utils.FileWriteToPath(fileName, []byte(resp))
	if err != nil {
		return fmt.Errorf("写入文件失败: %v", err)
	}
	// 执行下载的JNLP文件
	err = exec.Command("javaws", fileName).Run()
	if err != nil {
		return fmt.Errorf("执行JNLP文件失败: %v", err)
	}
	return nil
}
