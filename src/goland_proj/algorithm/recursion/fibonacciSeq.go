package recursion



func FibonacciSequence(n int) (int){

	if n < 2{
		return n
	}else{
		return FibonacciSequence(n - 1) + FibonacciSequence( n - 2)
	}
}





