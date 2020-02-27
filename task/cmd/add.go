package cmd

var AddCmd = &cobra.Command{
	Use:"add",
	Short:"Add a task to your list",
	Run: func(cmd *cobra.Command,args []string){
		fmt.Println("add called")
	}
}
