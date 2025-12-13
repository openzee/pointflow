package flow

import (
	"time"
)

func CreateBatchPoint() BatchPoint {
	arr := LoadExcel()

	batch := BatchPoint{}

	for _, a := range arr {

		batch = append(batch, &Point{
			Original:        a,
			Value:           1234,
			ChangeTimestamp: time.Now(),
		})
	}

	return batch
}
