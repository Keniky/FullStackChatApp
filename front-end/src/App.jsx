import './App.css'
import ChatPage from './pages/chatPage';
import LoginPage from './pages/loginPage'
// import Auth from './components/Auth'

import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { useState } from 'react';
import { RoomContext } from './context/roomContext';
import RoomPage from './pages/chatRoomPage';
import ChatRoomPage from './pages/chatRoomPage';

//before displaying anything get cookie and ask server if it is valid 
//if cookie is valid redirect to chat
//else redirect to log Page


const router = createBrowserRouter([
  {path: "/", element: <LoginPage/>},

  {
    //check if i can go to that page
      // element: <Auth/>,
      errorElement: [
      <ChatPage/>,
      <RoomPage/>
    ],
      //do children
      //like this i have auth for all pages no need to worry :)
      children:[
        {path: "/chat", element:<ChatPage/>},
        {path: "/chat/room", element:<ChatRoomPage/>},

      ]
    }
])


function App() {
  const [roomId , setRoomId ] = useState("")
  const [name , setName] = useState("guest")

  return (
    <RoomContext.Provider value = {{ roomId , setRoomId , name , setName}}>
      <RouterProvider router={router}/>
    </RoomContext.Provider>
  )
}

export default App
