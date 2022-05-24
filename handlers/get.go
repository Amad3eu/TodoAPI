package handlers


import "net/http
func Get(w http.ResponseWritter, r *http.Request){
id, err := strconv.Atoi(chi.URLParam(r,"id"))
if err != nil{
log.println("Erro ao fazer parse do id: %v", err)
http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
return
}


err = json.NewDecoder(r.body).Decode(&todo)
if err != nil{
log.println("Error decoding JSON: %v", err)
http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
return
}

todo, err := models.Get(int64(id))
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
json.NewEncoder(w).Encode(todo)
}
