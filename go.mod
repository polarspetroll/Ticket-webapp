module github.com/polarspetroll/Ticket-webapp

go 1.15


require (
	admin v0.0.0
	register v0.0.0
	list v0.0.0
)


replace (
	admin => ./admin
	register => ./register
	list => ./list
)