package container

import (
	"github.com/wishperera/GVAT/gen/internal/domain"
	"log"
)

const (
	moduleUnregistered = "module: [%s] is not registered, please use container.Bind() to register first\n"
)

// Container : Stores the dependencies for the application
type Container interface {
	// Bind : binds a dependency with the provided name
	Bind(name string, value interface{})
	// Resolve : resolves a dependency by name
	Resolve(name string) interface{}
	// Init : initialize the dependencies in the provided order
	Init(modules ...string)
	// GetModuleConfig : get the module config
	GetModuleConfig(name string) domain.Config
}

// AppContainer : container for runnable dependencies
type AppContainer interface {
	Container
	Start(modules ...string)
	ShutDown(modules ...string)
	SetModuleConfig(name string, value domain.Config)
}

// Runnable : Defines a runnable module ex: http router
type Runnable interface {
	// Run : run the module
	Run() error
	// Ready : return a channel that signals when the module is ready. useful for ensuring sequential execution
	// of runnable module initiations.
	Ready() chan struct{}
}

// Stoppable : Defines a stoppable module ex: http router
type Stoppable interface {
	// Stop : stop the module
	Stop() error
}

func NewContainment() AppContainer {
	c := new(Containment)
	c.bindings = make(map[string]interface{})
	c.configs = make(map[string]domain.Config)

	return c
}

// Containment : Implementation of AppContainer interface
type Containment struct {
	bindings map[string]interface{}
	configs  map[string]domain.Config
}

func (c *Containment) Bind(name string, value interface{}) {
	_, ok := c.bindings[name]
	if !ok { // check if the module is already inited
		c.bindings[name] = value
	} else {
		log.Fatalf("module: [%s] already registered\n", name)
	}
}

func (c *Containment) Resolve(name string) interface{} {
	binding, ok := c.bindings[name]
	if !ok {
		log.Fatalf("resolve failed,module: [%s] is not registered", binding)
	}

	return binding
}

func (c *Containment) Init(modules ...string) {
	for index := range c.configs {
		err := c.configs[index].Init()
		if err != nil {
			log.Fatalf("failed to initialize module config due: %s\n", err)
		}

		err = c.configs[index].Validate()
		if err != nil {
			log.Fatalf("failed to validate module config due: %s\n", err)
		}

		log.Printf("initialized config for module: %s config: %s\n", index, c.configs[index].Print())
	}

	for index := range modules {
		// check if the provided binding is registered
		binding, ok := c.bindings[modules[index]]
		if !ok {
			log.Fatalf(moduleUnregistered, binding)
		}

		// check if the provided binding is a module
		module, ok := binding.(Module)
		if !ok {
			log.Fatalf("provided binding: [%s] is not a module\n", modules[index])
		}

		// initialize the module
		err := module.Init(c)
		if err != nil {
			log.Fatalf("failed to initialize module: [%s] due: [%s]\n", modules[index], err)
		}

		log.Printf("initialized module: %s\n", modules[index])
	}
}

// Start : start the runnable modules in the provided order
func (c *Containment) Start(modules ...string) {
	for index := range modules {
		// check if the provided binding is registered
		binding, ok := c.bindings[modules[index]]
		if !ok {
			log.Fatalf(moduleUnregistered, binding)
		}

		// check if the provided binding is a runnable module
		module, ok := binding.(Runnable)
		if !ok {
			log.Fatalf("provided module: [%s] is not runnable\n", modules[index])
		}

		// run the module and wait for the ready state
		err := module.Run()
		if err != nil {
			log.Fatalf("failed to run module: [%s] due: [%s]\n", modules[index], err)
		}

		<-module.Ready()
		log.Printf("module: [%s] started\n", modules[index])
	}
}

// ShutDown : stop the runnable modules in the provided order
func (c *Containment) ShutDown(modules ...string) {
	for index := range modules {
		// check if the provided binding is registered
		binding, ok := c.bindings[modules[index]]
		if !ok {
			log.Fatalf(moduleUnregistered, binding)
		}

		// check if the provided binding is a runnable module
		module, ok := binding.(Stoppable)
		if !ok {
			log.Fatalf("provided module: [%s] is not runnable\n", modules[index])
		}

		// stop each module
		err := module.Stop()
		if err != nil {
			log.Printf("failed to stop module: [%s] due: [%s]\n", modules[index], err)
			continue
		}

		log.Printf("module: [%s] stopped\n", modules[index])
	}
}

func (c *Containment) SetModuleConfig(name string, value domain.Config) {
	c.configs[name] = value
}

func (c *Containment) GetModuleConfig(name string) domain.Config {
	return c.configs[name]
}
