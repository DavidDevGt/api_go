CREATE DATABASE IF NOT EXISTS ticket_sys_db;
USE ticket_sys_db;

-- Tabla de departamentos de Guatemala
CREATE TABLE departamento (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL
);

INSERT INTO departamento (nombre) VALUES
    ('Alta Verapaz'),
    ('Baja Verapaz'),
    ('Chimaltenango'),
    ('Chiquimula'),
    ('El Progreso'),
    ('Escuintla'),
    ('Guatemala'),
    ('Huehuetenango'),
    ('Izabal'),
    ('Jalapa'),
    ('Jutiapa'),
    ('Petén'),
    ('Quetzaltenango'),
    ('Quiché'),
    ('Retalhuleu'),
    ('Sacatepéquez'),
    ('San Marcos'),
    ('Santa Rosa'),
    ('Sololá'),
    ('Suchitepéquez'),
    ('Totonicapán'),
    ('Zacapa');

-- Otras tablas existentes
CREATE TABLE rol (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    descripcion VARCHAR(255),
    active BOOLEAN DEFAULT TRUE
);

CREATE TABLE usuario (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(255) NOT NULL,
    correo VARCHAR(255) NOT NULL,
    contrasena VARCHAR(255) NOT NULL,
    rol_id INT,
    fecha_creacion DATETIME DEFAULT CURRENT_TIMESTAMP,
    usuario_creaion_id INT,
    fecha_edicion DATETIME,
    usuario_edicion_id INT,
    active BOOLEAN DEFAULT TRUE
);

CREATE TABLE cliente (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(255) NOT NULL,
    direccion TEXT,
    telefono VARCHAR(20),
    fecha_creacion DATETIME DEFAULT CURRENT_TIMESTAMP,
    usuario_creaion_id INT,
    fecha_edicion DATETIME,
    usuario_edicion_id INT,
    active BOOLEAN DEFAULT TRUE
);

CREATE TABLE ruta (
    id INT AUTO_INCREMENT PRIMARY KEY,
    origen VARCHAR(255) NOT NULL,
    destino VARCHAR(255) NOT NULL,
    distancia_km DECIMAL(10, 2) NOT NULL,
    duracion_horas DECIMAL(5, 2) NOT NULL,
    active BOOLEAN DEFAULT TRUE
);

CREATE TABLE viaje_cliente (
    id INT AUTO_INCREMENT PRIMARY KEY,
    viaje_id INT NOT NULL,
    cliente_id INT NOT NULL,
    fecha_viaje DATETIME NOT NULL,
    fecha_creacion DATETIME DEFAULT CURRENT_TIMESTAMP,
    usuario_creaion_id INT,
    fecha_edicion DATETIME,
    usuario_edicion_id INT,
    active BOOLEAN DEFAULT TRUE
);

CREATE TABLE ticket (
    id INT AUTO_INCREMENT PRIMARY KEY,
    titulo VARCHAR(255) NOT NULL,
    descripcion TEXT,
    precio DECIMAL(10, 2) NOT NULL,
    cliente_id INT NOT NULL,
    estado INT NOT NULL DEFAULT 1,
    fecha_creacion DATETIME DEFAULT CURRENT_TIMESTAMP,
    usuario_creaion_id INT,
    fecha_edicion DATETIME,
    usuario_edicion_id INT,
    active BOOLEAN NOT NULL DEFAULT TRUE
);