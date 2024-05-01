package grs

type GitCommitResolver interface {
	GetCommitHash(url, ref string) (string, error)
}
