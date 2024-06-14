package main

import "fmt"

func main() {
	var usuario string
	var contraseña string
	var op1 int
	var name string
	var nroLaborerEliminado int
	var nroLaborerNuevo int
	var laborerNuevo []string
	for {
		fmt.Println("----------------------------")
		fmt.Println("Bienvenido, ingrese 1 si quiere iniciar sesión, si quiere terminar el programa, ingrese cualquier otro caracter")
		fmt.Scanf("%d\n", &op1)
		fmt.Println("----------------------------")
		if op1 == 1 {
			for {
				fmt.Println("----------------------------")
				fmt.Println("Bienvenido, por favor ingrese su usuario y contraseña.")
				fmt.Println("----------------------------")
				fmt.Println("Usuario:")
				fmt.Scanf("%s\n", &usuario)
				fmt.Println("----------------------------")
				fmt.Println("Contraseña:")
				fmt.Scanf("%s\n", &contraseña)
				fmt.Println("----------------------------")
				if usuario == "admin" && contraseña == "root" {
					fmt.Println("Bienvenido Administrador...")
					fmt.Println("----------------------------")
					for {
						fmt.Println("Ingrese 1 si quiere eliminar un Laborer, ingrese 2 si desea crear uno, cualquier otro caracter para cerrar sesión")
						fmt.Println("----------------------------")
						fmt.Scanf("%d\n", &op1)
						if op1 == 1 {
							fmt.Println("Ingrese el nombre del laborer que desea eliminar")
							fmt.Scanf("%s\n", &name)
							fmt.Println("----------------------------")
							fmt.Println("laborer eliminado")
							fmt.Println("----------------------------")
							nroLaborerEliminado = nroLaborerEliminado + 1
							fmt.Println("----------------------------")
							fmt.Println("Numero de Laborers eliminados: ", nroLaborerEliminado)
							fmt.Println("----------------------------")
						} else if op1 == 2 {
							fmt.Println("----------------------------")
							fmt.Println("Ingrese el nombre del nuevo laborer")
							fmt.Scanf("%s\n", &name)
							fmt.Println("----------------------------")
							laborerNuevo = append(laborerNuevo, "Laborer creado\n")
							nroLaborerNuevo = nroLaborerNuevo + 1
							fmt.Println("Numero de Usuarios creados: ", nroLaborerNuevo)
							fmt.Println("----------------------------")
						} else {
							fmt.Println("Cerrando sesión")
							fmt.Println("----------------------------")
							break
						}
					}

				} else if usuario == "seeker223" && contraseña == "seekr" {
					fmt.Println("Bienvenido Supervisor")
					fmt.Println("----------------------------")
					for {

						fmt.Println("Ingrese 1 si quiere crear registro de administrador, ingrese cualquier otro caracter para cerrar sesión")
						fmt.Scanf("%d\n", &op1)
						if op1 == 1 {
							fmt.Println("----------------------------")
							fmt.Println("laborer creados: ", nroLaborerNuevo)
							fmt.Println("----------------------------")
							fmt.Println("laborer eliminados: ", nroLaborerEliminado)
							fmt.Println("----------------------------")
						} else {
							break
						}

					}

				} else {
					fmt.Println("Usuario o contraseña incorrecto")
				}
			}
		} else {
			fmt.Println("Programa terminado")
			fmt.Println("----------------------------")
			break
		}

	}

}
