package repository

import (
	"gitee.com/cristiane/micro-mall-shop-cron/model/mysql"
	"gitee.com/kelvins-io/kelvins"
)

func ListShopInfo(sqlSelect string, pageSize, pageNum int) ([]mysql.ShopBusinessInfo, error) {
	var result = make([]mysql.ShopBusinessInfo, 0)
	err := kelvins.XORM_DBEngine.Table(mysql.TableShopBusinessInfo).Select(sqlSelect).Desc("update_time").
		Limit(pageSize, (pageNum-1)*pageSize).Find(&result)
	return result, err
}
