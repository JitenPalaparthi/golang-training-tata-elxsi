	•	Generator pattern reads data.txt
	•	No bufio / no readers (we use os.File.Read + manual newline detection)
	•	Each line is emitted as soon as newline is detected
	•	Each line carries a line number
	•	Same line is broadcast to multiple channels
	•	Separate goroutines compute:
	1.	words per line
	2.	chars per line
	3.	global word counts (fan-out workers, scalable)