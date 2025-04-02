package apps

import (
	"PeachDRAC/backend/constants"
	"PeachDRAC/backend/model"
	"PeachDRAC/backend/modules"
	"PeachDRAC/backend/orm"
	"PeachDRAC/backend/service/common"
	"PeachDRAC/backend/service/config"
	"PeachDRAC/backend/service/dell"
	"PeachDRAC/backend/service/inspur"
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx           context.Context
	logsService   *modules.ModulesLogs  // 日志服务
	ormService    *orm.SQLite           // 数据库服务
	CommonService *common.CommonService // 通用服务
	DellService   *dell.DellService     // 戴尔服务
	InspurService *inspur.InspurService // 浪潮服务
	configService *config.ConfigService // 配置服务
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		DellService:   dell.NewService(),
		InspurService: inspur.NewService(),
	}
}

// Startup is called at application startup
func (a *App) Startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx

	// 启动日志模块
	a.logsService = modules.NewModulesLogs()
	a.logsService.InitLogger(constants.PathLog)

	// 启动数据库模块
	a.ormService = orm.NewSQLite()
	a.ormService.SyncTable(model.Config{})

	// 启动配置服务
	a.configService = config.NewConfigService(a.ormService)

	// 启动通用服务
	a.CommonService = common.NewService(a.ctx, a.DellService, a.InspurService, a.ormService)
}

// DomReady is called after front-end resources have been loaded
func (a App) DomReady(ctx context.Context) {
	// Add your action here
}

// BeforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) BeforeClose(ctx context.Context) (prevent bool) {
	return false
}

// Shutdown is called at application termination
func (a *App) Shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

/*
探测指定IP范围内的设备，并且自动识别型号

参数:

	ips: 需要探测的IP列表

返回:

	[{ip: '设备IP',model: '设备型号'}...]

PS: 如果探测失败，则返回model为：未知/离线
*/
func (a *App) CommonSurvey(ips []string) interface{} {
	return a.CommonService.Survey(ips)
}

/*
执行对应的动作

参数:

	type ActionRequest struct {
			Action string   `json:"action"`
			IPs    []string `json:"ips"`
			Fan    struct {
					Speed int `json:"speed"` // 调整风扇的转速，如果为-1则表示自适应
			} `json:"fan"`
			NFS struct {
					Mount struct {
							IP   string `json:"ip"`   // 挂载NFS的IP
							Path string `json:"path"` // 挂载NFS的路径
					} `json:"mount"`
			} `json:"nfs"`
	}
*/
func (a *App) CommonAction(actions model.ActionRequest) {
	a.CommonService.Action(actions)
}

/*
获取所有配置
*/
func (a *App) ConfigGetAll() model.ConfigRespond {
	return a.configService.GetAll()
}

/*
添加或更新配置
*/
func (a *App) ConfigAddOrUpdate(config model.Config) model.ConfigRespond {
	return a.configService.AddOrUpdate(&config)
}

/*
删除配置
*/
func (a *App) ConfigDelete(id int) model.ConfigRespond {
	return a.configService.Delete(id)
}
