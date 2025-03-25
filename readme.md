go get github.com/google/uuid


```
myapp/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── core/
│   │   ├── domain/           # Entidades de negocio
│   │   │   └── user.go
│   │   ├── ports/            # Interfaces (contratos)
│   │   │   ├── repositories.go
│   │   │   └── services.go
│   │   └── services/         # Lógica de negocio
│   │       └── user_service.go
│   ├── features/             # Vertical slicing por funcionalidad
│   │   ├── users/            # Feature: Users
│   │   │   ├── transport/    # Adaptadores de entrada
│   │   │   │   ├── http/
│   │   │   │   │   └── handlers.go
│   │   │   │   └── websocket/
│   │   │   │       └── handlers.go
│   │   │   └── repository/   # Adaptadores de salida
│   │   │       └── user_repository.go
│   │   └── chat/             # Feature: Chat
│   │       ├── transport/
│   │       │   ├── http/
│   │       │   └── websocket/
│   │       └── repository/
│   ├── infrastructure/
│   │   ├── database/         # Implementación DB
│   │   ├── cache/            # Implementación Redis
│   │   └── websocket/        # Config WS server
│   ├── shared/
│   │   ├── middleware/       # Middlewares HTTP
│   │   ├── utils/            # Helpers comunes
│   │   └── config/           # Configuración
│   └── app/                  # Inicialización y DI
│       └── bootstrap.go
├── pkg/                      # Librerías reutilizables
│   ├── validator/
│   └── logger/
├── api/
│   └── openapi.yaml          # Especificación OpenAPI
├── scripts/                  # Migraciones, scripts
├── test/
│   ├── integration/
│   └── e2e/
├── deployments/              # Docker, k8s
└── go.mod
```

SECUENCIA DE EJECUCIÓN
→ Dominio (Estructuras de datos) 
→ Output Ports (Contratos hacia el exterior) 
→ Input Ports (Casos de uso) 
→ Implementaciones

DIAGRAMA DE INYECCIÓN DE DEPENDENCIAS
graph LR
    main.go --> Bootstrap
    Bootstrap -->|Crea| UserRepository
    Bootstrap -->|Crea| UserService
    UserService -->|Depende de| UserRepository
    main.go -->|Usa| UserService
