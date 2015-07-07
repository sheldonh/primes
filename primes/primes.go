package primes

import "math"
import "math/big"

func IsPrime(n int64) bool {
	switch {
	case n < 2:
		return false
	case n == 2:
		return true
	case n&1 == 0:
		return false
	default:
		switch {
		case n < 5963:
			return big.NewInt(int64(n)).ProbablyPrime(1)
		case n < 206981:
			return big.NewInt(int64(n)).ProbablyPrime(2)
		case n < 587861:
			return big.NewInt(int64(n)).ProbablyPrime(3)
		default:
			if big.NewInt(int64(n)).ProbablyPrime(4) {
				s := int64(math.Sqrt(float64(n)))
				for i := int64(3); i <= s; i += 2 {
					if n%i == 0 {
						return false
					}
				}
				return true
			} else {
				return false
			}
		}
	}
}
