package datamixer

import (
	"fmt"
	"testing"
)

func getData1(params string, limit, offset int64) (DataResponse, error) {
	d := []map[string]interface{}{
		map[string]interface{}{
			"a1": "a0",
			"a2": "a0",
		},
		map[string]interface{}{
			"a1": "a1",
			"a2": "a1",
		},
		map[string]interface{}{
			"a1": "a2",
			"a2": "a2",
		},
		map[string]interface{}{
			"a1": "a3",
			"a2": "a3",
		},
		map[string]interface{}{
			"a1": "a4",
			"a2": "a4",
		},
		map[string]interface{}{
			"a1": "a5",
			"a2": "a5",
		},
		map[string]interface{}{
			"a1": "a6",
			"a2": "a6",
		},
		map[string]interface{}{
			"a1": "a7",
			"a2": "a7",
		},
		map[string]interface{}{
			"a1": "a8",
			"a2": "a8",
		},
	}

	total := int64(len(d))
	lastOffset := offset + limit

	if total == 0 {
		return DataResponse{
			Data:  []map[string]interface{}{},
			Total: 0,
		}, nil
	}

	if offset > total {
		return DataResponse{
			Data:  []map[string]interface{}{},
			Total: total,
		}, nil
	}

	if lastOffset > total {
		lastOffset = total
	}

	return DataResponse{
		Data:  d[offset:lastOffset],
		Total: total,
	}, nil
}

func getData2(params string, limit, offset int64) (DataResponse, error) {
	d := []map[string]interface{}{
		map[string]interface{}{
			"b1": "b0",
			"b2": "b0",
		},
		map[string]interface{}{
			"b1": "b1",
			"b2": "b1",
		},
		map[string]interface{}{
			"b1": "b2",
			"b2": "b2",
		},
		map[string]interface{}{
			"b1": "b3",
			"b2": "b3",
		},
		map[string]interface{}{
			"b1": "b4",
			"b2": "b4",
		},
	}

	total := int64(len(d))
	lastOffset := offset + limit

	if total == 0 {
		return DataResponse{
			Data:  []map[string]interface{}{},
			Total: 0,
		}, nil
	}

	if offset > total {
		return DataResponse{
			Data:  []map[string]interface{}{},
			Total: total,
		}, nil
	}

	if lastOffset > total {
		lastOffset = total
	}

	return DataResponse{
		Data:  d[offset:lastOffset],
		Total: total,
	}, nil
}

func getData3(params string, limit, offset int64) (DataResponse, error) {
	d := []map[string]interface{}{
		map[string]interface{}{
			"c1": "c0",
			"c2": "c0",
		},
		map[string]interface{}{
			"c1": "c1",
			"c2": "c1",
		},
		map[string]interface{}{
			"c1": "c2",
			"c2": "c2",
		},
		map[string]interface{}{
			"c1": "c3",
			"c2": "c3",
		},
	}

	total := int64(len(d))
	lastOffset := offset + limit

	if total == 0 {
		return DataResponse{
			Data:  []map[string]interface{}{},
			Total: 0,
		}, nil
	}

	if offset > total {
		return DataResponse{
			Data:  []map[string]interface{}{},
			Total: total,
		}, nil
	}

	if lastOffset > total {
		lastOffset = total
	}

	return DataResponse{
		Data:  d[offset:lastOffset],
		Total: total,
	}, nil
}

func Test_Mixer1(t *testing.T) {
	datas := []SourceData{
		SourceData{
			Name:    "data1",
			Params:  "",
			Offset:  6,
			Weight:  4,
			GetData: getData1,
		},
		SourceData{
			Name:    "data2",
			Params:  "",
			Offset:  5,
			Weight:  3,
			GetData: getData2,
		},
		SourceData{
			Name:    "data3",
			Params:  "",
			Offset:  4,
			Weight:  3,
			GetData: getData3,
		},
	}

	m := &Mixer{
		GlobalLimit: 5,
	}

	d, offsetMap, err := m.Mix(datas)

	fmt.Println("err", err)
	fmt.Println("offsetMap", offsetMap)
	fmt.Println("resp", d)
}
