package client

import (
	"encoding/json"

	"github.com/regner/albiondata-client/lib"
	"github.com/regner/albiondata-client/log"
)

type operationGoldMarketGetAverageInfo struct {
}

func (op operationGoldMarketGetAverageInfo) Process(state *albionState, uploader iuploader) {
	log.Debug("Got GoldMarketGetAverageInfo operation...")
}

type operationGoldMarketGetAverageInfoResponse struct {
	GoldPrices []int `mapstructure:"0"`
	TimeStamps []int `mapstructure:"1"`
}

func (op operationGoldMarketGetAverageInfoResponse) Process(state *albionState, uploader iuploader) {
	log.Debug("Got response to GoldMarketGetAverageInfo operation...")

	data, err := json.Marshal(lib.GoldPricesUpload{
		Prices:     op.GoldPrices,
		TimeStamps: op.TimeStamps,
	})

	if err != nil {
		log.Errorf("Error while marshalling payload for gold prices: %v", err)
		return
	}

	uploader.sendToIngest(data, lib.NatsGoldPricesIngest)
}
