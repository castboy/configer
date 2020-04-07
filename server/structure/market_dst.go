package structure

type MarketDST struct {
	MarketOwnerType MarketType `json:"market_owner_type" xorm:"market_type"`
	DST             DSTType    `xorm:"dst_type"`
}

func (md *MarketDST) FormatCheck() error {
	return nil
}
