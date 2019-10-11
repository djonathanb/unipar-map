package unipar

// Environment struct
type Environment struct {
	key  int
	name string
}

// Name getter for name property
func (e Environment) Name() string {
	return e.name
}

// Environments map of environments
var Environments = map[string]Environment{
	"ENTRADA_SAIDA_FRENTE":               Environment{key: 0, name: "Entrada/Saída Frente"}, // Térreo
	"GUARDA":                             Environment{key: 1, name: "Guarda"},
	"SECRETARIA":                         Environment{key: 2, name: "Secretária"},
	"PROVE":                              Environment{key: 3, name: "PROVE"},
	"COORDENACAO":                        Environment{key: 4, name: "Coordenação"},
	"RH":                                 Environment{key: 5, name: "RH"},
	"ESCRITORIO_MODELO":                  Environment{key: 6, name: "Escritório Modelo"},
	"CAIXA_ELETRONICO":                   Environment{key: 7, name: "Caixa Eletrônico"},
	"ENTRADA_SAIDA_AUDITORIO":            Environment{key: 8, name: "Entrada/Saída Auditório"},
	"ALMOXARIFADO":                       Environment{key: 9, name: "Almoxarifado"},
	"CORREDOR_ESTETICA":                  Environment{key: 10, name: "Corredor Estética"},
	"ESCADA_ESTETICA":                    Environment{key: 11, name: "Escada Estética"},
	"CORREDOR_ESTETICA_2":                Environment{key: 12, name: "Corredor Estética 2"},
	"RAMPA_ENTRADA_FRENTE":               Environment{key: 13, name: "Rampa - Entrada da Frente"},
	"RAMPA_ENTRADA_FRENTE_PATAMAR":       Environment{key: 14, name: "Rampa - Entrada da Frente (Patamar)"},
	"RAMPA_ENTRADA_FRENTE_1_ANDAR":       Environment{key: 15, name: "Rampa - Entrada da Frente (1° Andar)"},
	"BANHEIRO_SECRETARIA":                Environment{key: 16, name: "Banheiro Secretaria"},
	"ENTRADA_SAIDA_LATERAL":              Environment{key: 17, name: "Entrada/Saída Lateral"},
	"PATIO":                              Environment{key: 18, name: "Pátio"},
	"PATIO_2":                            Environment{key: 19, name: "Pátio Baixo"},
	"CORREDOR_ARQUITETURA":               Environment{key: 20, name: "Corredor Arquitetura"},
	"ESCADA_ARQUITETURA":                 Environment{key: 21, name: "Escada Arquitetura"},
	"CANTINA":                            Environment{key: 22, name: "Cantina"},
	"EAD":                                Environment{key: 23, name: "EAD"},
	"ENTRADA_SAIDA_BLOCO_PATIO":          Environment{key: 24, name: "Entrada/Saída Bloco/Pátio"}, // Sub
	"CORREDOR_FISIO":                     Environment{key: 25, name: "Corredor Fisioterapia"},
	"BANHEIRO_FISIO":                     Environment{key: 26, name: "Banheiro Fisioterapia"},
	"ENTRADA_SAIDA_BLOCO_ESTACIONAMENTO": Environment{key: 27, name: "Entrada/Saída Bloco/Estacionamento"},
	"ESTACIONAMENTO":                     Environment{key: 28, name: "Estacionamento"},
	"PORTAO_ESTACIONAMENTO":              Environment{key: 29, name: "Portão Estacionamento"},
	"QUIOSQUE":                           Environment{key: 30, name: "Quiosque"},
	"LAGO":                               Environment{key: 31, name: "Lago"},
	"SALA_PROFESSORES":                   Environment{key: 32, name: "Sala dos Professores"},
	"ANFITEATRO":                         Environment{key: 33, name: "Anfiteatro"},
	"PORTAO_ANFITEATRO":                  Environment{key: 34, name: "Portão Anfiteatro"},
	"ESCADA_BIBLIOTECA_RAMPA":            Environment{key: 35, name: "Escada da Biblioteca/Rampa"}, // 1 Andar
	"ESCADA_BIBLIOTECA_PATIO":            Environment{key: 36, name: "Escada da Biblioteca/Patio"},
	"BIBLIOTECA":                         Environment{key: 37, name: "Biblioteca"},
	"DIRETORIA":                          Environment{key: 38, name: "Diretoria"},
	"FIES":                               Environment{key: 39, name: "FIES"},
	"LABORATORIOS_SISTEMAS":              Environment{key: 40, name: "Laboratórios de Sistemas"},
	"ATLETICA":                           Environment{key: 41, name: "Atlética (Lobão)"},
	"RAMPA_LABORATORIOS":                 Environment{key: 42, name: "Rampa dos Laboratórios"},
	"BANHEIRO_LABORATORIOS":              Environment{key: 43, name: "Banheiro dos Laboratórios"},
	"LABORATORIOS_SISTEMAS_2":            Environment{key: 44, name: "Laboratórios de Sistemas 2"},
	"TECNICA":                            Environment{key: 45, name: "Técnica"},
	"CORREDOR_DIREITO":                   Environment{key: 46, name: "Corredor Direito"},
	"BANHEIRO_DIREITO":                   Environment{key: 47, name: "Banheiro Direito"},
	"ESCADA_DIREITO":                     Environment{key: 48, name: "Escada Direito"},
	"CORREDOR_DIREITO_2":                 Environment{key: 49, name: "Corredor Direito 2"},
	"CEUP":                               Environment{key: 50, name: "CEUP"},
	"BANHEIRO_CEUP":                      Environment{key: 51, name: "Banheiro CEUP"},
}
