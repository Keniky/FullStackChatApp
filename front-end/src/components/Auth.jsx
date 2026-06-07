import { useEffect } from "react";
import { useState } from "react";
import { Navigate, Outlet } from "react-router-dom";


export default function Auth(){
   
    const [status , setStatus] = useState("checking")
    
    //run once rendering finished
    useEffect(() => {
        fetch("http://localhost:80/api/v1/auth" , {
            credentials: "include"
        })
        .then(resp => {
            console.log("authentication completed " , resp.status)    
            resp.status === 202 ? setStatus("yes") : setStatus("no")
        })
        .catch(() => setStatus("no"))
    } , [])



    //render loading while we are searchign the user
    if (status === "checking"){return <p>Loading .. </p>}
    //if there is no user go to navigate 
    if (status === "no"){return <Navigate to="/"/>}

    //if there is a user you have been authenticated
    //do what childrent will do 
    return <Outlet/>
}