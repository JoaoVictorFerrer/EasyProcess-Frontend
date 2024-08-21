import { ObjectId } from "bson"
import { StageProcess } from "./Stages"

export type TimeLine = {
    id:         ObjectId
	processType: string             
	stages:      StageProcess        
	createdAt:   Date         
	updatedAt:   Date  
}