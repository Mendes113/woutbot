package bot

import (
	"log"
	"strconv"
	"strings"


)

// Exercise representa um exercício com sets, repetições e carga
type Exercise struct {
    Name   string
    Sets   []int
    Reps   []int
    Weights []int
}
// NewExercise cria um novo exercício com o nome fornecido
func NewExercise(name string) *Exercise {
    return &Exercise{
        Name: name,
    }
}

// AddSet adiciona um conjunto (set) ao exercício
func (e *Exercise) AddSet(set int) {
    e.Sets = append(e.Sets, set)
}

// AddRep adiciona uma repetição (rep) ao exercício
func (e *Exercise) AddRep(rep int) {
    e.Reps = append(e.Reps, rep)
}

// AddWeight adiciona uma carga (weight) ao exercício
func (e *Exercise) AddWeight(weight int) {
    e.Weights = append(e.Weights, weight)
}

// Workload calcula a carga total do exercício
func (e *Exercise) Workload(weight int, reps int, set int) int {
    var workload int
    for i := 0; i < len(e.Sets); i++ {
        workload += e.Sets[i] * e.Reps[i] * e.Weights[i]
    }
    return workload
}

// GetExerciseDetails retorna os detalhes do exercício formatados
func (e *Exercise) GetExerciseDetails() string {
    var details []string
    for i := 0; i < len(e.Sets); i++ {
        set := strconv.Itoa(e.Sets[i])
        rep := strconv.Itoa(e.Reps[i])
        weight := strconv.Itoa(e.Weights[i])
        worload := strconv.Itoa(e.Workload(e.Weights[i], e.Reps[i], e.Sets[i]))
        details = append(details, set+" x "+rep+" x "+weight+"kg" + "\n" + "Carga total: " + worload + "kg")
    }
    return strings.Join(details, "\n")
}

// Workout representa um treino composto por uma lista de exercícios
type Workout struct {
    Exercises []*Exercise
    TotalWorkload int
}

// NewWorkout cria um novo treino vazio
func NewWorkout() *Workout {
    return &Workout{
        Exercises: make([]*Exercise, 0),
    }
}

// AddExercise adiciona um exercício ao treino
func (w *Workout) AddExercise(exercise *Exercise) {
    w.Exercises = append(w.Exercises, exercise)
}




// getTreinoDetails retorna os detalhes do treino formatados
func getTreinoDetails(treino *Workout) string {
    var details []string
    for _, exercicio := range treino.Exercises {
        details = append(details, exercicio.Name+":\n"+exercicio.GetExerciseDetails())
    }

    return strings.Join(details, "\n\n")
}



// Modificando a função MakeTrain para calcular o trabalho total do treino
func MakeTrain(chatID int64, message string) string {
	log.Print("Montar treino")
	lines := strings.Split(message, "\n")

	treino := NewWorkout()
	totalWorkload := 0 // Variável para armazenar o trabalho total do treino

	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) < 4 {
			continue
		}

		name := parts[0]
		sets, _ := strconv.Atoi(parts[1])
		reps, _ := strconv.Atoi(parts[2])
		weight, _ := strconv.Atoi(parts[3])

		exercicio := NewExercise(name)
		for i := 0; i < sets; i++ {
			exercicio.AddSet(sets)
			exercicio.AddRep(reps)
			exercicio.AddWeight(weight)
			totalWorkload += exercicio.Workload(weight, reps, sets) // Adiciona o trabalho deste exercício ao total
		}

		treino.AddExercise(exercicio)
	}


	return getTreinoDetails(treino)
}