import { BackgroundPaths } from "@/components/background-paths"
import { useState } from "react"

export default function ChatPage(){
    //get user data you must be logged in lol else idk man
    const [mode , setMode] = useState("dark")
    
    const [name , setName] = useState("Loading...")
    fetch("http://localhost:80/api/v1/user" , {
        method: "GET",
        credentials: "include"
    })
    .then(resp => resp.json())
    .then(data => { 
        if (data.name != null || data.name != undefined){
            setName(data.name)
            console.log("user name is " + data.name)
        } 
        
    })//check it not null

    //add switch mod to light or dark
    console.log(name)
    return (
        <div className={` ${mode}`}>
            <BackgroundPaths title={`Chat Now ${name}`}  />
        </div>
    )
 
}

    