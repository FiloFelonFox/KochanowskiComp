generate:
	antlr4 -Dlanguage=Go -o parser/ kochanowski.g4

runMain:
	go run . main.kochanowski kochanowski.ll
	lli kochanowski.ll