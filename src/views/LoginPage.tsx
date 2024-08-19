import { useContext } from "react";
import logo from "../img/logo2.png"
import {
  Card,
  Input,
  Button,
  Typography,
} from "@material-tailwind/react";
import { UserContext } from "../context/userContex";



export default function LoginPage() {

  const{loginUser} = useContext(UserContext)

  return (
    <div  className=" mx-auto py-24 flex flex-col justify-center items-center w-full h-screen bg-blue-gray-500">
      <Card color="transparent" shadow={false} className="bg-white m-2 p-8" >
        <div className="flex justify-center items-center ">
      <img src={logo} width= '200' alt="logo" />
        </div>
     <Typography variant="h4" color="blue-gray" className="text-center text-4xl mt-6 ">
       Área Clientes
     </Typography>

     <form className="mt-8 mb-2 w-80 max-w-screen-lg sm:w-96">
       <div className="mb-1 flex flex-col gap-6">

         <Typography variant="h6" color="blue-gray" className="-mb-3">
           Email
         </Typography>
         <Input
           size="lg"
           placeholder="name@mail.com"
           className=" !border-t-blue-gray-200 focus:!border-t-gray-900"
           labelProps={{
             className: "before:content-none after:content-none",
           }}
         />
         <Typography variant="h6" color="blue-gray" className="-mb-3">
           Contraseña
         </Typography>
         <Input
           type="password"
           size="lg"
           placeholder="********"
           className=" !border-t-blue-gray-200 focus:!border-t-gray-900"
           labelProps={{
             className: "before:content-none after:content-none",
           }}
         />
          </div>
          
       <Button className="mt-6" fullWidth onClick={loginUser} >
            Acceder
       </Button>

     </form>
   </Card>
    </div>
 );
}
