package ratelimit

type Rater interface {
	Allow() bool
}