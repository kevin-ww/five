package comm

type (
	Payload struct {
		TxId   string
		TxTime int64
		Memo   string
	}

	//payload interface {
	//	IsValid() bool
	//	GenKey() string
	//}
)
