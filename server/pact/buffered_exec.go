package pact

import "sync"

// @generate pkginfo
// @import pacstall.dev/website/types
// @replace TResult >> *types.PackageInfo
// @replace TInput >> string
// @replace AsyncBuffered >> AsyncBufferedPackageInfo

type TResult = interface{}
type TInput = interface{}

func AsyncBuffered(bufSize int, values []TInput, handle func(TInput) TResult) <-chan TResult {
	guard := make(chan int, bufSize)
	output := make(chan TResult)

	go func() {
		mutex := sync.Mutex{}
		amountLeft := len(values)

		for i, value := range values {
			guard <- i

			go func(value TInput, i int) {
				output <- handle(value)
				<-guard

				mutex.Lock()
				defer mutex.Unlock()

				amountLeft -= 1
				if amountLeft == 0 {
					close(guard)
					close(output)
				}

			}(value, i)
		}
	}()

	return output
}
