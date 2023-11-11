package presentations

type CreateEntryParams struct {
	AccountID int64 `json:"account_id"`
	Amount    int64 `json:"amount"`
}

type ListEntriesParams struct {
	AccountID int64 `json:"account_id"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}
