package bot

import (

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
func (e *Exercise) Workload() int {
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
        details = append(details, set+" x "+rep+" x "+weight+"kg")
    }
    return strings.Join(details, "\n")
}

// Workout representa um treino composto por uma lista de exercícios
type Workout struct {
    Exercises []*Exercise
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

// MontarTreino monta o treino com base na mensagem do usuário
func MontarTreino(chatID int64, mensagem string) string{
    // Criar um novo treino
    treino := NewWorkout()

    // Dividir a mensagem em linhas
    linhas := strings.Split(mensagem, "\n")

    // Processar cada linha para extrair informações de cada exercício
    for _, linha := range linhas {
        partes := strings.Fields(linha)
        if len(partes) != 4 {
            // SendMessage(chatID, "Formato inválido. Use o formato 'sets reps peso' separados por espaços.")
            return ""
        }

        // Extrair informações
        set, err1 := strconv.Atoi(partes[0])
        rep, err2 := strconv.Atoi(partes[1])
        peso, err3 := strconv.Atoi(strings.TrimSuffix(partes[3], "kg"))

        // Verificar erros de conversão
        if err1 != nil || err2 != nil || err3 != nil {
            // SendMessage(chatID, "Erro ao converter valores. Certifique-se de usar números válidos.")
            return ""
        }

        // Criar e adicionar o exercício ao treino
        exercicio := NewExercise(partes[2])
        exercicio.AddSet(set)
        exercicio.AddRep(rep)
        exercicio.AddWeight(peso)
        treino.AddExercise(exercicio)
    }

    // Retornar os detalhes do treino formatados
    return getTreinoDetails(treino)
}

// getTreinoDetails retorna os detalhes do treino formatados
func getTreinoDetails(treino *Workout) string {
    var details []string
    for _, exercicio := range treino.Exercises {
        details = append(details, exercicio.Name+":\n"+exercicio.GetExerciseDetails())
    }
    return strings.Join(details, "\n\n")
}