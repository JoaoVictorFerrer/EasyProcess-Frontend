
import performRequest from './services/Network/DataFetcher'
import { HTTPMethod } from './services/Network/Entities'
import { ServerEndpoints } from './services/Network/ServerEndpoints'

function App() {

  const checkServices = async () => {

    const body = {
        Email: "some@email.com",
        password: "12345678"
    }

    const response = await performRequest(HTTPMethod.POST, ServerEndpoints.login(),body)
    console.log('esto es response',response)
  }

  return (
    <>
      <div className='flex justify-center items-center flex-col size-full bg-slate-600 m-auto '>
        <h1 className='text-6xl'>Hola mundo</h1>
        <button onClick={checkServices}>click me</button>
        </div>
    </>
  )
}

export default App
