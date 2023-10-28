package scenes

type EntityManager *[][]Component

type Scene interface {
	GetEntityManager() EntityManager
	Update()
	DoAction(action *Action)
	Render()
}