package datamixer

type DataResponse struct {
	Data  []map[string]interface{}
	Total int64
}

type GetDataFunc func(params string, limit, offset int64) (DataResponse, error)

type SourceData struct {
	Name   string
	Params string
	Offset int64
	Weight int64

	GetData GetDataFunc
}
