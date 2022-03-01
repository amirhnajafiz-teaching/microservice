package cmd

func Execute() {
	router := New()

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
