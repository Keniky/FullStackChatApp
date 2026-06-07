import './App.css'
import ChatPage from './pages/chatPage';
import LoginPage from './pages/loginPage'
import Auth from './components/Auth'

import { createBrowserRouter, RouterProvider } from "react-router-dom";

//before displaying anything get cookie and ask server if it is valid 
//if cookie is valid redirect to chat
//else redirect to log Page

const router = createBrowserRouter([
  {path: "/", element: <LoginPage/>},

  {
    //check if i can go to that page
      // element: <Auth/>,
      //do children
      //like this i have auth for all pages no need to worry :)
      children:[
        {path: "/chat", element:<ChatPage/>},
      ]
    }
])

function App() {

  
  return (
    <RouterProvider router={router}/>
  )
}

export default App
