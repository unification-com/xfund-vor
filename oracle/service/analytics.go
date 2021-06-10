package service

import (
	"github.com/ethereum/go-ethereum/params"
	"math/big"
	"oracle/models/api"
	"oracle/models/database"
)

func (d *Service) Analytics(xFundEth, xFundUsd, fees float64, limit int, gasPrice int64, simulation int, consumer string) (interface{}, error) {

	requests, err := d.Store.Db.GetLastXRequests(limit, consumer)

	if err != nil {
		return nil, err
	}

	if len(requests) == 0 {
		return nil, nil
	}

	mostGasUsedContract := ""
	leastGasUSedContract := ""

	mgu, err := d.Store.Db.GetMostGasUsed()
	if err == nil {
		mostGasUsedContract = mgu.Sender
	}
	lgu, err := d.Store.Db.GetLeastGasUsed()
	if err == nil {
		leastGasUSedContract = lgu.Sender
	}

	analyticsData := process(requests, xFundEth, xFundUsd, fees, gasPrice, simulation)

	if len(consumer) == 0 {
		analyticsData.MostGasUsedConsumer = mostGasUsedContract
		analyticsData.LeastGasUsedConsumer = leastGasUSedContract
	}

	filters := api.AnalyticsFilter{
		ConsumerContract: consumer,
		Limit:            limit,
	}

	if simulation == 1 {
		return &api.AnalyticsSimResponse{
			AnalyticsData: analyticsData,
			Filters:       filters,
			SimValues: api.SimValues{
				IfFees: fees,
				IfGas:  uint64(gasPrice),
			},
		}, nil
	} else {
		return &api.AnalyticsResponse{
			AnalyticsData: analyticsData,
			Filters:       filters,
		}, nil
	}
}

func process(rows []database.RandomnessRequest, xFundEth, xFundUsd, fees float64, gasPrice int64, simulation int) api.AnalyticsData {

	numRows := uint64(len(rows))

	gasMin := big.NewInt(0)
	gasMax := big.NewInt(0)
	gasMean := big.NewFloat(0)
	gasSum := big.NewInt(0)

	gasPriceMin := big.NewInt(0)
	gasPriceMax := big.NewInt(0)
	gasPriceMean := big.NewFloat(0)
	gasPriceSum := big.NewInt(0)

	costMin := big.NewFloat(0)
	costMax := big.NewFloat(0)
	costMean := big.NewFloat(0)
	costSum := big.NewFloat(0)

	totalFees := big.NewFloat(0)

	for _, reqRow := range rows {

		gasVal := int64(reqRow.FulfillGasUsed)
		// gas
		if gasMin.Cmp(big.NewInt(0)) == 0 || gasMin.Cmp(big.NewInt(gasVal)) > 0 {
			gasMin = big.NewInt(gasVal)
		}

		if gasMax.Cmp(big.NewInt(gasVal)) < 0 {
			gasMax = big.NewInt(gasVal)
		}

		gasSum = big.NewInt(0).Add(gasSum, big.NewInt(gasVal))

		// gas prices
		gasPriceVal := int64(reqRow.FulfillGasPrice)

		// check for simulating gas prices
		if simulation == 1 {
			gasPriceVal = gasPrice * 1e9
		}
		if gasPriceMin.Cmp(big.NewInt(0)) == 0 || gasPriceMin.Cmp(big.NewInt(gasPriceVal)) > 0 {
			gasPriceMin = big.NewInt(gasPriceVal)
		}

		if gasPriceMax.Cmp(big.NewInt(gasPriceVal)) < 0 {
			gasPriceMax = big.NewInt(gasPriceVal)
		}

		gasPriceSum = big.NewInt(0).Add(gasPriceSum, big.NewInt(gasPriceVal))

		// cost
		cost := new(big.Float).Mul(new(big.Float).SetInt64(gasVal), new(big.Float).SetInt64(gasPriceVal))

		if costMin.Cmp(big.NewFloat(0)) == 0 || costMin.Cmp(cost) > 0 {
			costMin = cost
		}

		if costMax.Cmp(cost) < 0 {
			costMax = cost
		}

		costSum = big.NewFloat(0).Add(costSum, cost)

		// fees
		reqFee := reqRow.Fee
		if simulation == 1 {
			reqFee = uint64(fees * 1e9)
		}
		totalFees = big.NewFloat(0).Add(totalFees, new(big.Float).SetUint64(reqFee))
	}

	// gas
	gasMean = new(big.Float).Quo(new(big.Float).SetInt(gasSum), new(big.Float).SetUint64(numRows))
	gasMeanUint64, _ := gasMean.Uint64()

	// gas price
	gasPriceMinGwei := new(big.Float).Quo(new(big.Float).SetInt(gasPriceMin), big.NewFloat(params.GWei))
	gasPriceMinGweiUint64, _ := gasPriceMinGwei.Uint64()

	gasPriceMaxGwei := new(big.Float).Quo(new(big.Float).SetInt(gasPriceMax), big.NewFloat(params.GWei))
	gasPriceMaxGweiUint64, _ := gasPriceMaxGwei.Uint64()

	gasPriceMean = new(big.Float).Quo(new(big.Float).SetInt(gasPriceSum), new(big.Float).SetUint64(numRows))
	gasPriceMeanGwei := new(big.Float).Quo(gasPriceMean, big.NewFloat(params.GWei))
	gasPriceMeanGweiUint64, _ := gasPriceMeanGwei.Uint64()

	// cost
	costMean = new(big.Float).Quo(costSum, new(big.Float).SetUint64(numRows))
	costMinEth := new(big.Float).Quo(costMin, big.NewFloat(params.Ether))
	costMaxEth := new(big.Float).Quo(costMax, big.NewFloat(params.Ether))
	costMeanEth := new(big.Float).Quo(costMean, big.NewFloat(params.Ether))
	totalCostEth := new(big.Float).Quo(costSum, big.NewFloat(params.Ether))

	cMin, _ := costMinEth.Float64()
	cMax, _ := costMaxEth.Float64()
	cMean, _ := costMeanEth.Float64()
	tCost, _ := totalCostEth.Float64()

	// fees
	totalFeesTokens := new(big.Float).Quo(totalFees, big.NewFloat(params.GWei))
	totalFeesXfund, _ := totalFeesTokens.Float64()

	totalFeesEth := new(big.Float).Mul(big.NewFloat(totalFeesXfund), big.NewFloat(xFundEth))
	totalFeesEthFloat64, _ := totalFeesEth.Float64()

	profitLoss := new(big.Float).Sub(totalFeesEth, totalCostEth)
	profitLossFloat64, _ := profitLoss.Float64()

	return api.AnalyticsData{
		GasUsed: api.IntStats{
			Min:  gasMin.Uint64(),
			Max:  gasMax.Uint64(),
			Mean: gasMeanUint64,
		},
		GasPrice: api.IntStats{
			Min:  gasPriceMinGweiUint64,
			Max:  gasPriceMaxGweiUint64,
			Mean: gasPriceMeanGweiUint64,
		},
		EthCosts: api.FloatStats{
			Min:  cMin,
			Max:  cMax,
			Mean: cMean,
		},
		Earnings: api.EarningsStats{
			CurrentXfundPriceEth: xFundEth,
			TotalFeesEarnedXfund: totalFeesXfund,
			TotalFeesEarnedEth:   totalFeesEthFloat64,
			TotalCostsEth:        tCost,
			ProfitLossEth:        profitLossFloat64,
		},
		NumberAnalysed: numRows,
	}
}
