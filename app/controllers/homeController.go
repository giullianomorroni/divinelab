package controllers

import (
	cb "divinelab/app/controllers/base"
	"github.com/robfig/revel"
	"net/smtp"
)

type (
	Home struct {
		cb.BaseController
	}
)

func init() {}

func (this *Home) Index() revel.Result {
	return this.Render();
}

func (this *Home) Valores() revel.Result {
	return this.Render();
}

func (this *Home) Produtos() revel.Result {
	return this.Render();
}

func (this *Home) Contato(nome, email, texto string) revel.Result {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n";
	revel.TRACE.Printf("[EmailService.go] Disparo email EnviarEmailContato")

    texto = texto + " MENSAGEM ENVIADA POR "+ nome +", email para contato: "+ email;
    msg := []byte(mime + texto)

	auth := smtp.PlainAuth("Divine Slim", "suporte@divineslim.com.br", "ds20142015", "smtp.gmail.com")
	to := []string{"comercial@divinelab.com.br"}
	err := smtp.SendMail("smtp.gmail.com:587", auth, "contato@divineslim.com.br", to, msg)
	if err != nil {
		revel.TRACE.Printf("Erro no disparo de email %s", err)
	}
    this.Flash.Success("Recebemos sua mensagem com sucesso!")
    return this.Redirect((*Home).Index)
}