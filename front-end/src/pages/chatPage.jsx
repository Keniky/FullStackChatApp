import { BackgroundPaths } from "@/components/background-paths"
import { RoomContext } from "@/context/roomContext"
import { useContext } from "react"
import { useEffect } from "react"
import { useState } from "react"

export default function ChatPage(){
    //get user data you must be logged in lol else idk man
    const [mode , setMode] = useState("dark")
    const {name , setName , pfp ,setPfp} = useContext(RoomContext)    

    console.log("place: chatPage"+pfp)
    useEffect(() => {
        fetch("http://localhost:80/api/v1/user" , {
            method: "GET",
            credentials: "include"
        })
        .then(resp => resp.json())
        .then(data => { 
            if (data.name != null || data.name != undefined){
                setName(data.name)
                setMode("dark")
            } 
            if (data.pfp !=null || data.pfp != undefined ){
               setPfp(data.pfp)
                console.log("profile pic" + data.pfp)
            }
            
        })//check it not null
    }, [])

    //add switch mod to light or dark
    return (
        <div className={` ${mode}`}>
            <BackgroundPaths title={`Chat Now ${name}`}  />
        </div>
    )
 
}

    