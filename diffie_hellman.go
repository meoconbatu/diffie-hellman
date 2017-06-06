package diffiehellman

import (
	"math"
	"math/big"
	"math/rand"
	"time"
)

const testVersion = 1

func PrivateKey(p *big.Int) *big.Int {
	rand.Seed(time.Now().UTC().UnixNano())
	o := big.NewInt(2 + rand.Int63n(int64(math.Abs(float64(p.Int64()-2)))))
	return o
}
func PublicKey(private, p *big.Int, g int64) *big.Int {
	gBigInt := big.NewInt(g)
	return modulo(gBigInt, private, p)
}
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = PublicKey(private, p, g)
	return
}
func SecretKey(private1, public2, p *big.Int) *big.Int {
	return modulo(public2, private1, p)
}
func modulo(a, b, n *big.Int) *big.Int {
	x := big.NewInt(1)
	y := a
	for b.Cmp(big.NewInt(0)) >= 1 {
		if (new(big.Int).Mod(b, big.NewInt(2))).Cmp(big.NewInt(1)) == 0 {
			x = new(big.Int).Mod(new(big.Int).Mul(x, y), n)
		}
		y = new(big.Int).Mod(new(big.Int).Mul(y, y), n)
		b = new(big.Int).Div(b, big.NewInt(2))
	}
	return new(big.Int).Mod(x, n)
}
