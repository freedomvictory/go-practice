package breadthFirstSearch


// reference book <Algorithm diagram> passage 87

// FIFO is a FIFO queue
type FIFO struct {
	queue []interface{}
}

// New creates new FIFO and returns it
func NewFifo() *FIFO {
	return &FIFO{
		queue: make([]interface{}, 0),
	}
}

// Push pushed node to the back of the queue
func (f *FIFO) Push(node interface{}) {
	f.queue = append(f.queue, node)
}

// Front takes a value from the front of the queue and returns it
func (f *FIFO) Front() interface{} {
	if len(f.queue) == 0 {
		return nil
	}

	node := f.queue[0]
	f.queue[0] = nil
	f.queue = f.queue[1:]

	return node
}

type Figure struct{
	Items map[string][]string
}

func NewFigure() Figure{
	return Figure{
		Items: make(map[string][]string, 32),
	}

}


func IsWeWant(input string) bool {

	var final = len(input) -1
	if input[final] == 'm'{
		return true
	}
	return false
}

func Search(input Figure) interface{}{

	queue := NewFifo()
	for _, val := range input.Items["you"]{
		queue.Push(val)
	}
	searched := []string{}

	for{

		front := queue.Front()
		if front == nil{
			break
		}
		front_s := front.(string)
		for _, val := range searched{
			if front_s == val {
				continue
			}
		}
		if IsWeWant(front_s) == true {
			return front
		}else{
			for _, val := range input.Items[front_s]{
				queue.Push(val)
			}
			searched = append(searched, front_s)
		}


	}






	return nil
}
