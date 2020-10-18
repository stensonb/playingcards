package card

const (
	Clubs Suit = iota
	Diamonds
	Hearts
	Spades
)

// note the distict absence of numerical value to these card "values"
// this is intentionally detached so that any given game can assign "rank"
// to a card's value.
const (
	Ace Value = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Suit int

func (s *Suit) String() string {
	switch *s {
	case Spades:
		return "♠"
	case Diamonds:
		return "♦"
	case Hearts:
		return "♥"
	default: // Clubs
		return "♣"
	}
}

type Value int

func (v *Value) String() string {
	valueStrings := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	return valueStrings[*v]
}

var AllSuits = []Suit{Hearts, Spades, Diamonds, Clubs}
var AllValues = []Value{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}

type Card struct {
	Suit  Suit
	Value Value
}

func (c *Card) String() string {
	return c.Value.String() + c.Suit.String()
}

// returns true only when both cards' suits and values match
// ie, two "Ten of Hearts" would be true
// ie, "Ten of Hearts" and "Ten of Spades" would be false
func (c *Card) Equals(o *Card) bool {
	if c.Suit == o.Suit && c.Value == o.Value {
		return true
	}
	return false
}
