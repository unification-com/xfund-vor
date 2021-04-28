package service

import (
	"github.com/ethereum/go-ethereum/params"
	"math/big"
	"oracle/models/api"
)

func (d *Service) Analytics(xFundEth, xFundUsd float64, limit int) (*api.AnalyticsResponse, error) {

	requests, err := d.Store.Db.GetLastXRequests(limit)

	if err != nil {
		return &api.AnalyticsResponse{}, err
	}

	if len(requests) == 0 {
		return &api.AnalyticsResponse{}, nil
	}

	numRows := uint64(len(requests))

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

	for _, reqRow := range requests {

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
		totalFees = big.NewFloat(0).Add(totalFees, new(big.Float).SetUint64(reqRow.Fee))
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

	analytics := &api.AnalyticsResponse{
		GasUsedMin:           gasMin.Uint64(),
		GasUsedMax:           gasMax.Uint64(),
		GasUsedMean:          gasMeanUint64,
		GasPriceMin:          gasPriceMinGweiUint64,
		GasPriceMax:          gasPriceMaxGweiUint64,
		GasPriceMean:         gasPriceMeanGweiUint64,
		CostMinEth:           cMin,
		CostMaxEth:           cMax,
		CostMeanEth:          cMean,
		TotalCostEth:         tCost,
		TotalFeesEarnedXfund: totalFeesXfund,
		TotalFeesEarnedEth:   totalFeesEthFloat64,
		ProfitLossEth:        profitLossFloat64,
		NumberAnalysed:       numRows,
	}

	return analytics, nil

}
