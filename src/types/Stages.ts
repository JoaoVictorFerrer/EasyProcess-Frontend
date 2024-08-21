import { Content } from "./Content"


export type Stage = {
	name:        string            
	isCompleted: boolean            
	contents:   Content       
	createdAt:   Date         
	UupdatedAt:   Date       
}

export type StageProcess = Stage[]