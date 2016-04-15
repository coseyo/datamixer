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
				"a1": "a3",
				"a2": "a4",
			},
			map[string]interface{}{
				"a1": "a5",
				"a2": "a6",
			},
			map[string]interface{}{
				"a1": "a7",
				"a2": "a8",
			},
			map[string]interface{}{
				"a1": "a9",
				"a2": "a10",
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
				"b1": "b3",
				"b2": "b4",
			},
			map[string]interface{}{
				"b1": "b5",
				"b2": "b6",
			},
			map[string]interface{}{
				"b1": "b7",
				"b2": "b8",
			},
			//			map[string]interface{}{
			//				"b1": "b1",
			//				"b2": "b2",
			//			},
		},
		Total: 4,
	}, nil
}

func Test_Mixer1(t *testing.T) {
	datas := []SourceData{
		SourceData{
			Name:    "data1",
			Params:  "",
			Offset:  3,
			Weight:  5,
			GetData: getData1,
		},
		SourceData{
			Name:    "data2",
			Params:  "",
			Offset:  3,
			Weight:  5,
			GetData: getData2,
		},
	}

	m := &Mixer{
		GlobalLimit: 7,
	}

	data1, _ := getData2("", 0, 0)
	fmt.Println(data1)

	d, offsetMap, err := m.Mix(datas)

	fmt.Println(err)
	fmt.Println(offsetMap)
	fmt.Println(d)
}
