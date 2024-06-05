function imc(peso, altura) {
    const total = peso / (altura * altura)

    if (total <= 18.5) {
        return "Seu IMC é de : " + total + " Nível 0 -> - MAGREZA"
    } else if (total > 18.5 && total <= 24.9) {
        return "Seu IMC é de : " + total + " Nível 0 -> - NORMAL"
    } else if (total > 24.9 && total <= 29.9) {
        return "Seu IMC é de : " + total + " Nível I -> SOBREPESO"
    } else if (total > 29.9 && total <= 39.9) {
        return "Seu IMC é de : " + total + " Nível II -> OBESIDADE"
    } else (total > 40) 
    return "Seu IMC é de : " + total + " Nível III -> OBESIDADE GRAVE"
}

console.log(imc(49,1.54))