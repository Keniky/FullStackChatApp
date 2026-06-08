import { RoomIdContext } from "@/context/roomIdContext"
import { Cookie } from "lucide-react"
import { useState } from "react"
import { useEffect } from "react"
import { useContext } from "react"



export default function RoomPage(){
    const {roomId , setRoomId} = useContext(RoomIdContext)
    const [webSocket , setWebSocket] = useState();  

    //join room with cookie and room id
    useEffect(() =>{
        const url = new URL("ws://localhost:80/api/v1/chat")
        url.searchParams.set("room_id", roomId)
        const ws = new WebSocket(url);
        setWebSocket(ws);
        ws.onmessage = (msg) => console.log(msg.data);
        ws.onopen = () => ws.send(JSON.stringify({msg: "bruh"}));
    }, [])

    return (
        <div>test , {roomId}</div>
    )


}