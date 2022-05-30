package handlers

import( 
	"encoding/json"
	"log"
	"net/http"
 	"srtconv"

 	"github.com/Amad3eu/api-pstg-go/models"
)

func Update(w http.ResponseWriter, r *http.request){
	id, err := strconv.Atoi(chi.URLParam(r,"id"))
	if err != nil{
	log.println("Error decoding JSON: %v", err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	return
}

 var todo models.todo

 err = json.NewDecoder(r.body).Decode(&todo)
 if err != nil{
 log.println("Error decoding JSON: %v", err)
 http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
 return
}

rows, err := models.Update(int64(id,todo))
	if err != nil{
	log.println("Error ao atualizar registro: %v", err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	return
} 

if rows > 1{
	log.printf("Error: foram atualizados %d registros", rows)
}
resp := map[string]any{
	"Error":false
	"Message":"dados atualizados com sucesso!",
}

w.header().Add("Content-Type", "application/json")
json.NewEncoder(w).Encode(resp)
}
