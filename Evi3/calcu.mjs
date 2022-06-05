import soap from 'soap'
import express from 'express'
import fs from 'fs'

var soapService = {
    calculatorService: {
        calculatorServiceSoapPort: {
            Suma: (args) => {
                return {
                    result: args.a + args.b
                }
            },
            Mult: (args) => {
                return {
                    result: args.a * args.b
                }
            },
            Resta: (args) => {
                return {
                    result: args.a - args.b
                }
            },
            Divi: (args) => {
                if (args.b == 0) {
                    return {
                        result: 0
                    }
                }
                return {
                    result: args.a / args.b
                }
            }
        }
    }
}

var wsld = fs.readFileSync("ns.wsdl", "utf8");

var app = express();

app.listen(5050, () => { //Inicializar el servidor
    soap.listen(app, '/calculadora', soapService, wsld, function() {
        console.log("Iniciando Servidor");
    })
})