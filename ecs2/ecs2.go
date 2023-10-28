package ecs2

import (
	"sync"

	"github.com/gopxl/pixel"
)

type Entity uint32

const (
	maxComponents = 32
	maxEntities   = 1 << 20
)

var nextEntity Entity
var entityMutex sync.Mutex

func NewEntity() Entity {
	entityMutex.Lock()
	defer entityMutex.Unlock()

	nextEntity++
	return nextEntity
}

type ComponentType uint8
type ComponentMask uint32

var nextComponentType ComponentType

func NewComponentType() ComponentType {
	nextComponentType++
	return nextComponentType
}

type Component interface {
	Type() ComponentType
}

type EntityManager struct {
	entities       map[Entity]ComponentMask
	components     [maxEntities][maxComponents]Component
	componentMutex sync.RWMutex
}

func NewEntityManager() *EntityManager {
	return &EntityManager{
		entities: make(map[Entity]ComponentMask),
	}
}

func (em *EntityManager) AddComponent(e Entity, component Component) {
	em.componentMutex.Lock()
	defer em.componentMutex.Unlock()

	mask := uint32(1) << component.Type()
	em.entities[e] |= ComponentMask(mask)
	em.components[e][component.Type()] = component
}

func (em *EntityManager) GetComponent(e Entity, componentType ComponentType) Component {
	em.componentMutex.RLock()
	defer em.componentMutex.RUnlock()

	return em.components[e][componentType]
}

func (em *EntityManager) RemoveComponent(e Entity, componentType ComponentType) {
	em.componentMutex.Lock()
	defer em.componentMutex.Unlock()

	mask := ^(uint32(1) << componentType)
	em.entities[e] &= ComponentMask(mask)
	em.components[e][componentType] = nil
}

type System interface {
	Update(dt float64, em *EntityManager)
}

type PositionComponent struct {
	position pixel.Vec
}

func (p *PositionComponent) Type() ComponentType {
	return NewComponentType()
}

type VelocityComponent struct {
	velocity pixel.Vec
}

func (v *VelocityComponent) Type() ComponentType {
	return NewComponentType()
}

type MovementSystem struct{}

func (ms *MovementSystem) Update(dt float64, em *EntityManager) {
	for e, mask := range em.entities {
		if mask&((1<<(&PositionComponent{}).Type())|(1<<(&VelocityComponent{}).Type())) != 0 {
			pos := em.GetComponent(e, (&PositionComponent{}).Type()).(*PositionComponent)
			vel := em.GetComponent(e, (&VelocityComponent{}).Type()).(*VelocityComponent)
			pos.position.X += vel.velocity.X * dt
			pos.position.Y += vel.velocity.Y * dt
		}
	}
}
