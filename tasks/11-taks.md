
Take example 71-http-fileops-chan

1. create GetUsers handler that should gibe all the users in users.txt file 
    - iterate the file line by line
    - each line unmarshal it to the User struct 
    - create a slice of []User
    - write it back using r.Write 
    
2. The file name to be given thru optional /command line arguments. 
3. If no filename argument is given then it should take users.txt as a default file to create and also to be read from