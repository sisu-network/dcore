package extra

import "strconv"

type EvmApi struct {
	index int
}

func (evm *EvmApi) Snapshot() string {
	evm.index++
	return strconv.FormatInt(int64(evm.index), 16)
}
