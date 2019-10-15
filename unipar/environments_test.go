package unipar

import "testing"

func TestMinimumPath(t *testing.T) {
	from := Environments["ENTRADA_SAIDA_FRENTE"]
	to := Environments["TECNICA"]

	path, distance := MinimumPath(from, to)

	if distance != 0 {
		t.Errorf("wrong distance (%d)", distance)
	}

	if len(path) == 0 {
		t.Errorf("wrong path (%v)", path)
	}
}
