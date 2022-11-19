package strategies

type AccountStrategyManager struct {
	Strategies map[string]IAccountStrategy
}

func (sm AccountStrategyManager) GetStrategy(topic string) IAccountStrategy {
	return sm.Strategies[topic]
}
