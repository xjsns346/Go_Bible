// 进行磅和千克的转换
package weightconv

type Pound float64
type Kg float64

func PtoK(p Pound) Kg { return Kg(p * 0.45) }

func KtoP(k Kg) Pound { return Pound(k * 2.2) }
