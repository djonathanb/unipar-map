package unipar

// Environment struct
type Environment struct {
	key   int
	name  string
	floor float32
}

// Name getter for name property
func (e Environment) Name() string {
	return e.name
}

// Environments map of environments
var Environments = map[string]Environment{
	"ENTRADA_SAIDA_FRENTE":               Environment{key: 0, name: "Entrada/Saída Frente", floor: 0}, // Térreo
	"GUARDA":                             Environment{key: 1, name: "Guarda", floor: 0},
	"SECRETARIA":                         Environment{key: 2, name: "Secretária", floor: 0},
	"PROVE":                              Environment{key: 3, name: "PROVE", floor: 0},
	"COORDENACAO":                        Environment{key: 4, name: "Coordenação", floor: 0},
	"RH":                                 Environment{key: 5, name: "RH", floor: 0},
	"ESCRITORIO_MODELO":                  Environment{key: 6, name: "Escritório Modelo", floor: 0},
	"CAIXA_ELETRONICO":                   Environment{key: 7, name: "Caixa Eletrônico", floor: 0},
	"ENTRADA_SAIDA_ANFITEATRO":           Environment{key: 8, name: "Entrada/Saída Anfiteatro", floor: 0},
	"ALMOXARIFADO":                       Environment{key: 9, name: "Almoxarifado", floor: 0},
	"CORREDOR_ESTETICA_SUL":              Environment{key: 10, name: "Corredor Estética (Sul)", floor: 0},
	"ESCADA_ESTETICA":                    Environment{key: 11, name: "Escada Estética", floor: 0},
	"CORREDOR_ESTETICA_NORTE":            Environment{key: 12, name: "Corredor Estética (Norte)", floor: 0},
	"RAMPA_ENTRADA_FRENTE":               Environment{key: 13, name: "Rampa - Entrada da Frente", floor: 0},
	"RAMPA_ENTRADA_FRENTE_PATAMAR":       Environment{key: 14, name: "Rampa - Entrada da Frente (Patamar)", floor: 0.5},
	"BANHEIRO_SECRETARIA":                Environment{key: 15, name: "Banheiro Secretaria", floor: 0},
	"ENTRADA_SAIDA_LATERAL":              Environment{key: 16, name: "Entrada/Saída Lateral", floor: 0},
	"PATIO_ALTO":                         Environment{key: 17, name: "Pátio (Alto)", floor: 0},
	"PATIO_BAIXO":                        Environment{key: 18, name: "Pátio (Baixo)", floor: 0},
	"CORREDOR_ARQUITETURA":               Environment{key: 19, name: "Corredor Arquitetura", floor: 0},
	"ESCADA_ARQUITETURA":                 Environment{key: 20, name: "Escada Arquitetura", floor: 0},
	"CANTINA":                            Environment{key: 21, name: "Cantina", floor: -1},
	"EAD":                                Environment{key: 22, name: "EAD", floor: -1},
	"ENTRADA_SAIDA_BLOCO_PATIO":          Environment{key: 23, name: "Entrada/Saída Bloco/Pátio", floor: -1}, // Sub
	"CORREDOR_FISIO":                     Environment{key: 24, name: "Corredor Fisioterapia", floor: -1},
	"BANHEIRO_FISIO":                     Environment{key: 25, name: "Banheiro Fisioterapia", floor: -1},
	"ENTRADA_SAIDA_BLOCO_ESTACIONAMENTO": Environment{key: 26, name: "Entrada/Saída Bloco/Estacionamento", floor: -1},
	"ESTACIONAMENTO":                     Environment{key: 27, name: "Estacionamento", floor: -1},
	"PORTAO_ESTACIONAMENTO":              Environment{key: 28, name: "Portão Estacionamento", floor: -1},
	"QUIOSQUE":                           Environment{key: 29, name: "Quiosque", floor: -1},
	"LAGO":                               Environment{key: 30, name: "Lago", floor: -1},
	"SALA_PROFESSORES":                   Environment{key: 31, name: "Sala dos Professores", floor: -1},
	"ANFITEATRO":                         Environment{key: 32, name: "Anfiteatro", floor: -1},
	"PORTAO_ANFITEATRO":                  Environment{key: 33, name: "Portão Anfiteatro", floor: -1},
	"ESCADA_BIBLIOTECA_RAMPA":            Environment{key: 34, name: "Escada da Biblioteca/Rampa", floor: 1}, // 1 Andar
	"ESCADA_BIBLIOTECA_PATIO":            Environment{key: 35, name: "Escada da Biblioteca/Patio", floor: 1},
	"BIBLIOTECA":                         Environment{key: 36, name: "Biblioteca", floor: 1},
	"DIRETORIA":                          Environment{key: 37, name: "Diretoria", floor: 1},
	"RAMPA_DIRETORIA":                    Environment{key: 38, name: "Rampa - Diretoria", floor: 1},
	"FIES":                               Environment{key: 39, name: "FIES", floor: 1},
	"LABORATORIOS_SISTEMAS_LESTE":        Environment{key: 40, name: "Laboratórios de Sistemas (Leste)", floor: 1},
	"ATLETICA":                           Environment{key: 41, name: "Atlética (Lobão)", floor: 1},
	"RAMPA_LABORATORIOS":                 Environment{key: 42, name: "Rampa dos Laboratórios", floor: 1},
	"BANHEIRO_LABORATORIOS":              Environment{key: 43, name: "Banheiro dos Laboratórios", floor: 1},
	"LABORATORIOS_SISTEMAS_OESTE":        Environment{key: 44, name: "Laboratórios de Sistemas (Oeste)", floor: 1},
	"TECNICA":                            Environment{key: 45, name: "Técnica", floor: 1},
	"CORREDOR_DIREITO_SUL":               Environment{key: 46, name: "Corredor Direito (Sul)", floor: 1},
	"BANHEIRO_DIREITO_NORTE":             Environment{key: 47, name: "Banheiro Direito (Norte)", floor: 1},
	"ESCADA_DIREITO":                     Environment{key: 48, name: "Escada Direito", floor: 1},
	"CORREDOR_DIREITO_2":                 Environment{key: 49, name: "Corredor Direito 2", floor: 1},
	"CEUP":                               Environment{key: 50, name: "CEUP", floor: 1},
	"BANHEIRO_CEUP":                      Environment{key: 51, name: "Banheiro CEUP", floor: 1},
}

// Edge edge between two environments
type Edge struct {
	from      Environment
	to        Environment
	direction string
	distance  int
}

// Edges list of edges between environments
var Edges = []Edge{
	Edge{from: Environments["ENTRADA_SAIDA_FRENTE"], to: Environments["SECRETARIA"], direction: "N", distance: 5},
	Edge{from: Environments["SECRETARIA"], to: Environments["BANHEIRO_SECRETARIA"], direction: "N", distance: 15},
	Edge{from: Environments["SECRETARIA"], to: Environments["PROVE"], direction: "W", distance: 5},
	Edge{from: Environments["PROVE"], to: Environments["RAMPA_ENTRADA_FRENTE"], direction: "N", distance: 5},
	Edge{from: Environments["PROVE"], to: Environments["GUARDA"], direction: "W", distance: 5},
	Edge{from: Environments["GUARDA"], to: Environments["COORDENACAO"], direction: "W", distance: 2},
	Edge{from: Environments["COORDENACAO"], to: Environments["MECANOGRAFIA"], direction: "W", distance: 5},
	Edge{from: Environments["COORDENACAO"], to: Environments["RH"], direction: "W", distance: 24},
	Edge{from: Environments["RH"], to: Environments["ESCRITORO_MODELO"], direction: "W", distance: 5},
	Edge{from: Environments["ESCRITORIO_MODELO"], to: Environments["CAIXA_ELETRONICO"], direction: "W", distance: 4},
	Edge{from: Environments["CAIXA_ELETRONICO"], to: Environments["ENTRADA_SAIDA_ANFITEATRO"], direction: "W", distance: 10},
	Edge{from: Environments["CAIXA_ELETRONICO"], to: Environments["ALMOXARIFADO"], direction: "N", distance: 5},
	Edge{from: Environments["ALMOXARIFADO"], to: Environments["CORREDOR_ESTETICA_SUL"], direction: "N", distance: 15},
	Edge{from: Environments["CORREDOR_ESTETICA_SUL"], to: Environments["ESCADA_ESTETICA"], direction: "N", distance: 15},
	Edge{from: Environments["ESCADA_ESTETICA"], to: Environments["CORREDOR_ESTETICA_NORTE"], direction: "N", distance: 15},
	Edge{from: Environments["ESCADA_ESTETICA"], to: Environments["ENTRADA_SAIDA_BLOCO_PATIO"], direction: "N", distance: 5},
	Edge{from: Environments["MECANOGRAFIA"], to: Environments["PATIO_ALTO"], direction: "N", distance: 20},
	Edge{from: Environments["RAMPA_ENTRADA_FRENTE"], to: Environments["RAMPA_ENTRADA_FRENTE_PATAMAR"], direction: "N", distance: 7},    // 0 to 1
	Edge{from: Environments["RAMPA_ENTRADA_FRENTE_PATAMAR"], to: Environments["DIRETORIA"], direction: "S", distance: 7},               // 0 to 1
	Edge{from: Environments["RAMPA_ENTRADA_FRENTE_PATAMAR"], to: Environments["ESCADA_BIBLIOTECA_RAMPA"], direction: "N", distance: 1}, // 0 to 1/
	Edge{from: Environments["ESCADA_BIBLIOTECA_RAMPA"], to: Environments["BIBLIOTECA"], direction: "N", distance: 3},                   // 0 to 1
	Edge{from: Environments["ENTRADA_SAIDA_LATERAL"], to: Environments["BANHEIRO_SECRETARIA"], direction: "S", distance: 5},
	Edge{from: Environments["ENTRADA_SAIDA_LATERAL"], to: Environments["CORREDOR_ARQUITETURA"], direction: "N", distance: 15},
	Edge{from: Environments["BANHEIRO_SECRETARIA"], to: Environments["CORREDOR_ARQUITETURA"], direction: "N", distance: 15},
	Edge{from: Environments["PATIO_ALTO"], to: Environments["ENTRADA_SAIDA_BLOCO_ARQUITETURA"], direction: "E", distance: 10},
	Edge{from: Environments["PATIO_ALTO"], to: Environments["ESCADA_BIBLIOTECA_PATIO"], direction: "E", distance: 10},
	Edge{from: Environments["PATIO_ALTO"], to: Environments["PATIO_BAIXO"], direction: "N", distance: 15},
	Edge{from: Environments["ENTRADA_SAIDA_BLOCO_ARQUITETURA"], to: Environments["CORREDOR_ARQUITETURA"], direction: "E", distance: 10},
	Edge{from: Environments["ENTRADA_SAIDA_BLOCO_ARQUITETURA"], to: Environments["BANHEIRO_SECRETARIA"], direction: "E", distance: 10},
	Edge{from: Environments["CORREDOR_ARQUITETURA"], to: Environments["ESCADA_ARQUITETURA"], direction: "N", distance: 15},
	Edge{from: Environments["ESCADA_ARQUITETURA"], to: Environments["CANTINA"], direction: "N", distance: 8},
	Edge{from: Environments["CANTINA"], to: Environments["EAD"], direction: "N", distance: 8},
	Edge{from: Environments["CANTINA"], to: Environments["PATIO_BAIXO"], direction: "W", distance: 8},
	Edge{from: Environments["ENTRADA_SAIDA_BLOCO_PATIO"], to: Environments["PATIO_BAIXO"], direction: "E", distance: 20},    // 0 to -1
	Edge{from: Environments["ENTRADA_SAIDA_BLOCO_PATIO"], to: Environments["ESCADA_ESTETICA"], direction: "W", distance: 5}, // 0 to -1
	Edge{from: Environments["ENTRADA_SAIDA_BLOCO_PATIO"], to: Environments["CORREDOR_FISIO"], direction: "N", distance: 15},
	Edge{from: Environments["ENTRADA_SAIDA_BLOCO_PATIO"], to: Environments["BANHEIRO_FISIO"], direction: "W", distance: 15},
	Edge{from: Environments["ENTRADA_SAIDA_BLOCO_ESTACIONAMENTO"], to: Environments["BANHEIRO_FISIO"], direction: "E", distance: 5},
	Edge{from: Environments["ENTRADA_SAIDA_BLOCO_ESTACIONAMENTO"], to: Environments["BANHEIRO_FISIO"], direction: "E", distance: 5},
	Edge{from: Environments["ENTRADA_SAIDA_BLOCO_ESTACIONAMENTO"], to: Environments["ESTACIONAMENTO"], direction: "N", distance: 25},
	Edge{from: Environments["ENTRADA_SAIDA_BLOCO_ESTACIONAMENTO"], to: Environments["SALA_PROFESSORES"], direction: "W", distance: 7},
	Edge{from: Environments["ESTACIONAMENTO"], to: Environments["SALA_PROFESSORES"], direction: "S", distance: 20},
	Edge{from: Environments["ESTACIONAMENTO"], to: Environments["LAGO"], direction: "SW", distance: 25},
	Edge{from: Environments["ESTACIONAMENTO"], to: Environments["QUIOESQUE"], direction: "W", distance: 25},
	Edge{from: Environments["ESTACIONAMENTO"], to: Environments["PORTAO_ESTACIONAMENTO"], direction: "N", distance: 15},
	Edge{from: Environments["QUIOESQUE"], to: Environments["LAGO"], direction: "N", distance: 5},
	Edge{from: Environments["ANFITEATRO"], to: Environments["SALA_PROFESSORES"], direction: "N", distance: 40},
	Edge{from: Environments["ANFITEATRO"], to: Environments["PORTAO_ANFITEATRO"], direction: "S", distance: 10},
	Edge{from: Environments["PORTAO_ANFITEATRO"], to: Environments["ENTRADA_SAIDA_ANFITEATRO"], direction: "E", distance: 10},
	Edge{from: Environments["DIRETORIA"], to: Environments["FIES"], direction: "E", distance: 10},
	Edge{from: Environments["DIRETORIA"], to: Environments["LABORATORIOS_SISTEMAS_LESTE"], direction: "W", distance: 20},
	Edge{from: Environments["FIES"], to: Environments["BIBLIOTECA"], direction: "N", distance: 30},
	Edge{from: Environments["BIBLIOTECA"], to: Environments["ESCADA_BIBLIOTECA_PATIO"], direction: "W", distance: 10},
	Edge{from: Environments["LABORATORIOS_SISTEMAS_LESTE"], to: Environments["ATLETICA"], direction: "W", distance: 20},
	Edge{from: Environments["ATLETICA"], to: Environments["CORREDOR_DIREITO_SUL"], direction: "N", distance: 20},
	Edge{from: Environments["ATLETICA"], to: Environments["BANHEIRO_LABORATORIOS"], direction: "W", distance: 25},
	Edge{from: Environments["CORREDOR_DIREITO_SUL"], to: Environments["BANHEIRO_DIREITO"], direction: "N", distance: 5},
	Edge{from: Environments["BANHEIRO_DIREITO"], to: Environments["ESCADA_DIREITO"], direction: "N", distance: 10},
	Edge{from: Environments["ESCADA_DIREITO"], to: Environments["CORREDOR_DIREITO_NORTE"], direction: "N", distance: 20},
	Edge{from: Environments["CORREDOR_DIREITO_NORTE"], to: Environments["CEUP"], direction: "N", distance: 20},
	Edge{from: Environments["CEUP"], to: Environments["BANHEIRO_CEUP"], direction: "E", distance: 5},
	Edge{from: Environments["BANHEIRO_LABORATORIOS"], to: Environments["LABORATORIOS_SISTEMAS_OESTE"], direction: "W", distance: 5},
	Edge{from: Environments["LABORATORIOS_SISTEMAS_OESTE"], to: Environments["TECNICA"], direction: "W", distance: 10},
	Edge{from: Environments["TECNICA"], to: Environments["TECNICA"], direction: "W", distance: 10},
	Edge{from: Environments["ESCADA_DIREITO"], to: Environments["ESCADA_ESTETICA"], direction: "W", distance: 7}, // 1 to 0
}
