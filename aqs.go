package aqs

import (
	"hash/maphash"
	"math"
	"math/rand"
	"time"

	"github.com/lucasb-eyer/go-colorful"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// These variables contain the saturation and luminance used for name color
// generation. By default, they use pastel-esque numbers.
var (
	// Saturation is the default saturation used for name color generation.
	Saturation = 0.70
	// Luminance is the default luminance used for name color generation.
	Luminance = 0.80
)

// Characters stores multiple anime characters. This slice is empty by default.
// To fill them up with default data, add this to the imports:
//
//     _ "github.com/diamondburned/aqs/data"
var Characters []Character

// Character represents a single anime character.
type Character struct {
	Name     string   `json:"name"`
	Anime    string   `json:"anime"`
	ImageURL string   `json:"image"`
	Quotes   []string `json:"quotes"`
}

// RandomCharacter returns a random character, or a zero-value if there's none.
func RandomCharacter() Character {
	if len(Characters) == 0 {
		return Character{}
	}

	return Characters[rand.Intn(len(Characters))]
}

// SearchCharacter searches for a character using exact match. A zero-value is
// returned if none is found.
func SearchCharacter(name string) Character {
	for _, ch := range Characters {
		if ch.Name == name {
			return ch
		}
	}

	return Character{}
}

// RandomQuote returns a random quote from the character.
func (c Character) RandomQuote() string {
	if len(c.Quotes) == 0 {
		return ""
	}

	return c.Quotes[rand.Intn(len(c.Quotes))]
}

const maxu64 = float64(math.MaxUint64)

// NameColor returns a consistent name color for the character.
func (c Character) NameColor() colorful.Color {
	hash := maphash.Hash{}
	hash.WriteString(c.Name)
	hash.Sum64()

	hue := float64(hash.Sum64()) / maxu64 * 360
	return colorful.Hsl(hue, Saturation, Luminance)
}
