// Package incr implements a global random incremental map for characters.
package incr

import (
	"hash/maphash"
	"log"
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

var hasher maphash.Hash

func init() {
	hasher = maphash.Hash{}
	hasher.SetSeed(maphash.MakeSeed())
}

// RandomQuote returns a random quote using a global randomly-permutated
// increment map. This function is safe for concurrent use. The unique
// constraint for this function is the character's name and anime.
func RandomQuote(char aqs.Character) string {
	randmut.Lock()
	defer randmut.Unlock()

	hasher.WriteString(char.Name)
	hasher.WriteString(char.Anime)
	u := hasher.Sum64()
	hasher.Reset()

	st, ok := randmap[u]
	if !ok {
		st = New(len(char.Quotes))
		randmap[u] = st
	}

	log.Println("State:", u, char.Name, st)

	return char.Quotes[st.Next()]
}
