// Package incr implements a global random incremental map for characters.
package incr

import (
	"hash/maphash"
	"math/rand"
	"sync"

	"github.com/diamondburned/aqs"
)

// State provides an incremental permutation state. It is thread-safe.
type State struct {
	permus []int
	ptr    int
}

func New(n int) *State {
	return &State{permus: rand.Perm(n), ptr: 0}
}

// Next returns the next random index.
func (s *State) Next() int {
	i := s.ptr
	s.ptr = (i + 1) % len(s.permus)

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
	if len(char.Quotes) == 0 {
		return ""
	}

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

	return char.Quotes[st.Next()]
}
