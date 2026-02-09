class MyCircularQueue:
    def __init__(self, k: int):
        self.data = [0] * k
        self.maxSize = k
        self.head = 0
        self.tail = -1
    def enQueue(self, val: int) -> bool:
        if self.isFull(): return False
        self.tail = (self.tail + 1) % self.maxSize
        self.data[self.tail] = val
        return True
    def deQueue(self) -> bool:
        if self.isEmpty(): return False
        if self.head == self.tail: self.head, self.tail = 0, -1
        else: self.head = (self.head + 1) % self.maxSize
        return True
    def Front(self) -> int:
        return -1 if self.isEmpty() else self.data[self.head]
    def Rear(self) -> int:
        return -1 if self.isEmpty() else self.data[self.tail]
    def isEmpty(self) -> bool:
        return self.tail == -1
    def isFull(self) -> bool:
        return not self.isEmpty() and (self.tail + 1) % self.maxSize == self.head
    
# Example usage:
circularQueue = MyCircularQueue(3)
# check if the queue is empty
print(circularQueue.isEmpty())  # returns True
print(circularQueue.enQueue(1))  # returns True
print(circularQueue.enQueue(2))  # returns True
print(circularQueue.enQueue(3))  # returns True
print(circularQueue.enQueue(4))  # returns False, the queue is full
print(circularQueue.Rear())      # returns 3
print(circularQueue.isFull())    # returns True
print(circularQueue.deQueue())   # returns True 
print(circularQueue.isEmpty())  # returns False