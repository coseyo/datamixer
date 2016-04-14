package datamixer

type GetDataFunc func(params string, limit, offset int64) (DataResponse, error)

type SourceData struct {
	Name   string
	Params string
	Offset int64
	Weight int64

	GetData GetDataFunc
}
