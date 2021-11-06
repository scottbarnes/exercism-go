package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// Key to the keys.
//  g    int64    // prime, generator
//  p    *big.Int // prime, modulus
//  a, b *big.Int // private keys
//  A, B *big.Int // public keys
//  s    *big.Int // secret key

// PrivateKey randomly generates private keys, from 2 to the limit p.
func PrivateKey(p *big.Int) *big.Int {
	bigTwo := big.NewInt(2)            // rand.Int returns a random number from 0 to limit. Add two to sort this.
	max := new(big.Int).Sub(p, bigTwo) // subtract two from the limit to account for the added two.
	privKey, err := rand.Int(rand.Reader, max)

	if err != nil {
		panic(err)
	}

	return privKey.Add(privKey, bigTwo)
}

// PublicKey randomly generates public keys
func PublicKey(private, p *big.Int, g int64) *big.Int {
	var pubKey big.Int
	// Exp sets z = x**y mod |m| (i.e. the sign of m is ignored), and returns z.
	return pubKey.Exp(big.NewInt(g), private, p)
}

// NewPair generates a pair of private and public keys for use with Diffie-Hellman key exchange.
func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	a := PrivateKey(p)
	A := PublicKey(a, p, g)
	return a, A
}

// SecretKey generates a shared secret key.
func SecretKey(private1, public2, p *big.Int) *big.Int {
	var secKey big.Int
	return secKey.Exp(public2, private1, p)
}
