package repositories

type DataRepository[K comparable, V any] interface {
	Get(key K) (V, error)
}
