package tendencia

import "errors"

type Tendencia struct {
	Valores []float32
}

func (t *Tendencia) Promedio() (float32, error) {
	if len(t.Valores) == 0 {
		return -1, errors.New("El Array no puede estar vacio")
	}

	var acum float32 = 0
	for _, v := range t.Valores {
		acum += v
	}
	return acum / float32(len(t.Valores)), nil
}

func (t *Tendencia) Moda() (int, error) {
	if len(t.Valores) == 0 {
		return -1, errors.New("El Array no puede estar vacio")
	}

	i := 0
	for j := i + 1; i < len(t.Valores)-1 && j < len(t.Valores); i, j = i+1, j+1 {
		if t.Valores[i] > t.Valores[j] {
			aux := t.Valores[j]
			t.Valores[j] = t.Valores[i]
			t.Valores[i] = aux
		}

	}

	return 0, nil
}

func (t *Tendencia) Mediana() (float32, error) {
	if len(t.Valores) == 0 {
		return -1, errors.New("El Array no puede estar vacio")
	}

	if len(t.Valores)%2 != 0 {
		return t.Valores[(len(t.Valores)-1)/2], nil
	}

	der := len(t.Valores) / 2
	izq := der - 1

	return (t.Valores[der] + t.Valores[izq]) / 2, nil
}
