package strategies

type IAccountStrategyManager interface {
	GetStrategy(topic string) IAccountStrategy
}
