package handlers


import ("net/http"
 "strconv"
"github.com/Amad3eu/api-pstg-go/models"
)

func Get(w http.ResponseWritter, r *http.Request){
id, err := strconv.Atoi(chi.URLParam(r,"id"))
if err != nil{
log.println("Erro ao fazer parse do id: %v", err)
http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
return
}

todo, err := models.Get(int64(id))
if err != nil{
log.println("Error ao atualizar registro: %v", err)
http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
return
} 

w.header().Add("Content-Type", "application/json")
json.NewEncoder(w).Encode(todo)
}
