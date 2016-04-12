package datamixer

import (
	"errors"
	"fmt"
	"sync"
)

type processDataType struct {
	Name      string
	Resp      DataResponse
	DataCount int64
	Weight    int64
}

type Mixer struct {
	GlobalLimit int64
}

func (m *Mixer) Mix(datas []SourceData) (dataRes DataResponse, err error) {
	var wg sync.WaitGroup
	pdts := []processDataType{}

	for _, sd := range datas {
		go func() {
			wg.Add(1)
			defer wg.Done()

			sDataRes, sErr := sd.GetData(sd.Params, m.GlobalLimit, sd.Offset)
			if sErr != nil {
				err = errors.New(fmt.Sprintf("%s %s:%s", err.Error(), sd.Name, sErr.Error()))
			}

			dataLen := int64(len(sDataRes.Data))
			if dataLen > 0 {
				pdts = append(pdts, processDataType{
					Name:      sd.Name,
					Resp:      sDataRes,
					DataCount: dataLen,
					Weight:    sd.Weight,
				})
			}
		}()
	}

	wg.Wait()

	if err != nil {
		return
	}

	return m.mixResp(pdts)
}

func (m *Mixer) mixResp(pdts []processDataType) (retResp DataResponse, err error) {

	var (
		limit,
		leftDataCount,
		dataCount int64
	)

	totalWeight := m.getTotalWeight(pdts)

	for _, pdt := range pdts {

		limit = (pdt.Weight / totalWeight) * m.GlobalLimit

		dataCount = pdt.DataCount

		if dataCount < limit {
			leftDataCount = limit - dataCount

			limit = dataCount
		}

		retResp.Data = append(retResp.Data, pdt.Resp.Data[:limit]...)
		retResp.Total += pdt.Resp.Total
	}

	return
}

func (m *Mixer) getTotalWeight(pdts []processDataType) (totalWeight int64) {
	for _, pdt := range pdts {
		totalWeight += pdt.Weight
	}
	return
}

func (m *Mixer) getRealLimit(pdts []processDataType, totalWeight int64) (limitMap map[string]int64) {

	pdtsLen := len(pdts)
	leftDataCountMap := make(map[string]int64, pdtsLen)
	var (
		weightPercent,
		theoryLimit,
		realLimit,
		needFillCount int64
	)

	for _, pdt := range pdts {

		weightPercent = pdt.Weight / totalWeight
		theoryLimit = weightPercent * m.GlobalLimit

		if pdt.DataCount > theoryLimit {
			leftDataCountMap[pdt.Name] = pdt.DataCount - theoryLimit
			realLimit = theoryLimit
		} else {
			needFillCount += theoryLimit - pdt.DataCount
			realLimit = pdt.DataCount
		}

		limitMap[pdt.Name] = realLimit
	}

	if needFillCount > 0 && len(leftDataCount) > 0 {
		for name, leftDataCount := range leftDataCountMap {
			if needFillCount <= 0 {
				break
			}

			if needFillCount > leftDataCount {
				limitMap[name] += leftDataCount
				needFillCount = needFillCount - leftDataCount
			} else {
				limitMap[name] += needFillCount
				needFillCount = 0
			}
		}
	}

	return
}
