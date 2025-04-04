Ejecuta este comando en la carpeta /home/user/Desktop/2025/golang/struct:
-	go mod init struct

Esto generará un archivo go.mod que debe verse algo así:

	module struct

	go 1.20  # La versión de Go puede variar según tu instalación



struct/                   <- Carpeta raíz del módulo
├── go.mod                <- Archivo de módulo Go (lo crearemos)
├── 26_1struct.go         <- Archivo principal
└── mypackage/            <- Carpeta del paquete
    └── mypackage.go      <- Archivo del paquete

