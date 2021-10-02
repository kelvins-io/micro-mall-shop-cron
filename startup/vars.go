package startup

import (
	"gitee.com/cristiane/micro-mall-shop-cron/model/args"
	"gitee.com/cristiane/micro-mall-shop-cron/vars"
	"gitee.com/kelvins-io/kelvins"
	"gitee.com/kelvins-io/kelvins/setup"
	"gitee.com/kelvins-io/kelvins/util/queue_helper"
)

// SetupVars 加载变量
func SetupVars() error {
	var err error
	if vars.QueueAMQPSettingShopInfoSearchNotice != nil && vars.QueueAMQPSettingShopInfoSearchNotice.Broker != "" {
		vars.QueueServerShopInfoSearch, err = setup.NewAMQPQueue(vars.QueueAMQPSettingShopInfoSearchNotice, nil)
		if err != nil {
			return err
		}
		vars.QueueServerShopInfoSearchPusher, err = queue_helper.NewPublishService(vars.QueueServerShopInfoSearch, &queue_helper.PushMsgTag{
			DeliveryTag:    args.ShopInfoSearchNoticeTag,
			DeliveryErrTag: args.ShopInfoSearchNoticeTagErr,
			RetryCount:     vars.QueueAMQPSettingShopInfoSearchNotice.TaskRetryCount,
			RetryTimeout:   vars.QueueAMQPSettingShopInfoSearchNotice.TaskRetryTimeout,
		}, kelvins.BusinessLogger)
		if err != nil {
			return err
		}
	}
	return err
}

func StopFunc() error {
	var err error
	// 当kelvins收到退出信号时将会调用
	// 本应用的资源回收通常在这里进行，kelvins自动加载的资源回收由框架进行
	return err
}
