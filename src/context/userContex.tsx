import { createContext, ReactNode} from "react";

type UserProviderProps = {
    children: ReactNode
}

type UserContextProps = {
    loginUser:  () => void
}
export const UserContext = createContext<UserContextProps>(null!);


export const UserContextProvider = ({ children} : UserProviderProps ) => { 
    const loginUser = () => {
        console.log('desde el contect user')
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