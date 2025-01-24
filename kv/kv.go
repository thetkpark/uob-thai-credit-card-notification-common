package kv

type KV interface {
	Add(k string) error
	Exist(k string) (bool, error)
}
