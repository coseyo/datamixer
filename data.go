package datamixer

// DataResponse is a standard response struct
type DataResponse struct {
	Data  []map[string]interface{}
	Total int64
}

// GetDataFunc is a interface for get data, every source data could get data by
// this func
type GetDataFunc func(params string, limit, offset int64) (DataResponse, error)

type SourceData struct {
	Name   string
	Params string
	Offset int64
	Weight int64

	GetData GetDataFunc
}
