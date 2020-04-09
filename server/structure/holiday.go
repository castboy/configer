package structure

type Holiday struct {
	ID          int  `json:"id" xorm:"id"`
	Enable      bool `json:"enable" xorm:"enable"`
	DateSymbol  `xorm:"extends"`
	From        string          `json:"from" xorm:"from"`
	To          string          `json:"to" xorm:"to"`
	Category    HolidayCategory `json:"category" xorm:"category"`
	Description string          `json:"description" xorm:"description"`

	HolidayTime []HolidayTime `xorm:"-"`
}

type DateSymbol struct {
	Date   string `json:"date" xorm:"date"`
	Symbol string `json:"symbol" xorm:"symbol"`
}

type HolidayCategory int

const (
	HolidayAll HolidayCategory = iota
	HolidaySecurity
	HolidaySymbol
	HolidaySource
)

type HolidayTime struct {
	From,
	To string
}

func (ho *Holiday) FormatCheck() error {
	return nil
}