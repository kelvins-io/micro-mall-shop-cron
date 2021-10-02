package startup

import (
	"gitee.com/cristiane/micro-mall-shop-cron/service"
	"gitee.com/cristiane/micro-mall-shop-cron/vars"
	"gitee.com/kelvins-io/kelvins"
)

func GenCronJobs() []*kelvins.CronJob {
	tasks := make([]*kelvins.CronJob, 0)
	if vars.ShopInfoSearchSyncTaskSetting != nil {
		if vars.ShopInfoSearchSyncTaskSetting.Cron != "" {
			tasks = append(tasks, &kelvins.CronJob{
				Name: "店铺信息-搜索同步",
				Spec: vars.ShopInfoSearchSyncTaskSetting.Cron,
				Job:  service.ShopInfoSearchSync,
			})
		}
	}

	return tasks
}
