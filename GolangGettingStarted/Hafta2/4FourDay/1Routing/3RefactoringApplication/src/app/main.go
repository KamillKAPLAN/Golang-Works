package main

func main() {
	r := registerRoutes()
	// bu projeyi çalıştırmak için ; aşağıdaki yol izlenir
	// kamilkaplan@kamilkaplan:~/go/src/github.com/GO/Hafta2/4FourDay/Routing/3RefactoringApplication/src/app$ ./app
	r.Run(":3000")
}
