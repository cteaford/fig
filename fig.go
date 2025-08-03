
package fig

import (
    "fig/container"
    "fig/source"
)

type Entry struct {
	Name string
	Source string
	DefaultValue string
}

type Config struct {
	Entries []Entry
	RefreshInterval int64
}

func Configure(c Config) (container.Container, error) {
	// load sources
    source.Sources["env"] = source.EnvSource{}


	// load config
    container := container.Container{}
    container.Init()
    for _, entry := range c.Entries {
        // implement getting via source here
        source := source.Sources[entry.Source]
        value, err := source.Get(entry.Name)
        if err != nil || value == "" {
            value = entry.DefaultValue
        }
        container.Add(entry.Name, value)
    }

	//get a soda
	return container, nil
}
