package args

const (
	ShopInfoSearchNoticeType = 10001
)

const (
	ShopInfoSearchNoticeTag    = "shop_info_search_notice"
	ShopInfoSearchNoticeTagErr = "shop_info_search_notice_err"
)

type SearchStoreShop struct {
	ShopId       int64  `json:"shop_id,omitempty"`
	NickName     string `json:"nick_name,omitempty"`
	FullName     string `json:"full_name,omitempty"`
	ShopCode     string `json:"shop_code,omitempty"`
	RegisterAddr string `json:"register_addr,omitempty"`
	BusinessAddr string `json:"business_addr,omitempty"`
	BusinessDesc string `json:"business_desc,omitempty"`
}

type CommonBusinessMsg struct {
	Type    int    `json:"type"`
	Tag     string `json:"tag"`
	UUID    string `json:"uuid"`
	Content string `json:"content"`
}
