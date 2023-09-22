package graphql

type SubscriptionResult[T any] struct {
	data  T
	error error
}

func SubscriptionResultData[T any](data T) SubscriptionResult[T] {
	return SubscriptionResult[T]{
		data: data,
	}
}

func SubscriptionResultError[T any](err error) SubscriptionResult[T] {
	return SubscriptionResult[T]{
		error: err,
	}
}

func (sr *SubscriptionResult[T]) Result() (T, error) {
	return sr.data, sr.error
}
