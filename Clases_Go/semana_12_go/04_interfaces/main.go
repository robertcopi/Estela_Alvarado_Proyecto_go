package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

// ==========================================
// PASO 1: INTERFACES BÁSICAS
// ==========================================

// Notificador define la funcionalidad básica de envío
type Notificador interface {
	EnviarNotificacion(destinatario, mensaje string) error
}

// ValidadorMensaje valida contenido antes del envío
type ValidadorMensaje interface {
	ValidarMensaje(mensaje string) error
	ValidarDestinatario(destinatario string) error
}

// Rastreador permite hacer seguimiento de notificaciones
type Rastreador interface {
	ObtenerEstado(id string) (string, error)
	ObtenerEstadisticas() map[string]int
}

// Logger registra eventos del sistema
type Logger interface {
	Log(nivel, mensaje string)
	LogError(error)
	LogInfo(string)
}

// ==========================================
// PASO 2: INTERFACES COMPUESTAS
// ==========================================

// NotificadorCompleto combina funcionalidades básicas
type NotificadorCompleto interface {
	Notificador
	ValidadorMensaje
}

// NotificadorAvanzado incluye todas las funcionalidades
type NotificadorAvanzado interface {
	Notificador
	ValidadorMensaje
	Rastreador
	Logger
}

// ==========================================
// PASO 3: STRUCTS Y TIPOS DE DATOS
// ==========================================

type TipoNotificacion string

const (
	Email TipoNotificacion = "email"
	SMS   TipoNotificacion = "sms"
	Push  TipoNotificacion = "push"
	Slack TipoNotificacion = "slack"
)

type EstadoNotificacion string

const (
	Pendiente EstadoNotificacion = "pendiente"
	Enviada   EstadoNotificacion = "enviada"
	Fallida   EstadoNotificacion = "fallida"
	Entregada EstadoNotificacion = "entregada"
)

type RegistroNotificacion struct {
	ID           string
	Tipo         TipoNotificacion
	Destinatario string
	Mensaje      string
	Estado       EstadoNotificacion
	Timestamp    time.Time
	Intentos     int
	Error        string
}

type ConfiguracionNotificacion struct {
	MaxIntentos     int
	TimeoutSegundos int
	ReintentoAuto   bool
}

// ==========================================
// PASO 4: IMPLEMENTACIONES CONCRETAS
// ==========================================

// EmailNotificador - Implementa múltiples interfaces
type EmailNotificador struct {
	servidor      string
	puerto        int
	usuario       string
	password      string
	configuracion ConfiguracionNotificacion
	registros     map[string]*RegistroNotificacion
}

// Constructor para EmailNotificador
func NuevoEmailNotificador(servidor string, puerto int, usuario, password string) *EmailNotificador {
	return &EmailNotificador{
		servidor: servidor,
		puerto:   puerto,
		usuario:  usuario,
		password: password,
		configuracion: ConfiguracionNotificacion{
			MaxIntentos:     3,
			TimeoutSegundos: 30,
			ReintentoAuto:   true,
		},
		registros: make(map[string]*RegistroNotificacion),
	}
}

// Implementa Notificador
func (e *EmailNotificador) EnviarNotificacion(destinatario, mensaje string) error {
	if err := e.ValidarDestinatario(destinatario); err != nil {
		return err
	}
	if err := e.ValidarMensaje(mensaje); err != nil {
		return err
	}

	id := fmt.Sprintf("email_%d", time.Now().UnixNano())
	registro := &RegistroNotificacion{
		ID:           id,
		Tipo:         Email,
		Destinatario: destinatario,
		Mensaje:      mensaje,
		Estado:       Pendiente,
		Timestamp:    time.Now(),
		Intentos:     1,
	}

	e.registros[id] = registro
	e.LogInfo(fmt.Sprintf("Enviando email a %s", destinatario))
	time.Sleep(5 * time.Millisecond)

	// Simular éxito (podemos falsearlo)
	registro.Estado = Enviada
	e.LogInfo(fmt.Sprintf("Email enviado exitosamente: %s", id))
	return nil
}

// Implementa ValidadorMensaje
func (e *EmailNotificador) ValidarMensaje(mensaje string) error {
	if len(mensaje) == 0 {
		return errors.New("mensaje no puede estar vacío")
	}
	if len(mensaje) > 1000 {
		return errors.New("mensaje muy largo (máximo 1000 caracteres)")
	}
	return nil
}

// Implementa ValidadorMensaje
func (e *EmailNotificador) ValidarDestinatario(destinatario string) error {
	if !strings.Contains(destinatario, "@") || !strings.Contains(destinatario, ".") {
		return errors.New("email inválido")
	}
	return nil
}

// Implementa Rastreador
func (e *EmailNotificador) ObtenerEstado(id string) (string, error) {
	if registro, existe := e.registros[id]; existe {
		return string(registro.Estado), nil
	}
	return "", errors.New("notificación no encontrada")
}

func (e *EmailNotificador) ObtenerEstadisticas() map[string]int {
	stats := map[string]int{
		"total":      0,
		"enviadas":   0,
		"fallidas":   0,
		"pendientes": 0,
	}

	for _, registro := range e.registros {
		stats["total"]++
		switch registro.Estado {
		case Enviada:
			stats["enviadas"]++
		case Fallida:
			stats["fallidas"]++
		case Pendiente:
			stats["pendientes"]++
		}
	}
	return stats
}

// Implementa Logger
func (e *EmailNotificador) Log(nivel, mensaje string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] EMAIL [%s]: %s\n", timestamp, nivel, mensaje)
}

func (e *EmailNotificador) LogError(err error) {
	e.Log("ERROR", err.Error())
}

func (e *EmailNotificador) LogInfo(mensaje string) {
	e.Log("INFO", mensaje)
}

// ==========================================

// SMSNotificador - Otra implementación
type SMSNotificador struct {
	apiKey    string
	proveedor string
	registros map[string]*RegistroNotificacion
}

func NuevoSMSNotificador(apiKey, proveedor string) *SMSNotificador {
	return &SMSNotificador{
		apiKey:    apiKey,
		proveedor: proveedor,
		registros: make(map[string]*RegistroNotificacion),
	}
}

// Implementa Notificador
func (s *SMSNotificador) EnviarNotificacion(destinatario, mensaje string) error {
	id := fmt.Sprintf("sms_%d", time.Now().UnixNano())
	registro := &RegistroNotificacion{
		ID:           id,
		Tipo:         SMS,
		Destinatario: destinatario,
		Mensaje:      mensaje,
		Estado:       Enviada,
		Timestamp:    time.Now(),
		Intentos:     1,
	}
	s.registros[id] = registro
	s.LogInfo(fmt.Sprintf("Enviando SMS a %s via %s", destinatario, s.proveedor))
	return nil
}

func (s *SMSNotificador) LogInfo(mensaje string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] SMS [INFO]: %s\n", timestamp, mensaje)
}

// SlackNotificador - Implementación más simple
type SlackNotificador struct {
	webhook string
	canal   string
}

func NuevoSlackNotificador(webhook, canal string) *SlackNotificador {
	return &SlackNotificador{
		webhook: webhook,
		canal:   canal,
	}
}

// Solo implementa Notificador
func (sl *SlackNotificador) EnviarNotificacion(destinatario, mensaje string) error {
	fmt.Printf("🔔 Slack -> Canal: %s | Usuario: %s | Mensaje: %s\n",
		sl.canal, destinatario, mensaje)
	time.Sleep(2 * time.Millisecond)
	return nil
}

// ==========================================
// PASO 5: SERVICIO PRINCIPAL
// ==========================================

type ServicioNotificaciones struct {
	notificadores []Notificador
	logger        Logger
}

func NuevoServicioNotificaciones() *ServicioNotificaciones {
	return &ServicioNotificaciones{
		notificadores: make([]Notificador, 0),
	}
}

func (sn *ServicioNotificaciones) AgregarNotificador(notificador Notificador) {
	sn.notificadores = append(sn.notificadores, notificador)
}

func (sn *ServicioNotificaciones) EstablecerLogger(logger Logger) {
	sn.logger = logger
}

// Enviar a todos los notificadores
func (sn *ServicioNotificaciones) EnviarATodos(destinatario, mensaje string) map[string]error {
	resultados := make(map[string]error)
	for _, notificador := range sn.notificadores {
		tipoNotificador := fmt.Sprintf("%T", notificador)
		err := notificador.EnviarNotificacion(destinatario, mensaje)
		resultados[tipoNotificador] = err
	}
	return resultados
}

// Type switch para manejar diferentes tipos
func ProcesarNotificadorPorTipo(n Notificador, destinatario, mensaje string) {
	switch notificador := n.(type) {
	case *EmailNotificador:
		fmt.Println("📧 Procesando como EmailNotificador...")
		fmt.Printf("   Servidor: %s:%d\n", notificador.servidor, notificador.puerto)
		notificador.EnviarNotificacion(destinatario, mensaje)
	case *SMSNotificador:
		fmt.Println("📱 Procesando como SMSNotificador...")
		fmt.Printf("   Proveedor: %s\n", notificador.proveedor)
		notificador.EnviarNotificacion(destinatario, mensaje)
	case *SlackNotificador:
		fmt.Println("💬 Procesando como SlackNotificador...")
		fmt.Printf("   Canal: %s\n", notificador.canal)
		notificador.EnviarNotificacion(destinatario, mensaje)
	default:
		fmt.Printf("❓ Tipo desconocido: %T\n", notificador)
		notificador.EnviarNotificacion(destinatario, mensaje)
	}
}

// ==========================================
// PASO 7: FUNCIÓN PRINCIPAL DEMOSTRATIVA
// ==========================================

func main() {
	fmt.Println("🔔 SISTEMA DE NOTIFICACIONES - INTERFACES EN ACCIÓN")
	fmt.Println("============================================================")

	email := NuevoEmailNotificador("smtp.gmail.com", 587, "app@empresa.com", "password")
	sms := NuevoSMSNotificador("api-key-123", "Twilio")
	slack := NuevoSlackNotificador("https://hooks.slack.com/...", "#general")

	servicio := NuevoServicioNotificaciones()
	servicio.EstablecerLogger(email) // Email también funciona como logger

	servicio.AgregarNotificador(email)
	servicio.AgregarNotificador(sms)
	servicio.AgregarNotificador(slack)

	fmt.Println("\n📋 1. POLIMORFISMO BÁSICO:")
	fmt.Println("----------------------------------------")
	notificadores := []Notificador{email, sms, slack}
	for _, n := range notificadores {
		fmt.Printf("\n🧪 Probando %T:\n", n)
		err := n.EnviarNotificacion("usuario@ejemplo.com", "¡Hola desde Go!")
		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
		}
	}

	fmt.Println("\n📋 2. TYPE SWITCH EN ACCIÓN:")
	fmt.Println("----------------------------------------")
	for _, n := range notificadores {
		ProcesarNotificadorPorTipo(n, "+54911234567", "Mensaje tipo específico")
	}

	fmt.Println("\n📋 3. ESTADÍSTICAS Y TYPE ASSERTIONS:")
	fmt.Println("----------------------------------------")
	for _, n := range notificadores {
		if rastreador, implementa := n.(Rastreador); implementa {
			nombre := fmt.Sprintf("%T", n)
			stats := rastreador.ObtenerEstadisticas()
			fmt.Printf("📊 Estadísticas de %s:\n", nombre)
			statsJSON, _ := json.MarshalIndent(stats, "    ", "   ")
			fmt.Printf("    %s\n\n", string(statsJSON))
		}
	}
}
