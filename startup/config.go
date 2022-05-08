package startup

import (
	"gitee.com/cristiane/micro-mall-shop-cron/vars"
	"gitee.com/kelvins-io/kelvins/config"
	"gitee.com/kelvins-io/kelvins/config/setting"
)

const (
	SectionEmailConfig       = "email-config"
	ShopInfoSearchNotice     = "shop-info-search-notice"
	ShopInfoSearchNoticeTask = "shop-info-search-notice-task"
)

// LoadConfig 加载配置对象映射
func LoadConfig() error {
	// 邮箱
	vars.EmailConfigSetting = new(vars.EmailConfigSettingS)
	config.MapConfig(SectionEmailConfig, vars.EmailConfigSetting)

	vars.QueueAMQPSettingShopInfoSearchNotice = new(setting.QueueAMQPSettingS)
	config.MapConfig(ShopInfoSearchNotice, vars.QueueAMQPSettingShopInfoSearchNotice)

	vars.ShopInfoSearchSyncTaskSetting = new(vars.ShopInfoSearchSyncTaskSettingS)
	config.MapConfig(ShopInfoSearchNoticeTask, vars.ShopInfoSearchSyncTaskSetting)
	return nil
}
