package pacscript

import "pacstall.dev/webserver/types/pac"

type PacscriptListWrapper []*pac.Script

func (w PacscriptListWrapper) Len() int {
	return len(w)
}

func (w PacscriptListWrapper) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}
