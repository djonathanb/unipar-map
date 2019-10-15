package unipar

import (
	"log"
	"math"

	"github.com/djonathanb/unipar-map/ds"
)

// Environment struct
type Environment struct {
	Key        int     `json:"key"`
	Name       string  `json:"name"`
	Floor      float32 `json:"floor"`
	IsBathroom bool    `json:"-"`
}

// Environments map of environments
var Environments = map[string]Environment{
	"ENTRADA_SAIDA_FRENTE":               Environment{Key: 0, Name: "Entrada/Saída Frente", Floor: 0}, // Térreo
	"GUARDA":                             Environment{Key: 1, Name: "Guarda", Floor: 0},
	"SECRETARIA":                         Environment{Key: 2, Name: "Secretária", Floor: 0},
	"PROVE":                              Environment{Key: 3, Name: "PROVE", Floor: 0},
	"COORDENACAO":                        Environment{Key: 4, Name: "Coordenação", Floor: 0},
	"MECANOGRAFIA":                       Environment{Key: 4, Name: "Mecanografia", Floor: 0},
	"RH":                                 Environment{Key: 5, Name: "RH", Floor: 0},
	"ESCRITORIO_MODELO":                  Environment{Key: 6, Name: "Escritório Modelo", Floor: 0},
	"CAIXA_ELETRONICO":                   Environment{Key: 7, Name: "Caixa Eletrônico", Floor: 0},
	"ENTRADA_SAIDA_ANFITEATRO":           Environment{Key: 8, Name: "Entrada/Saída Anfiteatro", Floor: 0},
	"ALMOXARIFADO":                       Environment{Key: 9, Name: "Almoxarifado", Floor: 0},
	"CORREDOR_ESTETICA_SUL":              Environment{Key: 10, Name: "Corredor Estética (Sul)", Floor: 0},
	"ESCADA_ESTETICA":                    Environment{Key: 11, Name: "Escada Estética", Floor: 0},
	"CORREDOR_ESTETICA_NORTE":            Environment{Key: 12, Name: "Corredor Estética (Norte)", Floor: 0},
	"RAMPA_ENTRADA_FRENTE":               Environment{Key: 13, Name: "Rampa - Entrada da Frente", Floor: 0},
	"RAMPA_ENTRADA_FRENTE_PATAMAR":       Environment{Key: 14, Name: "Rampa - Entrada da Frente (Patamar)", Floor: 0.5},
	"BANHEIRO_SECRETARIA":                Environment{Key: 15, Name: "Banheiro Secretaria", Floor: 0, IsBathroom: true},
	"ENTRADA_SAIDA_LATERAL":              Environment{Key: 16, Name: "Entrada/Saída Lateral", Floor: 0},
	"PATIO_ALTO":                         Environment{Key: 17, Name: "Pátio (Alto)", Floor: 0},
	"PATIO_BAIXO":                        Environment{Key: 18, Name: "Pátio (Baixo)", Floor: 0},
	"ENTRADA_SAIDA_BLOCO_ARQUITETURA":    Environment{Key: 19, Name: "Entrada/Saída Bloco (Arquitetura)", Floor: 0},
	"CORREDOR_ARQUITETURA":               Environment{Key: 19, Name: "Corredor Arquitetura", Floor: 0},
	"ESCADA_ARQUITETURA":                 Environment{Key: 20, Name: "Escada Arquitetura", Floor: 0},
	"CANTINA":                            Environment{Key: 21, Name: "Cantina", Floor: -1},
	"EAD":                                Environment{Key: 22, Name: "EAD", Floor: -1},
	"ENTRADA_SAIDA_BLOCO_PATIO":          Environment{Key: 23, Name: "Entrada/Saída Bloco/Pátio", Floor: -1}, // Sub
	"CORREDOR_FISIO":                     Environment{Key: 24, Name: "Corredor Fisioterapia", Floor: -1},
	"BANHEIRO_FISIO":                     Environment{Key: 25, Name: "Banheiro Fisioterapia", Floor: -1, IsBathroom: true},
	"ENTRADA_SAIDA_BLOCO_ESTACIONAMENTO": Environment{Key: 26, Name: "Entrada/Saída Bloco/Estacionamento", Floor: -1},
	"ESTACIONAMENTO":                     Environment{Key: 27, Name: "Estacionamento", Floor: -1},
	"PORTAO_ESTACIONAMENTO":              Environment{Key: 28, Name: "Portão Estacionamento", Floor: -1},
	"QUIOSQUE":                           Environment{Key: 29, Name: "Quiosque", Floor: -1},
	"LAGO":                               Environment{Key: 30, Name: "Lago", Floor: -1},
	"SALA_PROFESSORES":                   Environment{Key: 31, Name: "Sala dos Professores", Floor: -1},
	"ANFITEATRO":                         Environment{Key: 32, Name: "Anfiteatro", Floor: -1},
	"PORTAO_ANFITEATRO":                  Environment{Key: 33, Name: "Portão Anfiteatro", Floor: -1},
	"ESCADA_BIBLIOTECA_RAMPA":            Environment{Key: 34, Name: "Escada da Biblioteca/Rampa", Floor: 1}, // 1 Andar
	"ESCADA_BIBLIOTECA_PATIO":            Environment{Key: 35, Name: "Escada da Biblioteca/Patio", Floor: 1},
	"BIBLIOTECA":                         Environment{Key: 36, Name: "Biblioteca", Floor: 1},
	"DIRETORIA":                          Environment{Key: 37, Name: "Diretoria", Floor: 1},
	"RAMPA_DIRETORIA":                    Environment{Key: 38, Name: "Rampa - Diretoria", Floor: 1},
	"FIES":                               Environment{Key: 39, Name: "FIES", Floor: 1},
	"LABORATORIOS_SISTEMAS_LESTE":        Environment{Key: 40, Name: "Laboratórios de Sistemas (Leste)", Floor: 1},
	"ATLETICA":                           Environment{Key: 41, Name: "Atlética (Lobão)", Floor: 1},
	"RAMPA_LABORATORIOS":                 Environment{Key: 42, Name: "Rampa dos Laboratórios", Floor: 1},
	"BANHEIRO_LABORATORIOS":              Environment{Key: 43, Name: "Banheiro dos Laboratórios", Floor: 1, IsBathroom: true},
	"LABORATORIOS_SISTEMAS_OESTE":        Environment{Key: 44, Name: "Laboratórios de Sistemas (Oeste)", Floor: 1},
	"TECNICA":                            Environment{Key: 45, Name: "Técnica", Floor: 1},
	"CORREDOR_DIREITO_SUL":               Environment{Key: 46, Name: "Corredor Direito (Sul)", Floor: 1},
	"CORREDOR_DIREITO_NORTE":             Environment{Key: 47, Name: "Corredor Direito (Norte)", Floor: 1},
	"BANHEIRO_DIREITO":                   Environment{Key: 47, Name: "Banheiro Direito", Floor: 1, IsBathroom: true},
	"ESCADA_DIREITO":                     Environment{Key: 48, Name: "Escada Direito", Floor: 1},
	"CEUP":                               Environment{Key: 49, Name: "CEUP", Floor: 1},
	"BANHEIRO_CEUP":                      Environment{Key: 50, Name: "Banheiro CEUP", Floor: 1, IsBathroom: true},
}

func getEnvironmentByKey(key int) *Environment {
	for _, env := range Environments {
		if env.Key == key {
			return &env
		}
	}
	return nil
}

// EnvironmentEdge edge between two environments
type EnvironmentEdge struct {
	from      Environment
	to        Environment
	direction string
	distance  int
}

func createEdge(slice []EnvironmentEdge, from string, to string, direction string,
	reversedDirection string, distance int) []EnvironmentEdge {

	fromEnv := Environments[from]
	if fromEnv.Name == "" {
		log.Printf("%s is not a valid environment", from)
	}

	toEnv := Environments[to]
	if toEnv.Name == "" {
		log.Printf("%s is not a valid environment", to)
	}

	slice = append(slice, EnvironmentEdge{
		from:      fromEnv,
		to:        toEnv,
		direction: direction,
		distance:  distance},
	)

	slice = append(slice, EnvironmentEdge{
		from:      toEnv,
		to:        fromEnv,
		direction: reversedDirection,
		distance:  distance},
	)

	return slice
}

func initEdges() []EnvironmentEdge {
	var edges = []EnvironmentEdge{}

	edges = createEdge(edges, "ENTRADA_SAIDA_FRENTE", "SECRETARIA", "N", "S", 5)
	edges = createEdge(edges, "SECRETARIA", "BANHEIRO_SECRETARIA", "N", "S", 15)
	edges = createEdge(edges, "SECRETARIA", "PROVE", "W", "E", 5)
	edges = createEdge(edges, "PROVE", "RAMPA_ENTRADA_FRENTE", "N", "S", 5)
	edges = createEdge(edges, "PROVE", "GUARDA", "W", "E", 5)
	edges = createEdge(edges, "GUARDA", "COORDENACAO", "W", "E", 2)
	edges = createEdge(edges, "COORDENACAO", "MECANOGRAFIA", "W", "E", 5)
	edges = createEdge(edges, "COORDENACAO", "RH", "W", "E", 24)
	edges = createEdge(edges, "RH", "ESCRITORIO_MODELO", "W", "E", 5)
	edges = createEdge(edges, "ESCRITORIO_MODELO", "CAIXA_ELETRONICO", "W", "E", 4)
	edges = createEdge(edges, "CAIXA_ELETRONICO", "ENTRADA_SAIDA_ANFITEATRO", "W", "E", 10)
	edges = createEdge(edges, "CAIXA_ELETRONICO", "ALMOXARIFADO", "N", "S", 5)
	edges = createEdge(edges, "ALMOXARIFADO", "CORREDOR_ESTETICA_SUL", "N", "S", 15)
	edges = createEdge(edges, "CORREDOR_ESTETICA_SUL", "ESCADA_ESTETICA", "N", "S", 15)
	edges = createEdge(edges, "ESCADA_ESTETICA", "CORREDOR_ESTETICA_NORTE", "N", "S", 15)
	edges = createEdge(edges, "ESCADA_ESTETICA", "ENTRADA_SAIDA_BLOCO_PATIO", "N", "S", 5)
	edges = createEdge(edges, "MECANOGRAFIA", "PATIO_ALTO", "N", "S", 20)
	edges = createEdge(edges, "RAMPA_ENTRADA_FRENTE", "RAMPA_ENTRADA_FRENTE_PATAMAR", "N", "S", 7)    // 0 to 1
	edges = createEdge(edges, "RAMPA_ENTRADA_FRENTE_PATAMAR", "DIRETORIA", "S", "N", 7)               // 0 to 1
	edges = createEdge(edges, "RAMPA_ENTRADA_FRENTE_PATAMAR", "ESCADA_BIBLIOTECA_RAMPA", "N", "S", 1) // 0 to 1/
	edges = createEdge(edges, "ESCADA_BIBLIOTECA_RAMPA", "BIBLIOTECA", "N", "S", 3)                   // 0 to 1
	edges = createEdge(edges, "ENTRADA_SAIDA_LATERAL", "BANHEIRO_SECRETARIA", "S", "N", 5)
	edges = createEdge(edges, "ENTRADA_SAIDA_LATERAL", "CORREDOR_ARQUITETURA", "N", "S", 15)
	edges = createEdge(edges, "BANHEIRO_SECRETARIA", "CORREDOR_ARQUITETURA", "N", "S", 15)
	edges = createEdge(edges, "PATIO_ALTO", "ENTRADA_SAIDA_BLOCO_ARQUITETURA", "E", "W", 10)
	edges = createEdge(edges, "PATIO_ALTO", "ESCADA_BIBLIOTECA_PATIO", "E", "W", 10)
	edges = createEdge(edges, "PATIO_ALTO", "PATIO_BAIXO", "N", "S", 15)
	edges = createEdge(edges, "ENTRADA_SAIDA_BLOCO_ARQUITETURA", "CORREDOR_ARQUITETURA", "E", "W", 10)
	edges = createEdge(edges, "ENTRADA_SAIDA_BLOCO_ARQUITETURA", "BANHEIRO_SECRETARIA", "E", "W", 10)
	edges = createEdge(edges, "CORREDOR_ARQUITETURA", "ESCADA_ARQUITETURA", "N", "S", 15)
	edges = createEdge(edges, "ESCADA_ARQUITETURA", "CANTINA", "N", "S", 8)
	edges = createEdge(edges, "CANTINA", "EAD", "N", "S", 8)
	edges = createEdge(edges, "CANTINA", "PATIO_BAIXO", "W", "E", 8)
	edges = createEdge(edges, "ENTRADA_SAIDA_BLOCO_PATIO", "PATIO_BAIXO", "E", "W", 20)    // 0 to -1
	edges = createEdge(edges, "ENTRADA_SAIDA_BLOCO_PATIO", "ESCADA_ESTETICA", "W", "E", 5) // 0 to -1
	edges = createEdge(edges, "ENTRADA_SAIDA_BLOCO_PATIO", "CORREDOR_FISIO", "N", "S", 15)
	edges = createEdge(edges, "ENTRADA_SAIDA_BLOCO_PATIO", "BANHEIRO_FISIO", "W", "E", 15)
	edges = createEdge(edges, "ENTRADA_SAIDA_BLOCO_ESTACIONAMENTO", "BANHEIRO_FISIO", "E", "W", 5)
	edges = createEdge(edges, "ENTRADA_SAIDA_BLOCO_ESTACIONAMENTO", "BANHEIRO_FISIO", "E", "W", 5)
	edges = createEdge(edges, "ENTRADA_SAIDA_BLOCO_ESTACIONAMENTO", "ESTACIONAMENTO", "N", "S", 25)
	edges = createEdge(edges, "ENTRADA_SAIDA_BLOCO_ESTACIONAMENTO", "SALA_PROFESSORES", "W", "E", 7)
	edges = createEdge(edges, "ESTACIONAMENTO", "SALA_PROFESSORES", "S", "N", 20)
	edges = createEdge(edges, "ESTACIONAMENTO", "LAGO", "SW", "NE", 25)
	edges = createEdge(edges, "ESTACIONAMENTO", "QUIOSQUE", "W", "E", 25)
	edges = createEdge(edges, "ESTACIONAMENTO", "PORTAO_ESTACIONAMENTO", "N", "S", 15)
	edges = createEdge(edges, "QUIOSQUE", "LAGO", "N", "S", 5)
	edges = createEdge(edges, "ANFITEATRO", "SALA_PROFESSORES", "N", "S", 40)
	edges = createEdge(edges, "ANFITEATRO", "PORTAO_ANFITEATRO", "S", "N", 10)
	edges = createEdge(edges, "PORTAO_ANFITEATRO", "ENTRADA_SAIDA_ANFITEATRO", "E", "W", 10)
	edges = createEdge(edges, "DIRETORIA", "FIES", "E", "W", 10)
	edges = createEdge(edges, "DIRETORIA", "LABORATORIOS_SISTEMAS_LESTE", "W", "E", 20)
	edges = createEdge(edges, "FIES", "BIBLIOTECA", "N", "S", 30)
	edges = createEdge(edges, "BIBLIOTECA", "ESCADA_BIBLIOTECA_PATIO", "W", "E", 10)
	edges = createEdge(edges, "LABORATORIOS_SISTEMAS_LESTE", "ATLETICA", "W", "E", 20)
	edges = createEdge(edges, "ATLETICA", "CORREDOR_DIREITO_SUL", "N", "S", 20)
	edges = createEdge(edges, "ATLETICA", "BANHEIRO_LABORATORIOS", "W", "E", 25)
	edges = createEdge(edges, "CORREDOR_DIREITO_SUL", "BANHEIRO_DIREITO", "N", "S", 5)
	edges = createEdge(edges, "BANHEIRO_DIREITO", "ESCADA_DIREITO", "N", "S", 10)
	edges = createEdge(edges, "ESCADA_DIREITO", "CORREDOR_DIREITO_NORTE", "N", "S", 20)
	edges = createEdge(edges, "CORREDOR_DIREITO_NORTE", "CEUP", "N", "S", 20)
	edges = createEdge(edges, "CEUP", "BANHEIRO_CEUP", "E", "W", 5)
	edges = createEdge(edges, "BANHEIRO_LABORATORIOS", "LABORATORIOS_SISTEMAS_OESTE", "W", "E", 5)
	edges = createEdge(edges, "LABORATORIOS_SISTEMAS_OESTE", "TECNICA", "W", "E", 10)
	edges = createEdge(edges, "TECNICA", "TECNICA", "W", "E", 10)
	edges = createEdge(edges, "ESCADA_DIREITO", "ESCADA_ESTETICA", "W", "E", 7) // 1 to 0

	return edges
}

// Edges list of edges between environments
var environmentEdges = initEdges()

func buildEnvironmentsGraph() ds.Graph {
	vertices := make([]ds.Vertex, len(Environments))
	for _, env := range Environments {
		v := ds.NewVertex(env.Key, env)
		vertices[env.Key] = v
	}

	edges := ds.NewListContainer(len(vertices))
	for _, edge := range environmentEdges {
		edges.InsertEdge(edge.from.Key, edge.to.Key, edge.distance)
		edges.InsertEdge(edge.to.Key, edge.from.Key, edge.distance)
	}

	g := ds.NewGraph(vertices, edges)
	return g
}

var environmentsGraph = buildEnvironmentsGraph()

// MinimumPath returns the path and distance between two environments
func MinimumPath(from Environment, to Environment) ([]Environment, int) {
	path, distance := environmentsGraph.MinimumPathTo(from.Key, to.Key)

	envPath := make([]Environment, len(path))
	for i, key := range path {
		envPath[i] = *getEnvironmentByKey(key)
	}

	return envPath, distance
}

// MinimumPathToBathroom returns the path and distance between two environments
func MinimumPathToBathroom(from Environment) ([]Environment, int) {
	path, distance := environmentsGraph.MinimumPath(from.Key)

	var nearestBathroom *Environment
	var nearestDistance int = math.MaxInt32
	for i, key := range path {
		if key > -1 {
			env := getEnvironmentByKey(key)
			if env.IsBathroom {
				if distance[i] < nearestDistance {
					nearestBathroom = env
					nearestDistance = distance[i]
				}
			}
		}
	}

	return MinimumPath(from, *nearestBathroom)
}
