package main

import "fmt"

func main() {
	q := NewDeque()
	q.EnqueueBack(4)

	q.EnqueueBack(50)

	fmt.Println(q.PeekFront())

	fmt.Println(q.PeekBack())

	q.EnqueueFront(36)

	q.EnqueueFront(96)

	fmt.Println(q.PeekFront())

	fmt.Println(q.PeekBack())

	data, ok := q.DequeueBack()
	if ok {
		fmt.Printf("DequeueBack成功: 取り出した値 = %d\n", data)
	} else {
		fmt.Println("DequeueBack失敗: キューは空でした。")
	}

	data2, ok := q.DequeueBack()
	if ok {
		fmt.Printf("DequeueBack成功: 取り出した値 = %d\n", data2)
	} else {
		fmt.Println("DequeueBack失敗: キューは空でした。")
	}

	fmt.Println(q.PeekFront())

	fmt.Println(q.PeekBack())

	q.EnqueueFront(20)

	data3, ok := q.DequeueBack()
	if ok {
		fmt.Printf("DequeueBack成功: 取り出した値 = %d\n", data3)
	} else {
		fmt.Println("DequeueBack失敗: キューは空でした。")
	}
}

type Node struct {
	data int32
	next *Node
	prev *Node
}

type Deque struct {
	head *Node
	tail *Node
}

func NewDeque() *Deque {
	return &Deque{}
}

func (d *Deque) PeekFront() int32 {
	return d.head.data
}

func (d *Deque) PeekBack() int32 {
	return d.tail.data
}

// 両端キューの前に追加
func (d *Deque) EnqueueFront(data int32) {
	// 前 = head = prevがnil
	// 連結リストのhead、つまり一番手前側のこと

	// まだキューに要素がない場合
	if d.head == nil {
		// headとtail両方に新しいNodeを入れる
		d.head = &Node{data: data}
		d.tail = d.head
	} else {
		// headの１つ前、存在してないprevに新しいNodeを追加
		d.head.prev = &Node{data: data}
		// 新しいNodeから見て、nextに旧headを追加
		d.head.prev.next = d.head
		// headが指す位置を、新しいNodeにずらす
		d.head = d.head.prev
	}
}

// 両端キューの後ろに追加
func (d *Deque) EnqueueBack(data int32) {
	// 後ろ = tail = nextがnil
	// 連結リストのtail、つまり一番奥側のこと

	if d.head == nil {
		d.head = &Node{data: data}
		d.tail = d.head
	} else {
		// headの1つ次、存在してないnextに新しいNodeを追加
		d.tail.next = &Node{data: data}
		// 新しいNodeから見て、1つ前に旧tailを追加
		d.tail.next.prev = d.tail
		// tailが指す位置を新しいNodeにずらす
		d.tail = d.tail.next
	}
}

// 両端キューの前から取り出し、その値を返す
func (d *Deque) DequeueFront() (int32, bool) {
	// front = head = 一番手前が対象
	// headの値を保持しておきつつ、headをhead.nextに移動させる

	if d.head == nil {
		return 0, false
	}

	temp := d.head
	d.head = temp.next
	// 要素が複数ある場合
	if d.head != nil {
		d.head.prev = nil
	} else {
		// d.headがnilつまり空なのでtailも
		d.tail = nil
	}

	return temp.data, true
}

// 両端キューの後ろから取り出し、その値を貸す
func (d *Deque) DequeueBack() (int32, bool) {
	// back = tail = 一番奥が対象

	// 要素が存在しない場合
	if d.tail == nil {
		return 0, false
	}

	temp := d.tail
	d.tail = temp.prev
	// 要素が複数ある場合
	if d.tail != nil {
		d.tail.next = nil
	} else {
		// 元の要素が1つしか無い場合、取り出しで消えるため、head, tailともにnilになるようにする
		d.head = nil
	}

	return temp.data, true
}
