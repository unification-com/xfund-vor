package api

type SimValues struct {
	IfGas  uint64  `json:"if_gas"`
	IfFees float64 `json:"if_fees"`
}

type AnalyticsFilter struct {
	ConsumerContract string `json:"consumer_contract,omitempty"`
	Limit            int    `json:"limit,omitempty"`
}

type IntStats struct {
	Max  uint64 `json:"max"`
	Min  uint64 `json:"min"`
	Mean uint64 `json:"mean"`
}

type FloatStats struct {
	Max  float64 `json:"max"`
	Min  float64 `json:"min"`
	Mean float64 `json:"mean"`
}

type EarningsStats struct {
	CurrentXfundPriceEth float64 `json:"current_xfund_price_eth"`
	TotalFeesEarnedXfund float64 `json:"total_fees_xfund"`
	TotalFeesEarnedEth   float64 `json:"total_fees_eth"`
	TotalCostsEth        float64 `json:"total_cost_eth"`
	ProfitLossEth        float64 `json:"profit_loss_eth"`
}

type AnalyticsData struct {
	GasUsed              IntStats      `json:"gas_used"`
	GasPrice             IntStats      `json:"gas_price"`
	EthCosts             FloatStats    `json:"eth_costs"`
	Earnings             EarningsStats `json:"earnings"`
	MostGasUsedConsumer  string        `json:"most_gas_used_consumer,omitempty"`
	LeastGasUsedConsumer string        `json:"least_gas_used_consumer,omitempty"`
	NumberAnalysed       uint64        `json:"number_requests_analysed"`
}

type AnalyticsResponse struct {
	AnalyticsData
	Filters AnalyticsFilter `json:"filters"`
}

type AnalyticsSimResponse struct {
	AnalyticsData
	SimValues SimValues       `json:"simulation_values"`
	Filters   AnalyticsFilter `json:"filters"`
}

type ConsumerAnalytics struct {
	Consumer   string  `json:"consumer"`
	CurrentFee float64 `json:"current_fee"`
	AnalyticsData
}

type ConsumersResponse struct {
	Consumers []ConsumerAnalytics `json:"consumers"`
}
