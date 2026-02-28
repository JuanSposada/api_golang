package executor

import (
	"bytes"
	"errors"
	"os/exec"
	"os/user"
	"strconv"
	"syscall"
)

// Mapa de comandos permitidos (Whitelist)
var AllowedCommands = map[string]string{
	"uptime": "/usr/bin/uptime",
	"free":   "/usr/bin/free",
	"df":     "/usr/bin/df",
	"ls":     "/bin/ls",
	"whoami": "/usr/bin/whoami",
	"pwd":    "/bin/pwd",
	"ps":     "/bin/ps",
	"cat":    "/bin/cat",
}

func ExecuteAsUser(commandName string, params []string, username string) (string, error) {
	path, exists := AllowedCommands[commandName]
	if !exists {
		return "", errors.New("comando no autorizado en la whitelist")
	}

	// 1. Buscamos al usuario en el sistema operativo para obtener su UID y GID
	u, err := user.Lookup(username)
	if err != nil {
		return "", err
	}

	uid, _ := strconv.Atoi(u.Uid)
	gid, _ := strconv.Atoi(u.Gid)

	cmd := exec.Command(path, params...)

	// 2. ¡MAGIA! Le decimos al comando que se ejecute con la identidad de ese usuario
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.SysProcAttr.Credential = &syscall.Credential{
		Uid: uint32(uid),
		Gid: uint32(gid),
	}

	// Buffers para capturar la salida (Stdout y Stderr)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err = cmd.Run()

	// Si hay error, devolvemos lo que diga la consola (stderr)
	if err != nil {
		if stderr.Len() > 0 {
			return stderr.String(), err
		}
		return "", err
	}

	return out.String(), nil
}
