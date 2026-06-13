import { RoomContext } from "@/context/roomContext"
import { useContext } from "react"


export function RoomHeader( ){
    //an a sync function that runs forever and get the room members
    const {roomId} = useContext(RoomContext)

    return(
        <header className="border border-solid border-black bg-[#2B2D31] h-[3rem] w-[full] flex items-center justify-start flex-row "> 
            <section className="flex flex-row gap-[1rem] text-white p-[1rem]">
                <span role="img" aria-label="rocket">#🌍</span> 
                <span>global-chat</span>
                <span>Room id : {roomId}</span> 
            </section>
        </header>
    )
}