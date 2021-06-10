package service

import (
	"github.com/ethereum/go-ethereum/params"
	"math/big"
	"oracle/models/api"
)

func (d *Service) Consumers(xFundEth, xFundUsd float64, consumer string) (interface{}, error) {

	consumers, err := d.Store.Db.GetDistinctConsumers(consumer)

	if err != nil {
		return nil, err
	}

	if len(consumers) == 0 {
		return nil, nil
	}

	consumerResponse := api.ConsumersResponse{}

	for _, consumer := range consumers {
		var analyticsData api.AnalyticsData
		currentXfundFee := 0.0
		consumerRows, err := d.Store.Db.GetLastXRequests(0, consumer.Sender)

		if err == nil {
			analyticsData = process(consumerRows, xFundEth, xFundUsd, 0, 0, 0)
		}

		currentFee, err := d.VORCoordinatorCaller.QueryFees(consumer.Sender)
		currentFeeTokens := new(big.Float).Quo(new(big.Float).SetInt(currentFee), big.NewFloat(params.GWei))
		currentXfundFee, _ = currentFeeTokens.Float64()

		consumerAnalytics := api.ConsumerAnalytics{
			Consumer:      consumer.Sender,
			CurrentFee:    currentXfundFee,
			AnalyticsData: analyticsData,
		}

		consumerResponse.Consumers = append(consumerResponse.Consumers, consumerAnalytics)
	}

	if len(consumer) > 0 {
		if len(consumerResponse.Consumers) > 0 {
			return consumerResponse.Consumers[0], nil
		}
	}
	return consumerResponse, nil
}
