package strategies

//go:generate mockgen -destination=./mocks/iaccount_strategy_manager.go -package=mocks --build_flags=--mod=mod . IAccountStrategyManager
type IAccountStrategyManager interface {
	GetStrategy(topic string) IAccountStrategy
}
