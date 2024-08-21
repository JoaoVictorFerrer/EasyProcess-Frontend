
import { ObjectId } from "bson"
import { TimeLine } from "./TimeLine"

export type Process = {
	id:            ObjectId
	ownerUserId:   string            
	paymentStatus: string            
	clientId:      string            
	description:   string            
	startDate:     Date        
	endDate:       Date    
	timeline:      TimeLine   
	processType:   string        
	isCompleted:   boolean         
	currentStage:  string        
	createdAt:     Date       
	updatedAt:     Date    
}



export enum PaymentStatus {
    Payed          = "Pago",
	NotPayed       = "Não Pago",
	PaymentPending = "Pagamento Pendente",
}

export enum PossibleProcess {
    Citizenship = "Cidadania"
}