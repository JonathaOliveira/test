def CalculaIMC(peso, altura):
    total = peso / (altura * altura)

    if total < 18.5:
        return "Seu IMC é de : ", total, " Magro"
    elif 18.5 > total <= 25.00:
        return "Seu IMC é de : ", total, " Normal"
    elif 25.00 > total <= 30.00:
        return "Seu IMC é de : ", total, " Sobrepeso"
    elif 30.00 > total <= 35.00:
        return "Seu IMC é de : ", total, " Obesidade nível 1"
    else:
        return  "Seu IMC é de : ", total, " Obesidade nível 2"

Resultado = CalculaIMC(98, 1.79)

print(Resultado)