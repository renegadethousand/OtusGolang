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
	size  int
	front *ListItem
	back  *ListItem
}

func (l *list) Len() int {
	return l.size
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	oldFront := l.Front()
	newFront := ListItem{v, oldFront, nil}
	if oldFront != nil {
		oldFront.Prev = &newFront
	}
	l.front = &newFront
	if l.Back() == nil {
		l.back = &newFront
	}
	l.size++
	return &newFront
}

func (l *list) PushBack(v interface{}) *ListItem {
	oldBack := l.Back()
	newBack := ListItem{v, nil, oldBack}
	if oldBack != nil {
		oldBack.Next = &newBack
	}
	l.back = &newBack
	if l.Front() == nil {
		l.front = &newBack
	}
	l.size++
	return &newBack
}

func (l *list) Remove(i *ListItem) {
	left := i.Prev
	right := i.Next
	if left != nil {
		left.Next = right
	}
	if right != nil {
		right.Prev = left
	}
	if l.Front() == i && right != nil {
		l.front = right
	}
	if l.Back() == i && left != nil {
		l.back = left
	}
	l.size--
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}

func NewList() List {
	return new(list)
}
