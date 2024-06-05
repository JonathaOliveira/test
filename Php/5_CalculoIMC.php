<?php

function imc($peso, $altura){
    $total = $peso / ($altura * $altura);
    
    if($total <= 1.85) {
        return "Seu IMC é: " . $total . " Você está abaixo do peso - MAGRO";
    } else if ($total > 1.85 && $total <= 25) {
        return "Seu IMC é: " . $total . " Você está com o peso ideal - NORMAL";
    } else if ($total > 25 && $total <= 30) {
        return "Seu IMC é: " . $total . " Você está com o peso acima do ideal - SOBREPESO";
    }  else if ($total > 30 && $total <= 35) {
        return "Seu IMC é: " . $total . " Você está com o peso muito acima do ideal - OBESIDADE NÍVEL 1";
    } else if ($total > 35) {
        return "Seu IMC é: " . $total . " Você está com o peso critico - OBESIDADE NÍVEL 2";
    }
}

echo(imc(98, 1.79));
    
    

?>