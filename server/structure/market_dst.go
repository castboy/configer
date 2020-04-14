package structure

import "configer/server/structure/indexNameID"

type MarketDST struct {
	MarketOwnerType indexNameID.MarketType `json:"market_owner_type" xorm:"market_type"`
	DST             DSTType                `xorm:"dst_type"`
}

func (md *MarketDST) FormatCheck() error {
	return nil
}
func (md *MarketDST) IndexCheck() error {
	return nil
}

func (md *MarketDST) AutoCondition() (cond string) {
	return
}
