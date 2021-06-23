package task

import (
	"github.com/ZongweiBai/golang-in-action/config"
	cron "github.com/robfig/cron/v3"
)

func SetupTasks() {
	config.LOG.Infof("初始化定时任务")
	InitSystemTask()
}

func InitSystemTask() {
	crontab := cron.New()
	task := func() {
		config.LOG.Debugf("====================>>>>")
		GetCpuInfo()
		GetCpuLoad()
		GetMemInfo()
		GetHostInfo()
		GetDiskInfo()
		GetNetInfo()
		config.LOG.Debugf("====================>>>>")
	}

	// 添加定时任务，表示每分钟执行一次
	// crontab.AddFunc("0 */1 * * * ?", task)
	crontab.AddFunc("@every 1m", task)
	// 启动定时器
	crontab.Start()
	// 定时任务是另起协程执行的,这里使用 select 简答阻塞.实际开发中需要
	// 根据实际情况进行控制
	// select {}
}
