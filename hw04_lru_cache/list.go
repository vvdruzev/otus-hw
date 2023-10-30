package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	front  *ListItem
	back   *ListItem
	length int
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	l.length++
	front := &ListItem{
		Value: v,
		Next:  l.front,
		Prev:  nil,
	}
	if l.front == nil {
		l.front = front
		l.back = front
		return l.front
	}
	current := l.front
	l.front = front
	current.Prev = l.front
	return l.front
}

func (l *list) PushBack(v interface{}) *ListItem {
	l.length++
	current := l.back
	back := &ListItem{
		Value: v,
		Next:  nil,
		Prev:  current,
	}
	if l.back == nil {
		l.back = back
		l.front = back
		return back
	}
	current.Next = back
	l.back = back
	return back
}

func (l *list) Remove(i *ListItem) {
	l.length--
	if l.length == 0 {
		l.front = nil
		l.back = nil
		return
	}
	if i.Prev == nil {
		i.Next.Prev = nil
		l.front = i.Next
		return
	}
	if i.Next == nil {
		i.Prev.Next = nil
		l.back = i.Prev
		return
	}
	i.Prev.Next = i.Next
	i.Next.Prev = i.Prev
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}

func NewList() List {
	return new(list)
}
