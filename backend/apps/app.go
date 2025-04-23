package apps

import (
	"PeachDRAC/backend/farmework"
	"PeachDRAC/backend/model"
	"PeachDRAC/backend/service/config"
	"context"
)

// App struct
type App struct {
	ctx            context.Context
	config_service *config.ServiceConfig
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) Startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx

	// 初始化日志模块
	farmework.NewLogger()

	// 初始化数据库
	farmework.InitOrm()
	farmework.AutoMigrate(&model.TablePass{}, &model.TableJava{})

	// 初始化配置服务
	a.config_service = config.NewService()
}

// domReady is called after front-end resources have been loaded
func (a App) DomReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) BeforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) Shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Greet returns num + 2
func (a *App) Greet(num int) int {
	return num + 2
}

/*
获取密码列表
*/
func (a *App) ConfigPassGetList() model.WailsCommunicate {
	return a.config_service.GetAllPass()
}

/*
创建或修改密码配置
*/
func (a *App) ConfigPassAddOrUpdate(id int, Username string, Password string, Port string) model.WailsCommunicate {
	return a.config_service.AddOrUpdatePass(id, Username, Password, Port)
}

/*
删除密码配置
*/
func (a *App) ConfigPassDelete(id int) model.WailsCommunicate {
	return a.config_service.DeletePass(id)
}
