package Union_Find

type UF interface {
	Union(p, q int) error
	IsConnected(p, q int) (bool, error)
	GetID(p int) (int, error)
}
