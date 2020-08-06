package middlewares

/*
	// GET

	request has a jwt in a header or cookie?
	- no -> redirect to /login
	- yes ->
			auth := auth.New(name)
			token := auth.GenerateToken()

			check (GET) against Redis if KEY/VALUE (NAME/TOKEN)
			exists and if it is equal of what we have

			- yes ->
				route / login or /register?
					- yes -> redirect to / (already registered)
			- no -> redirect to login and clear this token

	// POST

	request has a jwt in a header or cookie?
	- no -> redirect to /login
	- yes ->
			auth := auth.New(name)
			token := auth.GenerateToken()

			check (GET) against Redis if KEY/VALUE (NAME/TOKEN)
			exists and if it is equal of what we have

			- yes ->
				route is / login or /register?
					- yes -> redirect to / (already registered)
					- no -> keep flow
			- no -> redirect to login and clear this token

*/
