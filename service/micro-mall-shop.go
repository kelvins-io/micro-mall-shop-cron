package service

import (
	"context"
	"gitee.com/cristiane/micro-mall-shop-cron/model/args"
	"gitee.com/cristiane/micro-mall-shop-cron/repository"
	"gitee.com/cristiane/micro-mall-shop-cron/vars"
	"gitee.com/kelvins-io/common/json"
	"gitee.com/kelvins-io/kelvins"
	"github.com/google/uuid"
)

var (
	pageSize = 50
	pageNum  = 1
)

func ShopInfoSearchSync() {
	if vars.ShopInfoSearchSyncTaskSetting != nil {
		if vars.ShopInfoSearchSyncTaskSetting.SingleSyncNum > 0 {
			pageSize = vars.ShopInfoSearchSyncTaskSetting.SingleSyncNum
		}
	}
	count := 0
	for {
		if count > 5 {
			break
		}
		count++
		shopInfoSearchSync(pageSize, pageNum)
		pageNum++
	}
}

const sqlSelectShopInfo = "shop_id,nick_name,full_name,shop_code,register_addr,business_addr,business_desc"

func shopInfoSearchSync(pageSize, pageNum int) {
	ctx := context.TODO()
	list, err := repository.ListShopInfo(sqlSelectShopInfo, pageSize, pageNum)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "ListShopInfo err； %v", err)
		return
	}
	if len(list) == 0 {
		return
	}
	for i := 0; i < len(list); i++ {
		info := &args.SearchStoreShop{
			ShopId:       list[i].ShopId,
			NickName:     list[i].NickName,
			FullName:     list[i].FullName,
			ShopCode:     list[i].ShopCode,
			RegisterAddr: list[i].RegisterAddr,
			BusinessAddr: list[i].BusinessAddr,
			BusinessDesc: list[i].BusinessDesc,
		}
		shopInfoNoticeSearchNotice(info)
	}
}

func shopInfoNoticeSearchNotice(info *args.SearchStoreShop) {
	var msg = &args.CommonBusinessMsg{
		Type:    args.ShopInfoSearchNoticeType,
		Tag:     "店铺搜索通知",
		UUID:    uuid.New().String(),
		Content: json.MarshalToStringNoError(info),
	}
	var ctx = context.TODO()
	vars.QueueServerShopInfoSearchPusher.PushMessage(ctx, msg)
	return
}
