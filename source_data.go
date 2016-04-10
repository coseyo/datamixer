package datamixer

type GetDataFunc func(params string, Limit, Offset int64) (DataResponse, error)

type SourceData struct {
	Name   string
	Params string
	Limit  int64
	Offset int64
	Weight int64

	GetData GetDataFunc
}
