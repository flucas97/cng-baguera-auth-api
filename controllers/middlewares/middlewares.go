package middlewares

/*
	// paths != /login
	request has a jwt in a header or cookie?
	- no -> redirect to login
	- yes -> auth := auth.New(name)
			 token := auth.GenerateToken()
			 check (GET) against Redis if KEY/VALUE (NAME/TOKEN) exists and if it is equal of what we have
				- yes -> redirect to cannabis service
				- no -> redirect to login and clear this token

	// path GET login
	request has a jwt in a header or cookie?
	- no -> ok
	- yes -> redirect user to / (already register)

	- yes -> redirect user to / (already logged in)

	// path GET signup
	request has a jwt in a header or cookie?
	- no -> ok
	- yes -> redirect user to / (already registered)

*/
