package executor

import (
	"bytes"
	"errors"
	"os/exec"
)

// Mapa de comandos permitidos: "nombre" -> "ruta absoluta en linux"
var AllowedCommands = map[string]string{
	"uptime": "/usr/bin/uptime",
	"free":   "/usr/bin/free",
	"df":     "/usr/bin/df",
}

// Execute ejecuta un comando si está en la lista blanca
func Execute(commandName string, params []string) (string, error) {
	path, exists := AllowedCommands[commandName]
	if !exists {
		return "", errors.New("comando no autorizado")
	}

	// exec.Command NO usa shell (/bin/sh), lo que evita inyecciones de tipo "ls ; rm -rf /"
	cmd := exec.Command(path, params...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return stderr.String(), err
	}

	return out.String(), nil
}
