package datamixer

import (
	"fmt"
	"testing"
)

func getData1(params string, limit, offset int64) (DataResponse, error) {
	return DataResponse{
		Data: []map[string]interface{}{
			map[string]interface{}{
				"a1": "a1",
				"a2": "a2",
			},
			map[string]interface{}{
				"a1": "a1",
				"a2": "a2",
			},
			map[string]interface{}{
				"a1": "a1",
				"a2": "a2",
			},
			map[string]interface{}{
				"a1": "a1",
				"a2": "a2",
			},
			map[string]interface{}{
				"a1": "a1",
				"a2": "a2",
			},
		},
		Total: 10,
	}, nil
}

func getData2(params string, limit, offset int64) (DataResponse, error) {
	return DataResponse{
		Data: []map[string]interface{}{
			map[string]interface{}{
				"b1": "b1",
				"b2": "b2",
			},
			map[string]interface{}{
				"b1": "b1",
				"b2": "b2",
			},
			map[string]interface{}{
				"b1": "b1",
				"b2": "b2",
			},
			map[string]interface{}{
				"b1": "b1",
				"b2": "b2",
			},
			map[string]interface{}{
				"b1": "b1",
				"b2": "b2",
			},
		},
		Total: 10,
	}, nil
}

func Test_Mixer1(t *testing.T) {
	datas := []SourceData{
		SourceData{
			Name:    "data1",
			Params:  "",
			Offset:  0,
			Weight:  5,
			GetData: getData1,
		},
		SourceData{
			Name:    "data2",
			Params:  "",
			Offset:  0,
			Weight:  5,
			GetData: getData2,
		},
	}

	m := &Mixer{
		GlobalLimit: 5,
	}

	data1, _ := getData2("", 0, 0)
	fmt.Println(data1)

	d, err := m.Mix(datas)

	fmt.Println(err)
	fmt.Println(d)
}
