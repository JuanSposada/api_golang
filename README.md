# Secure Linux Command Executor API (Golang)

Esta es una API REST robusta construida en **Go** diseñada para ejecutar comandos específicos de Linux en el servidor. El proyecto prioriza la seguridad, utilizando autenticación nativa del sistema operativo y mecanismos para prevenir la inyección de comandos.

## 🛡️ Controles de Seguridad Implementados

Para cumplir con el objetivo de prevenir vulnerabilidades, se implementaron los siguientes controles:

1.  **Whitelist de Comandos:** No se permite la ejecución de cualquier comando. Solo aquellos definidos explícitamente en el mapa de configuración (`internal/executor/executor.go`) son permitidos.
2.  **Prevención de Command Injection:** No se utiliza una shell intermedia (`/bin/sh` o `/bin/bash`). Se utiliza el paquete `os/exec` de Go, que trata los comandos y parámetros como argumentos aislados, neutralizando intentos de concatenación (ej. `;`, `&&`, `|`).
3.  **Autenticación Linux (PAM):** La API valida las credenciales (Basic Auth) directamente contra los usuarios reales del sistema operativo utilizando **PAM (Pluggable Authentication Modules)**.
4.  **Principio de Menor Privilegio:** Diseñada para ejecutarse bajo un usuario con permisos limitados.

## 🚀 Requisitos del Sistema

* **Go:** 1.20 o superior.
* **Dependencias de Linux:** Es necesario instalar las cabeceras de desarrollo de PAM para la compilación:
    ```bash
    sudo apt-get install libpam0g-dev
    ```

## 🛠️ Instalación y Configuración

1. **Clonar el repositorio:**
   ```bash
   git clone [https://github.com/tu-usuario/go-secure-api.git](https://github.com/tu-usuario/go-secure-api.git)
   cd go-secure-api

**Instalar dependencias de Go:**  
```bash
go mod tidy
```
2. 

**Ejecutar el servidor:** Dado que PAM requiere acceso a los módulos de autenticación del sistema, es posible que necesites ejecutar con privilegios:  
```bash
sudo go run cmd/api/main.go
```
3. *El servidor iniciará en el puerto `:8080` por defecto.*

## **📖 Guía de Uso**

### **Endpoint Principal**

`POST /system/execute`

**Autenticación:** Requiere **Basic Auth** (Usuario y contraseña de un usuario válido en el servidor Linux).

#### **Cuerpo de la Petición (JSON):**

JSON
```
{

  "command": "uptime",

  "params": []

}
```

#### **Ejemplo con CURL:**
```bash
curl -i -X POST http://localhost:8080/system/execute \
     -u "tu_usuario:tu_contraseña" \
     -H "Content-Type: application/json" \
     -d '{"command": "uptime", "params": []}'
```
#### **Comandos Permitidos (Whitelist):**

Actualmente, la API permite:

* `uptime`  
* `free`  
* `df`

## **📁 Estructura del Proyecto**

```plaintext

├── cmd/

│   └── api/             \# Punto de entrada (main.go)

├── internal/

│   ├── executor/        \# Lógica de ejecución segura y whitelist

│   ├── handlers/        \# Handlers de la API y Middleware de Auth PAM

├── go.mod               \# Gestión de dependencias

└── README.md
```
---

Desarrollado con ❤️ en **Go** para aprendizaje de sistemas seguros.

