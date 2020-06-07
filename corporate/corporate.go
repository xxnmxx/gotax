package corporate

type Corp struct {
	Cit  Cit
	Vat  Vat
	PLs  []PL
	BSs  []BS
	Name string
}

type Acc interface {
	Name() string
	Value() float64
	Type() string
}

type Pl interface {
	Acc


type Cit struct {
	Incomes []Income
	Adj     float64
	Credit  float64
	TaxRate float64
}

type Vat struct {
	Values  []Value
	TaxRate float64
	LumpSum bool
}
