package boundary

import (
	"fmt"
	"sort"
)

func keys(m map[string]int) []string {
	keys := make([]string, len(m))
	var i int
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func count(a []string) map[string]int {
	count := make(map[string]int)
	for _, s := range a {
		if _, ok := count[s]; ok {
			count[s]++
		} else {
			count[s] = 1
		}
	}
	return count
}

func mostCommonN(a []string, n int) []string {
	count := count(a)
	uniques := keys(count)
	byCountLargerFirst := func(i, j int) bool {
		return count[uniques[i]] > count[uniques[j]]
	}
	sort.Slice(uniques, byCountLargerFirst)
	return uniques[:n]
	// will cause runtime error if n > len(a) or n < 0; n==0 will give you an empty slice, which may be correct but is nonsensical.
	//further, this output is nondeterminstic, because map iteration order in golang is arbitrary.
	// i.e, mostCommonN([]string{"a", "b", "c", "d"}, 2) could give you any possible pair.
	// this is bad!
}

func mostCommonNFixed(a []string, n int) ([]string, error) {
	if n <= 0 {
		return nil, fmt.Errorf("mostCommonN takes a positive integer argument")
	} else if len(a) < n {
		return nil, fmt.Errorf("cannot take %d most common elements from a slice of len %d", n, len(a))
	}
	count := count(a)
	uniques := keys(count)
	byLargerCountThenLexigraphic := func(i, j int) bool {
		a, b := uniques[i], uniques[j]
		return count[a] > count[b] || (count[a] == count[b] && a < b)
	}
	sort.Slice(uniques, byLargerCountThenLexigraphic)
	return uniques[:n], nil
}
