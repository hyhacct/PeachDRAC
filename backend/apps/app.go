package apps

import (
	"PeachDRAC/backend/farmework"
	"PeachDRAC/backend/model"
	"PeachDRAC/backend/service/actions"
	"PeachDRAC/backend/service/config"
	"PeachDRAC/backend/service/survey"
	"context"
)

// App struct
type App struct {
	ctx             context.Context
	config_service  *config.ServiceConfig
	survey_service  *survey.ServiceSurvey
	actions_service *actions.ServiceActions
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

	// 初始化survey服务
	a.survey_service = survey.NewService(a.ctx)

	// 初始化actions服务
	a.actions_service = actions.NewService(a.ctx)
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
func (a *App) ConfigPassAddOrUpdate(form model.TablePass) model.WailsCommunicate {
	return a.config_service.AddOrUpdatePass(form)
}

/*
更新密码配置启用状态
*/
func (a *App) ConfigPassSwitch(id int, status bool) model.WailsCommunicate {
	return a.config_service.SwitchPass(id, status)
}

/*
删除密码配置
*/
func (a *App) ConfigPassDelete(id int) model.WailsCommunicate {
	return a.config_service.DeletePass(id)
}

/*
获取Java配置列表
*/
func (a *App) ConfigJavaGetList() model.WailsCommunicate {
	return a.config_service.GetAllJava()
}

/*
创建或修改Java配置
*/
func (a *App) ConfigJavaAddOrUpdate(form model.TableJava) model.WailsCommunicate {
	return a.config_service.AddOrUpdateJava(form)
}

/*
删除Java配置
*/
func (a *App) ConfigJavaDelete(id int) model.WailsCommunicate {
	return a.config_service.DeleteJava(id)
}

/*
开始探测
*/
func (a *App) SurveyStart(ips []string) model.WailsCommunicate {
	return a.survey_service.StartSurvey(ips)
}

/*
开始操作
*/
func (a *App) ActionsStart(ips []string, action string, fan int, nfs string) model.WailsCommunicate {
	return a.actions_service.Start(ips, action, fan, nfs)
}
