// Package incr implements a global random incremental map for characters.
package incr

import (
	"hash/maphash"
	"math/rand"
	"sync"

	"github.com/diamondburned/aqs"
)

type State struct {
	permus []int
	ptr    int
}

func New(n int) *State {
	// Numerical Recipes' LCG.
	return &State{rand.Perm(n), 0}
}

// Next returns the next random index.
func (s *State) Next() int {
	i := s.ptr

	s.ptr++
	if s.ptr >= len(s.permus) {
		s.ptr = 0
	}

	return s.permus[i]
}

var randmap = map[uint64]*State{}
var randmut = sync.Mutex{}

// RandomQuote returns a random quote using a global randomly-permutated
// increment map. This function is safe for concurrent use. The unique
// constraint for this function is the character's name and anime.
func RandomQuote(char aqs.Character) string {
	hash := maphash.Hash{}
	hash.WriteString(char.Name)
	hash.WriteString(char.Anime)
	u := hash.Sum64()

	randmut.Lock()
	defer randmut.Unlock()

	st, ok := randmap[u]
	if !ok {
		st = New(len(char.Quotes))
		randmap[u] = st
	}

	return char.Quotes[st.Next()]
}
