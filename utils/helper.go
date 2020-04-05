package utils

import "log"

// CheckErr checar o erro
func CheckErr(err error) (mensagem string) {
	mensagem = ""

	if err != nil {
		log.Printf("CheckErr(): %q\n", err)
		mensagem = err.Error()
	}

	return mensagem
}
