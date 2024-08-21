import { createContext, ReactNode} from "react";
import { BaseServerResponse, HTTPMethod } from "../services/Network/Entities";
import { ServerEndpoints } from "../services/Network/ServerEndpoints";
import performRequest from "../services/Network/DataFetcher";

type UserProviderProps = {
    children: ReactNode
}

type UserContextProps = {
    loginUser:  () => void
}
export const UserContext = createContext<UserContextProps>(null!);


export const UserContextProvider = ({ children }: UserProviderProps) => { 
    

    const loginUser = async () => {
        try {
            
            const body = {
                Email: "som@email.com",
                password: "12345678"
            }
            const response : BaseServerResponse = await performRequest(HTTPMethod.POST, ServerEndpoints.login(), body)
            console.log(response)
        } catch (error) {
           
           console.log({error}) 
        }
    
    }

    const userContextValue = {
        loginUser
    }

    return (
        <UserContext.Provider value={userContextValue}>
            {children}
        </UserContext.Provider>
    )

}