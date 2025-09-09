package interfaces

type LL interface{
	InsertAtBegin(data int)
	InsertAtEnd(data int)
	Delete(data int)

	CountOfNode() int
	CheckNegativeData() bool
	SumOfAllData() int
	SumAtOddIndexes() int

	Print() 
}


