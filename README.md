# goboxes
golang data structure include deque, queue, stack, priority queue

## Installation
`$go get github/JackyXiong/goboxes`

## Usage
```
import "github/JackyXiong/goboxes"

var deque *goboxes.Deque = goboxes.NewDeque()
```
#### Deque
* NewDeque
* Append
* AppendLeft
* Pop
* PopLeft
* First
* Last
* Extend
* Count
* Remove

#### Queue: FIFO queue
* NewQueue
* Put
* Get

#### Stack: LIFO queue
* NewQueue
* Put
* Get

#### PriorityQueue
* NewPriorityQueue
* Put
* Get


### TODO 
* Test priority queue



