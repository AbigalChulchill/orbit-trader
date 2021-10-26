package coinmonitor

func (s *Monitor) macd() float64 { //return difference between 12-26 moving average and 9 moving average of 12-26 moving average. negative indicates sell positive indicates buy.
	if assets, err := s.krakenapi.Assets(); err != nil {
		panic(err)
	}
	fmt.Println(assets)
}