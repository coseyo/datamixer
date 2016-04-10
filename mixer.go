package datamixer

import (
	"sync"
)

type Mixer struct {
	Datas []SourceData
}

func (m *Mixer) Mix() (dataRes DataResponse, err error) {
	wg := sync.WaitGroup

	for _, sourceData := range m.Datas {
		go func() {
			wg.Add(1)
			defer wg.Done()

			sourceData.GetData()
		}()
	}

	wg.Wait()
}
