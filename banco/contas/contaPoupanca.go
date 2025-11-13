package contas

type ContaPoupanca struct {
	Titular                        string
	Agencia, NumeroConta, Operacao int
	saldo                          float64
}

func (c *ContaPoupanca) Sacar(valor float64) string {
	podeSacar := valor > 0 && valor <= c.saldo
	if !podeSacar {
		return "saldo insuficiente"
	}
	c.saldo -= valor
	return "Saque realizado com sucesso"
}

func (c *ContaPoupanca) Depositar(valor float64) (string, float64) {
	if valor <= 0 {
		return "Valor de depósito inválido", c.saldo
	}
	c.saldo += valor
	return "Depósito realizado com sucesso", c.saldo
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
