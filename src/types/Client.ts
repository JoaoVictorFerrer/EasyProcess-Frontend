import { ObjectId } from "bson"


export type Client = {
	id:            ObjectId
	name:           string            
	surname:        string            
	email:          string            
	phone:          string            
	processIds:     string[]          
	documentType:   string            
	documentNumber: string            
	createdAt:      Date        
	updatedAt:      Date   
}

export enum DocumentClient {
    passport = "Passaporte",
	NIE      = "NIE",
	DNI      = "DNI",
	RG       = "RG",
	CPF      = "CPF",
}



