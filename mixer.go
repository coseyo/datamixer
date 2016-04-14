package datamixer

import (
	"errors"
	"fmt"
	"math"
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
	//	var wg sync.WaitGroup
	pdts := []processDataType{}

	for _, sd := range datas {
		//		go func() {
		//			wg.Add(1)
		//			defer wg.Done()

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
		//		}()
	}

	//	wg.Wait()

	if err != nil {
		return
	}

	if len(pdts) == 0 {
		return
	}

	return m.mixResp(pdts)
}

func (m *Mixer) mixResp(pdts []processDataType) (retResp DataResponse, err error) {

	var (
		limit int64
	)

	totalWeight := m.getTotalWeight(pdts)

	fmt.Println("totalWeight")
	fmt.Println(totalWeight)

	limitMap := m.getRealLimitMap(pdts, totalWeight)

	fmt.Println("limitMap")
	fmt.Println(limitMap)

	for _, pdt := range pdts {
		limit = limitMap[pdt.Name]
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

func (m *Mixer) getRealLimitMap(pdts []processDataType, totalWeight int64) (limitMap map[string]int64) {

	pdtsLen := len(pdts)
	leftDataCountMap := make(map[string]int64, pdtsLen)
	limitMap = make(map[string]int64, pdtsLen)

	var (
		weightPercent float64
		theoryLimit,
		totalTheoryLimit,
		realLimit,
		needFillCount int64
	)

	for k, pdt := range pdts {

		weightPercent = float64(pdt.Weight) / float64(totalWeight)

		theoryLimit = int64(round(weightPercent * float64(m.GlobalLimit)))
		totalTheoryLimit += theoryLimit

		if k == pdtsLen-1 && totalTheoryLimit > m.GlobalLimit {
			theoryLimit = m.GlobalLimit - (totalTheoryLimit - theoryLimit)
		}

		fmt.Println("m.GlobalLimit", m.GlobalLimit)
		fmt.Println(pdt.Name, "pdt.Weight", pdt.Weight)
		fmt.Println(pdt.Name, "totalWeight", totalWeight)
		fmt.Println(pdt.Name, "weightPercent", weightPercent)
		fmt.Println(pdt.Name, "pdt.DataCount", pdt.DataCount)
		fmt.Println(pdt.Name, "theoryLimit", theoryLimit)

		if pdt.DataCount > theoryLimit {
			leftDataCountMap[pdt.Name] = pdt.DataCount - theoryLimit
			realLimit = theoryLimit
		} else {
			needFillCount += theoryLimit - pdt.DataCount
			realLimit = pdt.DataCount
		}

		limitMap[pdt.Name] = realLimit

	}

	if needFillCount > 0 && len(leftDataCountMap) > 0 {
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

func round(input float64) float64 {
	if input < 0 {
		return math.Ceil(input - 0.5)
	}
	return math.Floor(input + 0.5)
}
