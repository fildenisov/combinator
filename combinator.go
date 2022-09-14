package combinator

import "context"

func Recombo[T any](ctx context.Context, a []T, results chan<- []T) {
	rc(ctx, []T{}, a, results)
	close(results)
}

func rc[T any](ctx context.Context, res []T, a []T, results chan<- []T) bool {
	select {
	case <-ctx.Done():
		return false
	default:
		break
	}

	if len(a) == 0 {
		results <- res
		return true
	}

	excluded := make([]T, len(a)-1)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if j < i {
				excluded[j] = a[j]
			} else if j > i {
				excluded[j-1] = a[j]
			}
		}

		if !rc(ctx, append(res, a[i]), excluded, results) {
			break
		}
	}

	return true
}
