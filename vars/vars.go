package vars

import (
	"gitee.com/kelvins-io/common/queue"
	"gitee.com/kelvins-io/kelvins/config/setting"
	"gitee.com/kelvins-io/kelvins/util/queue_helper"
)

var (
	EmailNoticeSetting                   *EmailNoticeSettingS
	EmailConfigSetting                   *EmailConfigSettingS
	ShopInfoSearchSyncTaskSetting        *ShopInfoSearchSyncTaskSettingS
	QueueAMQPSettingShopInfoSearchNotice *setting.QueueAMQPSettingS
	QueueServerShopInfoSearch            *queue.MachineryQueue
	QueueServerShopInfoSearchPusher      *queue_helper.PublishService
)
