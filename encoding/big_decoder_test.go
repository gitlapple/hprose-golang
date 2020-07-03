/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/big_decoder_test.go                             |
|                                                          |
| LastModified: Jul 3, 2020                                |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"math"
	"math/big"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeBigInt(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode(-1)
	enc.Encode(0)
	enc.Encode(1)
	enc.Encode(123)
	enc.Encode(math.MinInt64)
	enc.Encode(-math.MaxInt64)
	enc.Encode(uint64(math.MaxUint64))
	enc.Encode(true)
	enc.Encode(false)
	enc.Encode(nil)
	enc.Encode(math.NaN())
	enc.Encode(math.Inf(1))
	enc.Encode(math.Inf(-1))
	enc.Encode(3.14)
	enc.Encode("")
	enc.Encode("1")
	enc.Encode("123")
	enc.Encode("1234567890123456789012345678901234567890")
	enc.Encode("NaN")
	bi, _ := new(big.Int).SetString("1234567890123456789012345678901234567890", 10)
	enc.Encode(bi)
	dec := NewDecoder(([]byte)(sb.String()))
	var i big.Int
	dec.Decode(&i)
	assert.Equal(t, big.NewInt(-1), &i)
	dec.Decode(&i)
	assert.Equal(t, big.NewInt(0), &i)
	dec.Decode(&i)
	assert.Equal(t, big.NewInt(1), &i)
	dec.Decode(&i)
	assert.Equal(t, big.NewInt(123), &i)
	dec.Decode(&i)
	assert.Equal(t, big.NewInt(math.MinInt64), &i)
	dec.Decode(&i)
	assert.Equal(t, big.NewInt(-math.MaxInt64), &i)
	dec.Decode(&i)
	assert.Equal(t, new(big.Int).SetUint64(math.MaxUint64), &i)
	dec.Decode(&i)
	assert.Equal(t, big.NewInt(1), &i)
	dec.Decode(&i)
	assert.Equal(t, big.NewInt(0), &i)
	dec.Decode(&i)
	assert.Equal(t, big.NewInt(0), &i)
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, "hprose/encoding: can not parse NaN to big.Int")
	dec.Error = nil
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, "hprose/encoding: can not parse +Inf to big.Int")
	dec.Error = nil
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, "hprose/encoding: can not parse -Inf to big.Int")
	dec.Error = nil
	dec.Decode(&i)
	assert.Equal(t, big.NewInt(3), &i)
	dec.Decode(&i)
	assert.Equal(t, big.NewInt(0), &i)
	dec.Decode(&i)
	assert.Equal(t, big.NewInt(1), &i)
	dec.Decode(&i)
	assert.Equal(t, big.NewInt(123), &i)
	assert.NoError(t, dec.Error)
	dec.Decode(&i)
	assert.Equal(t, bi, &i)
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not parse "NaN" to big.Int`)
	dec.Error = nil
	dec.Decode(&i)
	assert.Equal(t, bi, &i)
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, "EOF")
}

func TestDecodeBigIntPtr(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode(-1)
	enc.Encode(0)
	enc.Encode(1)
	enc.Encode(123)
	enc.Encode(math.MinInt64)
	enc.Encode(-math.MaxInt64)
	enc.Encode(uint64(math.MaxUint64))
	enc.Encode(true)
	enc.Encode(false)
	enc.Encode(nil)
	enc.Encode(math.NaN())
	enc.Encode(math.Inf(1))
	enc.Encode(math.Inf(-1))
	enc.Encode(3.14)
	enc.Encode("")
	enc.Encode("1")
	enc.Encode("123")
	enc.Encode("1234567890123456789012345678901234567890")
	enc.Encode("NaN")
	bi, _ := new(big.Int).SetString("1234567890123456789012345678901234567890", 10)
	enc.Encode(bi)
	dec := NewDecoder(([]byte)(sb.String()))
	var i *big.Int
	dec.Decode(&i)
	assert.Equal(t, big.NewInt(-1), i)
	dec.Decode(&i)
	assert.Equal(t, big.NewInt(0), i)
	dec.Decode(&i)
	assert.Equal(t, big.NewInt(1), i)
	dec.Decode(&i)
	assert.Equal(t, big.NewInt(123), i)
	dec.Decode(&i)
	assert.Equal(t, big.NewInt(math.MinInt64), i)
	dec.Decode(&i)
	assert.Equal(t, big.NewInt(-math.MaxInt64), i)
	dec.Decode(&i)
	assert.Equal(t, new(big.Int).SetUint64(math.MaxUint64), i)
	dec.Decode(&i)
	assert.Equal(t, big.NewInt(1), i)
	dec.Decode(&i)
	assert.Equal(t, big.NewInt(0), i)
	dec.Decode(&i)
	assert.Nil(t, i)
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, "hprose/encoding: can not parse NaN to *big.Int")
	dec.Error = nil
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, "hprose/encoding: can not parse +Inf to *big.Int")
	dec.Error = nil
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, "hprose/encoding: can not parse -Inf to *big.Int")
	dec.Error = nil
	dec.Decode(&i)
	assert.Equal(t, big.NewInt(3), i)
	dec.Decode(&i)
	assert.Equal(t, big.NewInt(0), i)
	dec.Decode(&i)
	assert.Equal(t, big.NewInt(1), i)
	dec.Decode(&i)
	assert.Equal(t, big.NewInt(123), i)
	assert.NoError(t, dec.Error)
	dec.Decode(&i)
	assert.Equal(t, bi, i)
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not parse "NaN" to *big.Int`)
	dec.Error = nil
	dec.Decode(&i)
	assert.Equal(t, bi, i)
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, "EOF")
}

func TestDecodeBigFloat(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode(-1)
	enc.Encode(0)
	enc.Encode(1)
	enc.Encode(123)
	enc.Encode(math.MinInt64)
	enc.Encode(-math.MaxInt64)
	enc.Encode(uint64(math.MaxUint64))
	enc.Encode(true)
	enc.Encode(false)
	enc.Encode(nil)
	enc.Encode(math.NaN())
	enc.Encode(math.Inf(1))
	enc.Encode(math.Inf(-1))
	enc.Encode(3.14)
	enc.Encode("")
	enc.Encode("1")
	enc.Encode("123")
	enc.Encode("12345678901234567890.12345678901234567890")
	enc.Encode("NaN")
	bf, _ := new(big.Float).SetString("12345678901234567890.12345678901234567890")
	enc.Encode(bf)
	println(sb.String())
	dec := NewDecoder(([]byte)(sb.String()))
	var i big.Float
	dec.Decode(&i)
	assert.Equal(t, big.NewFloat(-1), &i)
	dec.Decode(&i)
	assert.Equal(t, big.NewFloat(0), &i)
	dec.Decode(&i)
	assert.Equal(t, big.NewFloat(1), &i)
	dec.Decode(&i)
	assert.Equal(t, big.NewFloat(123), &i)
	dec.Decode(&i)
	assert.Equal(t, new(big.Float).SetInt64(math.MinInt64), &i)
	dec.Decode(&i)
	assert.Equal(t, new(big.Float).SetInt64(-math.MaxInt64), &i)
	dec.Decode(&i)
	assert.Equal(t, new(big.Float).SetUint64(math.MaxUint64), &i)
	dec.Decode(&i)
	assert.Equal(t, big.NewFloat(1), &i)
	dec.Decode(&i)
	assert.Equal(t, big.NewFloat(0), &i)
	dec.Decode(&i)
	assert.Equal(t, big.NewFloat(0), &i)
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, "hprose/encoding: can not parse NaN to big.Float")
	dec.Error = nil
	dec.Decode(&i)
	assert.Equal(t, big.NewFloat(math.Inf(1)), &i)
	dec.Decode(&i)
	assert.Equal(t, big.NewFloat(math.Inf(-1)), &i)
	dec.Decode(&i)
	assert.Equal(t, "3.14", i.String())
	dec.Decode(&i)
	assert.Equal(t, big.NewFloat(0), &i)
	dec.Decode(&i)
	assert.Equal(t, "1", i.String())
	dec.Decode(&i)
	assert.Equal(t, "123", i.String())
	assert.NoError(t, dec.Error)
	dec.Decode(&i)
	assert.Equal(t, bf.String(), i.String())
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not parse "NaN" to big.Float`)
	dec.Error = nil
	dec.Decode(&i)
	assert.Equal(t, bf.String(), i.String())
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, "EOF")
}

func TestDecodeBigFloatPtr(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode(-1)
	enc.Encode(0)
	enc.Encode(1)
	enc.Encode(123)
	enc.Encode(math.MinInt64)
	enc.Encode(-math.MaxInt64)
	enc.Encode(uint64(math.MaxUint64))
	enc.Encode(true)
	enc.Encode(false)
	enc.Encode(nil)
	enc.Encode(math.NaN())
	enc.Encode(math.Inf(1))
	enc.Encode(math.Inf(-1))
	enc.Encode(3.14)
	enc.Encode("")
	enc.Encode("1")
	enc.Encode("123")
	enc.Encode("12345678901234567890.12345678901234567890")
	enc.Encode("NaN")
	bf, _ := new(big.Float).SetString("12345678901234567890.12345678901234567890")
	enc.Encode(bf)
	println(sb.String())
	dec := NewDecoder(([]byte)(sb.String()))
	var i *big.Float
	dec.Decode(&i)
	assert.Equal(t, big.NewFloat(-1), i)
	dec.Decode(&i)
	assert.Equal(t, big.NewFloat(0), i)
	dec.Decode(&i)
	assert.Equal(t, big.NewFloat(1), i)
	dec.Decode(&i)
	assert.Equal(t, big.NewFloat(123), i)
	dec.Decode(&i)
	assert.Equal(t, new(big.Float).SetInt64(math.MinInt64), i)
	dec.Decode(&i)
	assert.Equal(t, new(big.Float).SetInt64(-math.MaxInt64), i)
	dec.Decode(&i)
	assert.Equal(t, new(big.Float).SetUint64(math.MaxUint64), i)
	dec.Decode(&i)
	assert.Equal(t, big.NewFloat(1), i)
	dec.Decode(&i)
	assert.Equal(t, big.NewFloat(0), i)
	dec.Decode(&i)
	assert.Nil(t, i)
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, "hprose/encoding: can not parse NaN to *big.Float")
	dec.Error = nil
	dec.Decode(&i)
	assert.Equal(t, big.NewFloat(math.Inf(1)), i)
	dec.Decode(&i)
	assert.Equal(t, big.NewFloat(math.Inf(-1)), i)
	dec.Decode(&i)
	assert.Equal(t, "3.14", i.String())
	dec.Decode(&i)
	assert.Equal(t, big.NewFloat(0), i)
	dec.Decode(&i)
	assert.Equal(t, "1", i.String())
	dec.Decode(&i)
	assert.Equal(t, "123", i.String())
	assert.NoError(t, dec.Error)
	dec.Decode(&i)
	assert.Equal(t, bf.String(), i.String())
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not parse "NaN" to *big.Float`)
	dec.Error = nil
	dec.Decode(&i)
	assert.Equal(t, bf.String(), i.String())
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, "EOF")
}

func TestDecodeBigRat(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode(-1)
	enc.Encode(0)
	enc.Encode(1)
	enc.Encode(123)
	enc.Encode(math.MinInt64)
	enc.Encode(-math.MaxInt64)
	enc.Encode(uint64(math.MaxUint64))
	enc.Encode(true)
	enc.Encode(false)
	enc.Encode(nil)
	enc.Encode(math.NaN())
	enc.Encode(math.Inf(1))
	enc.Encode(math.Inf(-1))
	enc.Encode(3.14)
	enc.Encode("")
	enc.Encode("1")
	enc.Encode("123")
	enc.Encode("1234567890123456789012345678901234567890")
	enc.Encode("NaN")
	bi, _ := new(big.Int).SetString("1234567890123456789012345678901234567890", 10)
	enc.Encode(bi)
	br, _ := new(big.Rat).SetString("12345678901234567890/12345678901234567890")
	enc.Encode(br)
	dec := NewDecoder(([]byte)(sb.String()))
	var i big.Rat
	dec.Decode(&i)
	assert.Equal(t, big.NewRat(-1, 1), &i)
	dec.Decode(&i)
	assert.Equal(t, big.NewRat(0, 1), &i)
	dec.Decode(&i)
	assert.Equal(t, big.NewRat(1, 1), &i)
	dec.Decode(&i)
	assert.Equal(t, big.NewRat(123, 1), &i)
	dec.Decode(&i)
	assert.Equal(t, new(big.Rat).SetInt64(math.MinInt64), &i)
	dec.Decode(&i)
	assert.Equal(t, new(big.Rat).SetInt64(-math.MaxInt64), &i)
	dec.Decode(&i)
	assert.Equal(t, new(big.Rat).SetUint64(math.MaxUint64), &i)
	dec.Decode(&i)
	assert.Equal(t, big.NewRat(1, 1), &i)
	dec.Decode(&i)
	assert.Equal(t, big.NewRat(0, 1), &i)
	dec.Decode(&i)
	assert.Equal(t, big.NewRat(0, 1), &i)
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, "hprose/encoding: can not parse NaN to big.Rat")
	dec.Error = nil
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, "hprose/encoding: can not parse +Inf to big.Rat")
	dec.Error = nil
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, "hprose/encoding: can not parse -Inf to big.Rat")
	dec.Error = nil
	dec.Decode(&i)
	assert.Equal(t, new(big.Rat).SetFloat64(3.14), &i)
	dec.Decode(&i)
	assert.Equal(t, big.NewRat(0, 1), &i)
	dec.Decode(&i)
	assert.Equal(t, big.NewRat(1, 1), &i)
	dec.Decode(&i)
	assert.Equal(t, big.NewRat(123, 1), &i)
	assert.NoError(t, dec.Error)
	dec.Decode(&i)
	assert.Equal(t, new(big.Rat).SetInt(bi).RatString(), (&i).RatString())
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not parse "NaN" to big.Rat`)
	dec.Error = nil
	dec.Decode(&i)
	assert.Equal(t, new(big.Rat).SetInt(bi), &i)
	dec.Decode(&i)
	assert.Equal(t, br.RatString(), (&i).RatString())
	dec.Error = nil
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, "EOF")
}

func TestDecodeBigRatPtr(t *testing.T) {
	sb := new(strings.Builder)
	enc := NewEncoder(sb)
	enc.Encode(-1)
	enc.Encode(0)
	enc.Encode(1)
	enc.Encode(123)
	enc.Encode(math.MinInt64)
	enc.Encode(-math.MaxInt64)
	enc.Encode(uint64(math.MaxUint64))
	enc.Encode(true)
	enc.Encode(false)
	enc.Encode(nil)
	enc.Encode(math.NaN())
	enc.Encode(math.Inf(1))
	enc.Encode(math.Inf(-1))
	enc.Encode(3.14)
	enc.Encode("")
	enc.Encode("1")
	enc.Encode("123")
	enc.Encode("1234567890123456789012345678901234567890")
	enc.Encode("NaN")
	bi, _ := new(big.Int).SetString("1234567890123456789012345678901234567890", 10)
	enc.Encode(bi)
	br, _ := new(big.Rat).SetString("12345678901234567890/12345678901234567890")
	enc.Encode(br)
	dec := NewDecoder(([]byte)(sb.String()))
	var i *big.Rat
	dec.Decode(&i)
	assert.Equal(t, big.NewRat(-1, 1), i)
	dec.Decode(&i)
	assert.Equal(t, big.NewRat(0, 1), i)
	dec.Decode(&i)
	assert.Equal(t, big.NewRat(1, 1), i)
	dec.Decode(&i)
	assert.Equal(t, big.NewRat(123, 1), i)
	dec.Decode(&i)
	assert.Equal(t, new(big.Rat).SetInt64(math.MinInt64), i)
	dec.Decode(&i)
	assert.Equal(t, new(big.Rat).SetInt64(-math.MaxInt64), i)
	dec.Decode(&i)
	assert.Equal(t, new(big.Rat).SetUint64(math.MaxUint64), i)
	dec.Decode(&i)
	assert.Equal(t, big.NewRat(1, 1), i)
	dec.Decode(&i)
	assert.Equal(t, big.NewRat(0, 1), i)
	dec.Decode(&i)
	assert.Nil(t, i)
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, "hprose/encoding: can not parse NaN to *big.Rat")
	dec.Error = nil
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, "hprose/encoding: can not parse +Inf to *big.Rat")
	dec.Error = nil
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, "hprose/encoding: can not parse -Inf to *big.Rat")
	dec.Error = nil
	dec.Decode(&i)
	assert.Equal(t, new(big.Rat).SetFloat64(3.14), i)
	dec.Decode(&i)
	assert.Equal(t, big.NewRat(0, 1), i)
	dec.Decode(&i)
	assert.Equal(t, big.NewRat(1, 1), i)
	dec.Decode(&i)
	assert.Equal(t, big.NewRat(123, 1), i)
	assert.NoError(t, dec.Error)
	dec.Decode(&i)
	assert.Equal(t, new(big.Rat).SetInt(bi).RatString(), i.RatString())
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, `hprose/encoding: can not parse "NaN" to *big.Rat`)
	dec.Error = nil
	dec.Decode(&i)
	assert.Equal(t, new(big.Rat).SetInt(bi), i)
	dec.Decode(&i)
	assert.Equal(t, br.RatString(), i.RatString())
	dec.Error = nil
	dec.Decode(&i)
	assert.EqualError(t, dec.Error, "EOF")
}
