package middlewares

/*
	// ROUTES

	request has a jwt in a header or cookie?
	- no -> redirect to /login
	- yes ->
			auth := auth.New(name)
			token := auth.GenerateToken()

			check (GET) against Redis if KEY/VALUE (NAME/TOKEN)
			exists and if it is equal of what we have

			- yes ->
				is GET? AND	route / login or /register?
					- yes -> redirect to / (already registered)
			- no -> IS GET?
					- yes -> redirect to login and clear this token
					- no ->
			- yes ->
				route is / login or /register?
					- yes -> redirect to / (already registered)
					- no -> keep flow if success
						yes -> INSERT into Redis (NAME/TOKEN)
						no -> failed register
			- no -> redirect to login and clear this token

*/
