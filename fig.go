
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
    source.RegisterSource("env", source.EnvSource{})

    container := container.Container{}
    container.Init()
    for _, entry := range c.Entries {
        source := source.Sources[entry.Source]
        value, err := source.Get(entry.Name)
        if err != nil || value == "" {
            value = entry.DefaultValue
        }
        container.Add(entry.Name, value)
    }

	return container, nil
}
