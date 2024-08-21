import { ObjectId } from "bson"

 
export type ProcessGlance = {
	id:                   ObjectId
	ownerUserId:          string           
	representedProcessId: string           
	processType:          string          
	isCompleted:          boolean            
	currentStage:         string            
	startDate:            Date      
	endDate:              Date       
	clientName:           string           
	createdAt:            Date      
	updatedAt:            Date      
}