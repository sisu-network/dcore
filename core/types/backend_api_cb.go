package types

type BackendAPICallback struct {
	// callback before a transaction is submitted to a mempool.
	OnTxSubmitted func(*Transaction)
}
