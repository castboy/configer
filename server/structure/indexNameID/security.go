package indexNameID

type Security struct {
	ID           int      `json:"id" xorm:"id"`
	SecurityName string   `json:"security_name" xorm:"security_name"`
	Description  string   `json:"description" xorm:"description"`
	Symbols      []string `json:"symbols" xorm:"-"`
}

func (sec *Security) FormatCheck() error {
	return nil
}

func (sec *Security) AutoCondition() (cond string) {
	return
}

func (sec *Security) GetName() string {
	return sec.SecurityName
}

func (sec *Security) GetID() int {
	return sec.ID
}
