package entrypoints

import (
	"encoding/json"
	"github.com/davidveg/goapi/modules/internal/entrypoints/dto"
	"github.com/davidveg/goapi/modules/internal/entrypoints/queues"
	"net/http"
)

func SendMessages(w http.ResponseWriter, r *http.Request) {
	var message dto.SQSMessageRequest
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = queues.SendSQSMessages(&message)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode("mensagem enviada com sucesso !")
}
