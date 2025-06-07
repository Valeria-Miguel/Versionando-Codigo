# Versionando Codigo

API RESTful desarrollada con Go y Fiber para la gestión de usuarios y tareas, utilizando autenticación JWT y almacenamiento en MongoDB Atlas.

---

## Tecnologías usadas

- **Go**
- **Fiber** (Framework web)
- **MongoDB Atlas**
- **JWT** (Autenticación segura)
- **bcrypt** (Hash de contraseñas)

---

## Estructura del Proyecto

├── main.go → Punto de entrada
├── go.mod → Módulo de Go
├── config/ → Configuración de MongoDB
├── models/ → Modelos (User, Task)
├── handlers/ → Lógica de endpoints
├── routes/ → Definición de rutas
├── middleware/ → Middleware JWT
├── utils/ → Funciones auxiliares (JWT)
├── test/ → Pruebas automatizadas (Aun no implementada)
└── .env → Variables de entorno



---

## Autenticación

- Los usuarios se registran con contraseña hasheada (`bcrypt`)
- Los tokens JWT tienen duracion de **10 minutos**
- Las rutas de tareas estan protegidas mediante middleware

---

## Endpoints principales

### Auth
- `POST /api/users/register` – Registro
- `POST /api/users/login` – Login y obtención de token

### Users (requieren token)
- `GET /api/users` – Listar tareas del usuario autenticado
- `POST /api/users/get` - Ver solo un usuario se envia por json el ID 
{
   "id": "68426e0c8dd5602d4960730e"
}
- `PUT /api/users/update` – Actualizar usuario por ID enviado en json
- `DELETE /api/users/delete` – Eliminar usuario por ID 

### Tasks (requieren token)
- `GET /api/tasks/tasks` – Listar tareas del usuario autenticado
- `POST /api/tasks/create` – Crear tarea
- `POST /api/tasks/get` - Ver solo un tarea se envia por json el ID 
{
   "id": "68426e0c8dd5602d4960730e"
}
- `PUT /api/tasks/update` – Actualizar tarea por ID
- `DELETE /api/tasks/delete` – Eliminar tarea por ID

---

##  Configuración `.env`

```env
se encuentra la clave secreta y  URL de conexión del clúster en MongoDB Atlas
