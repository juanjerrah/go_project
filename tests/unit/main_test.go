package unit

import (
	"testing"
)

// TestMain pode ser usado para setup global de testes
func TestMain(m *testing.M) {
	// Setup global se necessário
	// ex: configurar timezone, conectar ao banco de teste, etc.
	
	m.Run()
	
	// Cleanup global se necessário
}