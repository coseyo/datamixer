package datamixer

import (
	"errors"
	"fmt"
	"sync"
)

type Mixer struct {
	Datas []SourceData
}

func (m *Mixer) GetData() (dataRes DataResponse, err error) {
	var wg sync.WaitGroup
	resps := []DataResponse{}

	for _, sd := range m.Datas {
		go func() {
			wg.Add(1)
			defer wg.Done()

			sDataRes, sErr := sd.GetData(sd.Params, sd.Limit, sd.Offset)
			if sErr != nil {
				err = errors.New(fmt.Sprintf("%s %s", err.Error(), sErr.Error()))
			}
			resps = append(resps, sDataRes)
		}()
	}

	wg.Wait()

	if err != nil {
		return
	}

	return m.mixResp(resps)
}

func (m *Mixer) mixResp(resps []DataResponse) (dataRes DataResponse, err error) {
	return
}
