package interfaces

type LL interface{
	InsertAtBegin(data int) 
	InsertAtEnd(data int)
	Delete(data int)
	Print()
	CountOfNode() 
	CheckNegativeData() 
	SumOfAllData()
	SumAtOddIndexes()
}

