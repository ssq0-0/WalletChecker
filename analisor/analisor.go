package analisor

import (
	"checkers/logger"
	"checkers/utils"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

type AddressData struct {
	Address common.Address
	Amount  float64
}

type SummaryData struct {
	Total   float64
	Average float64
}

type Aggregator struct {
	mu   sync.Mutex
	data map[common.Address]float64
}

func NewAggreagtor() *Aggregator {
	return &Aggregator{
		data: make(map[common.Address]float64),
	}
}

func (a *Aggregator) Add(addr common.Address, amount float64) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.data[addr] += amount
}

func (a *Aggregator) GetAggregatedData() []AddressData {
	a.mu.Lock()
	defer a.mu.Unlock()

	var ad []AddressData
	for addr, amount := range a.data {
		ad = append(ad, AddressData{
			Address: addr,
			Amount:  amount,
		})
	}

	return ad
}

func (a *Aggregator) GetAvgData() SummaryData {
	var total, avg float64
	for _, amount := range a.data {
		total += amount
	}

	avg = (total / float64(len(a.data)))
	return SummaryData{
		Total:   total,
		Average: avg,
	}
}

func (a *Aggregator) LogAnalizedData() {
	for wallet, amount := range a.data {
		logger.GlobalLogger.Infof("| Адрес: %s | Количество: %f|", wallet, amount)
	}

	totals := a.GetAvgData()
	logger.GlobalLogger.Infof("| Итого: %f | Среднее количесвто: %f|", totals.Total, totals.Average)
}

func (a *Aggregator) WriteAnalizedData() error {
	if err := utils.FileWriter("account/result.csv", a.data); err != nil {
		return err
	}

	return nil
}
