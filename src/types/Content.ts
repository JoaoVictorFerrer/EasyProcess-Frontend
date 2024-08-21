import { User } from "./User"

export type Content = {
    creator: User
    comment: string
    fileURLs: string[]
    createdAt: Date
    updatedAt: Date
}