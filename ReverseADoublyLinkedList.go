package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

type DoublyLinkedListNode struct {
    data int32
    next *DoublyLinkedListNode
    prev *DoublyLinkedListNode
}

type DoublyLinkedList struct {
    head *DoublyLinkedListNode
    tail *DoublyLinkedListNode
}

func (doublyLinkedList *DoublyLinkedList) insertNodeIntoDoublyLinkedList(nodeData int32) {
    node := &DoublyLinkedListNode {
        next: nil,
        prev: nil,
        data: nodeData,
    }

    if doublyLinkedList.head == nil {
        doublyLinkedList.head = node
    } else {
        doublyLinkedList.tail.next = node
        node.prev = doublyLinkedList.tail
    }

    doublyLinkedList.tail = node
}

func printDoublyLinkedList(node *DoublyLinkedListNode, sep string, writer *bufio.Writer) {
    for node != nil {
        fmt.Fprintf(writer, "%d", node.data)

        node = node.next

        if node != nil {
            fmt.Fprintf(writer, sep)
        }
    }
}

/*
 * Complete the 'reverse' function below.
 *
 * The function is expected to return an INTEGER_DOUBLY_LINKED_LIST.
 * The function accepts INTEGER_DOUBLY_LINKED_LIST llist as parameter.
 */

/*
 * For your reference:
 *
 * DoublyLinkedListNode {
 *     data int32
 *     next *DoublyLinkedListNode
 *     prev *DoublyLinkedListNode
 * }
 *
 */

func reverse(llist *DoublyLinkedListNode) *DoublyLinkedListNode {
    // Write your code here
    var pre *DoublyLinkedListNode = nil
    var cur *DoublyLinkedListNode = llist
    var tmp *DoublyLinkedListNode
    for cur != nil {
        tmp = cur.next
        cur.next = pre
        cur.prev = tmp
        pre = cur
        cur = tmp
    }
    return pre
}

func main() {
