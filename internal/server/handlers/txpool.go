package handlers

const txPoolRPCName = "txpool"

type TxPool struct {
	name string
}

func NewTxPool() *TxPool {
	return &TxPool{
		name: txPoolRPCName,
	}
}

func (t *TxPool) Name() string {
	return t.name
}

func (t *TxPool) Content() (any, error) {
	return nil, nil //nolint:nilnil // empty implementation
}
