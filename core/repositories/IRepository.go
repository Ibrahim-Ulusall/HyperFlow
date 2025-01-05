package repositories

type IRepository[TEntity any] interface {
	Any() bool
	Add(enttiy TEntity)
	Update(entitiy TEntity)
	Delete(entity TEntity)
	GetById(id int) TEntity
	GetList() []TEntity
}
