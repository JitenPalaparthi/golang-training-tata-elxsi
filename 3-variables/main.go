package main

const PI float32 = 3.14

const (
	MAX = 999
	MIN = 60
)

//group of constants
//var hours = 24 // Varaibles are evaluated at rumtime

func main() {
	const (
		HOUR = 60 * MIN    // This is done by the compiler , not by the runtime
		DAY  = 60*HOUR + 0 // The whole calculation should be known to the compiler
	)

}

/*
 Code/Text segment --> Your binary is loaded in this segment, the size if based on your binary size
	Some OS/Programmign Langauges store constants in Code/Text Segment
 Data Segment --> global,
		package level (static) variables are created but not deallocated, they are stored in data segment
		string literals
		constants
Stack
	Variables those are created as stack frames, local variables etc..
	Generally variables with fixed size are created in stack
	Stack pointer would push and pop in order to allocate and deallocate these data from the momory based on scope
	Stack size of 2mb (varies based on OS but it is constant)

Heap
	Dynamic memory , generally all objects are created in Heap, it can be grown or shrunk based on need
	Either Garbage collector or thru manual memory management , we can deallocate the memory

Note: Wheather the variable is created in stack or heap is purely decided by the compier process(escape analysis)
It follows some algorthm in order to allocate memory either on stack or heap
As a developer, we can expect it but sometimes, the compiler suprises us
*/
