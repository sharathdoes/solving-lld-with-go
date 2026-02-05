package cashdispenser


import "errors"

type Denomination struct {
	Value int
	Count int
}

type CoinChangeDispenser struct {
	denoms []Denomination
}

func NewCoinChangeDispenser(denoms []Denomination) *CoinChangeDispenser {
	return &CoinChangeDispenser{denoms: denoms}
}

func (c *CoinChangeDispenser) Dispense(amount int) (map[int]int, error) {
	result := make(map[int]int)

	for i := 0; i < len(c.denoms); i++ {
		d := &c.denoms[i]

		if amount <= 0 {
			break
		}

		maxNotes := amount / d.Value
		if maxNotes > d.Count {
			maxNotes = d.Count
		}

		if maxNotes > 0 {
			result[d.Value] = maxNotes
			amount -= maxNotes * d.Value
			d.Count -= maxNotes
		}
	}

	if amount != 0 {
		return nil, errors.New("cannot dispense exact amount")
	}

	return result, nil
}
