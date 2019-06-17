package doubly_linked_list

import (
	"testing"
)

func TestCreate(t *testing.T) {
	list := DoublyLinkedList{}
	if &list == nil {
		t.Fatalf("List is not create")
	}

	t.Logf("PASS")
}

func TestHeadPushOneElements(t *testing.T) {
	list := DoublyLinkedList{}
	list.HeadPush("1")

	if list.head.value != "1" ||
		list.tail.value != "1" ||
		list.head.next != nil ||
		list.tail.prev != nil {
		t.Fatalf("Push Failure")
	}

	t.Logf("PASS")
}

func TestHeadPushTwoElements(t *testing.T) {
	list := DoublyLinkedList{}
	list.HeadPush("1")
	list.HeadPush("2")

	if list.head.value != "2" ||
		list.tail.value != "1" ||
		list.head.next.value != "1" ||
		list.tail.prev.value != "2" {
		t.Fatalf("Push Failure")
	}

	t.Logf("PASS")
}

func TestHeadPushThreeElements(t *testing.T) {
	list := DoublyLinkedList{}
	list.HeadPush("1")
	list.HeadPush("2")
	list.HeadPush("3")

	if list.head.value != "3" || list.head.next.value != "2" || list.head.next.next.value != "1" {
		t.Fatalf("Push Failure")
	}

	t.Logf("PASS")
}

func TestTailPushOneElements(t *testing.T) {
	list := DoublyLinkedList{}
	list.TailPush("1")

	if list.head.value != "1" ||
		list.tail.value != "1" ||
		list.head.next != nil ||
		list.tail.prev != nil {
		t.Fatalf("Push Failure")
	}

	t.Logf("PASS")
}
func TestTailPushTwoElements(t *testing.T) {
	list := DoublyLinkedList{}
	list.TailPush("1")
	list.TailPush("2")

	if list.head.value != "1" ||
		list.tail.value != "2"||
		list.head.next.value != "2" ||
		list.tail.prev.value != "1" {
		t.Fatalf("Push Failure")
	}

	t.Logf("PASS")
}
func TestTailPushThreeElements(t *testing.T) {
	list := DoublyLinkedList{}
	list.HeadPush("1")
	list.TailPush("2")
	list.TailPush("3")

	if list.tail.value != "3" || list.tail.prev.value != "2" || list.tail.prev.prev.value != "1" {
		t.Fatalf("Push Failure")
	}

	t.Logf("PASS")
}

func TestFind(t *testing.T) {
	list := DoublyLinkedList{}
	list.HeadPush("1")
	list.HeadPush("2")
	list.HeadPush("3")

	got := list.Find("2")
	if got.value != "2" {
		t.Fatalf("Find Failure")
	}

	t.Logf("PASS")
}

func TestHeadPop(t *testing.T) {
	list := DoublyLinkedList{}
	list.HeadPush("1")
	list.HeadPush("2")
	list.HeadPush("3")
	list.HeadPop()

	if list.head.value != "2" ||
		list.head.next.value != "1" ||
		list.head.prev != nil ||
		list.tail.value != "1" ||
		list.tail.prev.value != "2" {
		t.Fatalf("Pop Failure")
	}

	t.Logf("PASS")
}

func TestHeadPopTwoTimes(t *testing.T) {
	list := DoublyLinkedList{}
	list.HeadPush("1")
	list.HeadPush("2")
	list.HeadPush("3")
	list.HeadPop()
	list.HeadPop()

	if list.head.value != "1" ||
		list.tail.value != "1" ||
		list.head.prev != nil ||
		list.tail.next != nil {
		t.Fatalf("Pop Failure")
	}

	t.Logf("PASS")
}

func TestHeadPopInEmtyList(t *testing.T) {
	list := DoublyLinkedList{}
	_, err := list.HeadPop()

	if err == nil {
		t.Fatalf("Pop Failure")
	}

	t.Logf("PASS")
}

func TestTailPop(t *testing.T) {
	list := DoublyLinkedList{}
	list.HeadPush("1")
	list.HeadPush("2")
	list.HeadPush("3")
	list.TailPop()

	if list.tail.value != "2" ||
		list.head.value != "3" ||
		list.head.next.value != "2" ||
		list.tail.prev.value != "3" ||
		list.tail.next != nil ||
		list.head.prev != nil  {
		t.Fatalf("Pop Failure")
	}

	t.Logf("PASS")
}

func TestTailPopTwoTimes(t *testing.T) {
	list := DoublyLinkedList{}
	list.HeadPush("1")
	list.HeadPush("2")
	list.HeadPush("3")
	list.TailPop()
	list.TailPop()

	if list.tail.value != "3" ||
		list.head.value != "3" ||
		list.tail.prev != nil ||
		list.head.next != nil ||
		list.tail.next != nil ||
		list.head.prev != nil  {
		t.Fatalf("Pop Failure")
	}

	t.Logf("PASS")
}

func TestTailPopInEmtyList(t *testing.T) {
	list := DoublyLinkedList{}
	_, err := list.TailPop()

	if err == nil {
		t.Fatalf("Pop Failure")
	}

	t.Logf("PASS")
}

func TestDelete(t *testing.T) {
	list := DoublyLinkedList{}
	list.HeadPush("1")
	list.HeadPush("2")
	list.HeadPush("3")
	list.Delete("2")

	if list.head.next.value != "1" || list.tail.prev.value != "3" {
		t.Fatalf("Pop Failure")
	}

	t.Logf("PASS")
}