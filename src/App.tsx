import { Route, Routes } from "react-router-dom"
import LoginPage from "./views/LoginPage.js"
import MainPage from "./views/MainPage.js"
import LayoutsMain from "./layouts/LayoutsMain.js"


function App() {
 
  return (
    <>
        <Routes>
        <Route element={<LoginPage/>} path="/" />
        <Route element={<LayoutsMain />}>
          <Route element={<MainPage />} path="/user/:id" />
        </Route>
  
        </Routes>

    </>
  )
}

export default App
