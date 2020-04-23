package main

import (
	"configer/server/structure"
	"github.com/shopspring/decimal"
)

const (
	SystemLeverage = 100
	DaysInYear     = 360
)

var SymbolLeverage float64

func init() {
	// TODO
	//cnf := config.GetConfigService("common")
	//SymbolLeverage = cnf.GetFloat64("symbol_leverage")
	//if SymbolLeverage == 0 {
	//	SymbolLeverage = 100
	//	sentry.CaptureMessage("symbol leverage is zero or null, set 100 default")
	//}
}

func MarginFormula(symb *ExportSymbol) func(lots, marketPrice decimal.Decimal) decimal.Decimal {
	switch symb.MarginMode {
	case structure.MarginForex:
		return func(lots, marketPrice decimal.Decimal) decimal.Decimal {
			// TODO : symbol.Leverage
			return lots.Mul(symb.ContractSize).Div(decimal.NewFromFloat(SymbolLeverage)).Mul(symb.Percentage).Div(decimal.NewFromFloat(SystemLeverage))
		}

	case structure.MarginCfd:
		return func(lots, marketPrice decimal.Decimal) decimal.Decimal {
			return lots.Mul(symb.ContractSize).Mul(marketPrice).Mul(symb.Percentage).Div(decimal.NewFromFloat(SystemLeverage))
		}

	case structure.MarginFutures:
		return func(lots, marketPrice decimal.Decimal) decimal.Decimal {
			return lots.Mul(symb.MarginInitial).Mul(symb.Percentage).Div(decimal.NewFromFloat(SystemLeverage))
		}

	case structure.MarginCfdIndex:
		// TODO
		// need update SourceFormatCheck() when symbol support this margin mode.

	case structure.MarginCfdLeverage:
		return func(lots, marketPrice decimal.Decimal) decimal.Decimal {
			// TODO : symbol.Leverage
			return lots.Mul(symb.ContractSize).Mul(marketPrice).Div(decimal.NewFromFloat(SymbolLeverage)).Mul(symb.Percentage).Div(decimal.NewFromFloat(SystemLeverage))
		}
	}

	// sentry and log it, when margin mode is wrong.
	//err := errors.NotValidf("margin mode: %d, symbol: %s", symb.MarginMode, symb.Symbol.Symbol)
	//sentry.CaptureException(err)
	//SymbolLog.Error(err)

	// `MarginCfdLeverage` is used frequently, use it as an substitution when margin mode is wrong.
	return func(lots, marketPrice decimal.Decimal) decimal.Decimal {
		return lots.Mul(symb.ContractSize).Mul(marketPrice).Div(decimal.NewFromFloat(SymbolLeverage)).Mul(symb.Percentage).Div(decimal.NewFromFloat(SystemLeverage))
	}
}

func ProfitFormula(symb *ExportSymbol) func(lots, openPrice, closePrice decimal.Decimal) decimal.Decimal {
	switch symb.ProfitMode {
	case structure.ProfitForex, structure.ProfitCfd:
		return func(lots, openPrice, closePrice decimal.Decimal) decimal.Decimal {
			return closePrice.Sub(openPrice).Mul(symb.ContractSize).Mul(lots)
		}

	case structure.ProfitFutures:
		// TODO
		// need update SourceFormatCheck() when symbol support this profit mode.
	}

	// sentry and log it, when profit mode is wrong.
	//err := constant.NewErr(constant.ArgsErr, errors.NotValidf("profit mode: %d, symbol: %s", symb.ProfitMode, symb.Symbol.Symbol))
	//sentry.CaptureException(err)
	//SymbolLog.Error(err)

	// `ProfitForex, ProfitCfd` is used frequently, use it as an substitution when profit mode is wrong.
	return func(lots, openPrice, closePrice decimal.Decimal) decimal.Decimal {
		return closePrice.Sub(openPrice).Mul(symb.ContractSize).Mul(lots)
	}
}

func SwapFormula(symb *ExportSymbol) func(lots, longOrShort decimal.Decimal, price ...decimal.Decimal) decimal.Decimal {
	switch symb.SwapType {
	case structure.ByPoints:
		return func(lots, longOrShort decimal.Decimal, price ...decimal.Decimal) decimal.Decimal {
			divider := decimal.New(1, int32(-symb.Digits))
			return lots.Mul(longOrShort).Mul(symb.ContractSize).Mul(divider)
		}

	case structure.ByMoney:
		// TODO
		// need update SourceFormatCheck() when symbol support this swap type.

	case structure.ByInterest:
		// TODO
		// need update SourceFormatCheck() when symbol support this swap type.

		// not support currently, comment below.
		//return func(lots, longOrShort decimal.Decimal, price ...decimal.Decimal) decimal.Decimal {
		//	return lots.Mul(longOrShort).Mul(symb.ContractSize).Div(decimal.NewFromFloat(SystemLeverage)).Div(decimal.NewFromFloat(DaysInYear))
		//}

	case structure.ByMoneyInMarginCurrency:
		// TODO
		// need update SourceFormatCheck() when symbol support this swap type.

	case structure.ByInterestOfCfds:
		return func(lots, longOrShort decimal.Decimal, price ...decimal.Decimal) decimal.Decimal {
			return lots.Mul(longOrShort).Mul(symb.ContractSize).Mul(price[0]).Div(decimal.NewFromFloat(SystemLeverage)).Div(decimal.NewFromFloat(DaysInYear))
		}

	case structure.ByInterestOfFutures:
		// TODO
		// need update SourceFormatCheck() when symbol support this swap type.
	}

	// sentry and log it, when swap type is wrong.
	//err := constant.NewErr(constant.ArgsErr, errors.NotValidf("swap type: %d, symbol: %s", symb.SwapType, symb.Symbol.Symbol))
	//sentry.CaptureException(err)
	//SymbolLog.Error(err)

	// `ByPoints` is used frequently, use it as an substitution when swap type is wrong.
	return func(lots, longOrShort decimal.Decimal, price ...decimal.Decimal) decimal.Decimal {
		divider := decimal.New(1, int32(-symb.Digits))
		return lots.Mul(longOrShort).Mul(symb.ContractSize).Mul(divider)
	}
}

