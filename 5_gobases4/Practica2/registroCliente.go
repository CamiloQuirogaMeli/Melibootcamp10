package main

//El mismo estudio del ejercicio anterior, también solicita una funcionalidad para poder registrar datos de nuevos clientes. Los datos requeridos para registrar
//a un cliente son:
//- Legajo
//- Nombre y Apellido
//- DNI
//- Número de teléfono
//- Domicilio
//El número de legajo debe ser asignado o generado por separado y en forma previa a la carga de los restantes gastos.
//Para cumplir con los objetivos de aprendizaje de este módulo, aún no vamos a escribir los datos en un archivo. Sin perjuicio de lo cual, supongamos que quieres
//constatar si el cliente se encuentra ya registrado y para ello precisas leer los datos de un archivo .txt. En algún lugar de tu código, simula que intentas leer
//un archivo llamado “customers.txt” (como en el ejercicio anterior). Claramente, este archivo que intentas leer no existirá, por lo que la función a utilizar,
//para intentar la lectura, devolverá un error que deberás manipular adecuadamente como hemos visto en el contenido hasta aquí. El error deberá generar un panic,
//luego lanzar por consola el mensaje: “error: el archivo indicado no fue encontrado o está dañado”, y continuar con la ejecución del programa normalmente.