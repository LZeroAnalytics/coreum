package main

import (
	"encoding/binary"
	"fmt"
	"math/big"
	"runtime"

	sdkmath "cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

const precision = 19

func main() {
	n1 := decimal.RequireFromString("2")
	n2 := decimal.RequireFromString("3")

	m := memAlloc()
	div := n1.DivRound(n2, precision)
	m()
	divStr := div.String()
	m()

	fmt.Println(divStr)

	fmt.Println("\n1 ================================\n")
	maxUint64Str := "18446744073709551615"
	m = memAlloc()
	n := decimal.RequireFromString(maxUint64Str)
	m()
	nStr := n.String()
	m()
	fmt.Println(nStr)

	fmt.Println("\n2 ================================\n")
	m = memAlloc()
	nPow2 := n.Mul(n)
	m()
	nStr = nPow2.String()
	m()
	fmt.Println(nStr)

	fmt.Println("\n3 ================================\n")
	maxUint := maxSDKUint()
	fmt.Println(maxUint)
	// maxUint.Add(sdkmath.NewUint(1)) // this should panic and really does

	m = memAlloc()
	maxDec := decimal.NewFromBigInt(maxUint.BigInt(), 0)
	m()
	maxDecStr := maxDec.String()
	m()
	fmt.Println(maxDecStr)

	fmt.Println("\n4 ================================\n")
	one := decimal.NewFromInt(1)

	m = memAlloc()
	price1 := maxDec.DivRound(one, precision)
	m()
	fmt.Println(price1)

	m = memAlloc()
	price2 := one.DivRound(maxDec, precision)
	m()
	fmt.Println(price2)

	fmt.Println("\n5 ================================\n")
	res := maxDec
	n10 := decimal.NewFromInt(10)
	n0 := decimal.Decimal{}
	for !res.Equals(n0) {
		res = res.DivRound(n10, precision)
		fmt.Println(res)
	}

	fmt.Println("\n6 ================================\n")
	sdkN1 := sdkmath.NewUint(math.MaxUint64)
	sdkN2 := sdkmath.NewUint(7)
	price, err := calculatePrice(sdkN1, sdkN2)
	if err != nil {
		panic(err)
	}
	fmt.Println(price)
	fmt.Println(priceToStoreKey(price))

	fmt.Println("\n7 ================================\n")

	sdkN1 = sdkmath.NewUint(100_000)
	sdkN2 = sdkmath.NewUint(1_000)
	priceSum := decimal.Decimal{}
	var count int64 = 0
	for {
		price, err = calculatePrice(sdkN1, sdkN2)
		if err != nil {
			break
		}
		priceSum = priceSum.Add(price)
		count++
		fmt.Println(price)

		sdkN1 = sdkN1.Sub(sdkmath.OneUint())
		sdkN2 = sdkN2.Sub(sdkmath.OneUint())
	}
	fmt.Printf("Average price: %s\n", priceSum.DivRound(decimal.NewFromInt(count), precision))
}

func memAlloc() func() {
	var m runtime.MemStats
	var prev uint64
	var prevTotal uint64
	f := func() {
		runtime.ReadMemStats(&m)
		fmt.Printf("Mem usage: %d, total: %d\n", m.HeapAlloc-prev, m.TotalAlloc-prevTotal)
		prev = m.HeapAlloc
	}

	runtime.ReadMemStats(&m)
	prev = m.HeapAlloc
	prevTotal = m.TotalAlloc
	return f
}

func maxSDKUint() sdkmath.Uint {
	n1 := new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil)
	return sdkmath.NewUintFromBigInt(n1.Sub(n1, big.NewInt(1)))
}

var minInvalidPrice = decimal.RequireFromString("10000000000000000000")

func calculatePrice(n1 sdkmath.Uint, n2 sdkmath.Uint) (decimal.Decimal, error) {
	if n1.Equal(sdkmath.ZeroUint()) || n2.Equal(sdkmath.ZeroUint()) {
		return decimal.Decimal{}, errors.New("zeros are bad")
	}

	n1Dec := decimal.NewFromBigInt(n1.BigInt(), 0)
	n2Dec := decimal.NewFromBigInt(n2.BigInt(), 0)

	res := n1Dec.DivRound(n2Dec, precision)
	if res.Equals(decimal.Decimal{}) {
		return decimal.Decimal{}, errors.New("price is too low")
	}
	if res.GreaterThanOrEqual(minInvalidPrice) {
		return decimal.Decimal{}, errors.New("price is too big")
	}

	return res, nil
}

var precisionMul = decimal.NewFromInt(10).Pow(decimal.NewFromInt(precision))

func priceToStoreKey(price decimal.Decimal) []byte {
	wholePart := price.Floor()
	decPart := price.Sub(wholePart).Mul(precisionMul)

	b := make([]byte, 16)
	binary.BigEndian.PutUint64(b, uint64(wholePart.IntPart()))
	binary.BigEndian.PutUint64(b[8:], uint64(decPart.IntPart()))

	return b
}
