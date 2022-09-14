package combinator

import "context"

func Recombo(ctx context.Context, a []rune, results chan<- []rune) {
	rc(ctx, []rune{}, a, results)
	close(results)
}

func rc(ctx context.Context, res []rune, a []rune, results chan<- []rune) bool {
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

	excluded := make([]rune, len(a)-1)
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
