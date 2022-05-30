package handlers

import(
	"http"
	"models"
	"log"
	"json"

	"github.com/Amad3eu/api-pstg-go/models"
)

func List(w http.ResponseWritter, r *http.Request){
 todos, err := models.GetAll()
 if err != nil{
	 log.Printf("Erro ao obter registros: %v", err)
 }

 w.Header().Add("Content-Type", "application/json")
 json.NewEncoder(w).Encode(todos)
}
