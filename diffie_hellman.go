package diffiehellman

import (
	"math/big"
	"math/rand"
	"time"
)

const testVersion = 1

func PrivateKey(p *big.Int) *big.Int {
	source := rand.New(rand.NewSource(time.Now().UnixNano()))
	return new(big.Int).Add(big.NewInt(2), new(big.Int).Rand(source, new(big.Int).Sub(p, big.NewInt(2))))
}
func PublicKey(private, p *big.Int, g int64) *big.Int {
	gBigInt := big.NewInt(g)
	return new(big.Int).Exp(gBigInt, private, p)
}
func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	private := PrivateKey(p)
	return private, PublicKey(private, p, g)
}
func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}
