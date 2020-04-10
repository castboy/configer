package structure

type Session struct {
	ID       int                `xorm:"id autoincr"`
	SourceID int                `xorm:"source_id"`
	Type     SessionType        `xorm:"type"`
	Dst      DSTType            `xorm:"dst_type"`
	Session  map[int32][]string `xorm:"session"`
}

type SessionType int

const (
	Quote SessionType = iota
	Trade
)

type DSTType int

const (
	DST DSTType = iota
	DSTNone
	None
)

func (se *Session) FormatCheck() error {
	return nil
}

func (se *Session) AutoCondition() (cond string) {
	return
}
