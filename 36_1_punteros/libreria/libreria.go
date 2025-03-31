package libreria 

type Pc struct {
	Ram int 
	Disk int 
	Brand string 
}

func (myPc *Pc)DuplicateRam(){
	myPc.Ram = myPc.Ram * 2
}


