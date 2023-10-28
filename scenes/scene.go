package scenes

type EntityManager []ComponentVector

type Scene interface {
	GetEntityManager() EntityManager
	AddEntity() *ComponentVector
	Update()
	DoAction(action *string)
	Render()
}
